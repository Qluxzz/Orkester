package repositories

import (
	"context"
	"goreact/ent"
	"goreact/ent/album"
	"goreact/ent/artist"
	"goreact/indexFiles"
	"log"
	"time"

	"github.com/gosimple/slug"
)

func AddTracks(tracks []*indexFiles.IndexedTrack, client *ent.Client, context context.Context) {
	tx, err := client.Tx(context)
	if err != nil {
		log.Fatal(err)
	}

	tracks_added := 0

	for _, track := range tracks {
		artists := []*ent.Artist{}

		for _, artist := range track.Artists {
			a := GetOrCreateArtist(artist.String, context, tx)

			artists = append(artists, a)
		}

		var albumArtist *ent.Artist

		if track.AlbumArtist.Valid {
			albumArtist = GetOrCreateArtist(track.AlbumArtist.String, context, tx)
		} else {
			albumArtist = artists[0]
		}

		var album *ent.Album

		if track.AlbumName.Valid {
			album = GetOrCreateAlbum(track, albumArtist, context, tx)
		}

		_, err := tx.
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
			log.Printf("Added track %s from album %s", track.Title.String, track.AlbumName.String)

			tracks_added += 1
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Added %d tracks", tracks_added)
}

func GetOrCreateAlbum(track *indexFiles.IndexedTrack, albumArtist *ent.Artist, context context.Context, client *ent.Tx) *ent.Album {
	a, err := client.
		Album.
		Query().
		Where(
			album.And(
				album.NameEQ(track.AlbumName.String),
				album.HasArtistWith(artist.ID(albumArtist.ID)),
			),
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
			SetArtist(albumArtist).
			Save(context)

		if err == nil {
			return a
		}
	}

	panic("failed to find or create album")
}

func GetOrCreateArtist(name string, context context.Context, client *ent.Tx) *ent.Artist {
	a, err := client.
		Artist.
		Query().
		Where(
			artist.NameEQ(name),
		).
		Only(context)

	if err == nil {
		return a
	}

	if _, ok := err.(*ent.NotFoundError); ok {
		a, err := client.Artist.
			Create().
			SetName(name).
			SetURLName(slug.Make(name)).
			Save(context)

		if err == nil {
			return a
		}
	}

	panic("failed to find or create artist")
}
