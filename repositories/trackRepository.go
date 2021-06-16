package repositories

import (
	"context"
	"goreact/ent"
	"goreact/ent/album"
	"goreact/ent/artist"
	"goreact/ent/track"
	"goreact/indexFiles"
	"log"
	"time"

	"github.com/gosimple/slug"
)

func AddTracks(tracks []*indexFiles.IndexedTrack, client *ent.Client, context context.Context) {
	tracks_added := 0

	for _, track := range tracks {
		artists := []*ent.Artist{}

		for _, artist := range track.Artists {
			a := GetOrCreateArtist(artist.String, context, client)

			artists = append(artists, a)
		}

		var albumArtist *ent.Artist

		if track.AlbumArtist.Valid {
			albumArtist = GetOrCreateArtist(track.AlbumArtist.String, context, client)
		}

		if albumArtist == nil {
			albumArtist = artists[0]
		}

		var album *ent.Album

		if track.AlbumName.Valid {
			album = GetOrCreateAlbum(track, albumArtist, context, client)
		}

		_, err := client.
			Track.
			Create().
			SetTitle(track.Title.String).
			SetTrackNumber(int(track.TrackNumber.Int32)).
			SetPath(track.Path.String).
			SetDate(time.Now()).
			SetLength(int(track.Length.Int32)).
			SetMimetype(track.MimeType.String).
			SetAlbum(album).
			AddArtists(artists...).
			Save(context)

		if err != nil {
			_, ok := err.(*ent.ConstraintError)
			if !ok {
				log.Fatalf("failed to create track %v", err)
			}
		} else {
			tracks_added += 1
		}

	}

	log.Printf("Added %d tracks", tracks_added)
}

func GetOrCreateAlbum(track *indexFiles.IndexedTrack, albumArtist *ent.Artist, context context.Context, client *ent.Client) *ent.Album {
	a, err := client.
		Album.
		Query().
		Where(
			album.NameEqualFold(track.AlbumName.String),
			album.HasArtistWith(artist.ID(albumArtist.ID)),
		).Only(context)

	if err == nil {
		return a
	}

	if _, ok := err.(*ent.NotFoundError); ok {
		a, err := client.Album.
			Create().
			SetName(track.AlbumName.String).
			SetURLName(slug.Make(track.AlbumName.String)).
			SetImage(track.Image.Data).
			SetImageMimeType(track.Image.MimeType.String).
			SetArtist(albumArtist).Save(context)

		if err != nil {
			log.Fatalf("failed to create album %v", err)
		}

		return a
	}

	panic("failed to find or create album")
}

func GetOrCreateArtist(name string, context context.Context, client *ent.Client) *ent.Artist {
	a, err := client.Artist.Query().Where(artist.NameEqualFold(name)).Only(context)

	if err == nil {
		return a
	}

	if _, ok := err.(*ent.NotFoundError); ok {
		a, err := client.Artist.
			Create().
			SetName(name).
			SetURLName(slug.Make(name)).
			Save(context)

		if err != nil {
			log.Fatalf("failed to create artist %v", err)
		}

		return a
	}

	panic("failed to find or create artist")
}

func RemoveDeletedEntities(tracks []*indexFiles.IndexedTrack, client *ent.Client, context context.Context) error {
	existing_tracks, err := client.
		Track.
		Query().
		WithAlbum(func(aq *ent.AlbumQuery) {
			aq.Select(album.FieldName)
		}).
		Select(track.FieldTrackNumber, track.FieldTitle).
		All(context)

	if err != nil {
		return err
	}

	removed_tracks := 0

	for _, dbTrack := range existing_tracks {
		exists := false

		for _, track := range tracks {
			if track.Title.String == dbTrack.Title &&
				track.TrackNumber.Int32 == int32(dbTrack.TrackNumber) &&
				track.AlbumName.String == dbTrack.Edges.Album.Name {
				exists = true
				break
			}
		}

		if !exists {
			err := client.Track.DeleteOneID(dbTrack.ID).Exec(context)
			if err != nil {
				return err
			}
			removed_tracks += 1
		}
	}

	log.Printf("Removed %d tracks", removed_tracks)

	// Remove albums without tracks
	removed_albums, err := client.Album.Delete().Where(album.Not(album.HasTracks())).Exec(context)
	if err != nil {
		return err
	}

	log.Printf("Removed %d albums", removed_albums)

	// Remove artists without albums
	removed_artists, err := client.Artist.Delete().Where(artist.Not(artist.HasAlbums())).Exec(context)
	if err != nil {
		return err
	}

	log.Printf("Removed %d artists", removed_artists)

	return nil
}
