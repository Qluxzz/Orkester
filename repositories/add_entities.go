package repositories

import (
	"context"
	"errors"
	"goreact/ent"
	"goreact/ent/album"
	"goreact/ent/artist"
	"goreact/ent/track"
	"goreact/indexFiles"

	"github.com/gosimple/slug"
)

func AddTracks(tracks []*indexFiles.IndexedTrack, client *ent.Client, context context.Context) ([]*ent.Track, error) {
	tx, err := client.Tx(context)
	if err != nil {
		return nil, err
	}

	added_tracks := []*ent.Track{}

	for _, track_on_disk := range tracks {
		artists := []*ent.Artist{}

		for _, artist := range track_on_disk.Artists {
			a, err := GetOrCreateArtist(artist, context, tx)
			if err != nil {
				tx.Rollback()
				return nil, err
			}

			artists = append(artists, a)
		}

		var albumArtist *ent.Artist

		if track_on_disk.AlbumArtist != "" {
			albumArtist, err = GetOrCreateArtist(track_on_disk.AlbumArtist, context, tx)
			if err != nil {
				tx.Rollback()
				return nil, err
			}

		} else {
			albumArtist = artists[0]
		}

		var album *ent.Album

		if track_on_disk.AlbumName != "" {
			album, err = GetOrCreateAlbum(track_on_disk.AlbumName, track_on_disk.Image, albumArtist, context, tx)
			if err != nil {
				tx.Rollback()
				return nil, err
			}
		}

		added_track, err := tx.
			Track.
			Create().
			SetTitle(track_on_disk.Title).
			SetTrackNumber(int(track_on_disk.TrackNumber)).
			SetPath(track_on_disk.Path).
			SetLength(int(track_on_disk.Length)).
			SetMimetype(track_on_disk.MimeType).
			SetAlbum(album).
			AddArtists(artists...).
			Save(context)

		if err == nil {
			track, err := GetTrackById(added_track.ID, context, tx)

			if err == nil {
				added_tracks = append(added_tracks, track)
			}
		} else {
			_, ok := err.(*ent.ConstraintError)
			if !ok {
				tx.Rollback()
				return nil, err
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return added_tracks, nil
}

func GetTrackById(id int, context context.Context, client *ent.Tx) (*ent.Track, error) {
	return client.
		Track.
		Query().
		Where(track.ID(id)).
		WithAlbum(func(aq *ent.AlbumQuery) {
			aq.Select(album.FieldName)
		}).
		WithArtists().
		Only(context)
}

func GetOrCreateAlbum(albumName string, albumImage *indexFiles.Image, albumArtist *ent.Artist, context context.Context, client *ent.Tx) (*ent.Album, error) {
	a, err := client.
		Album.
		Query().
		Where(
			album.And(
				album.NameEqualFold(albumName),
				album.HasArtistWith(artist.ID(albumArtist.ID)),
			),
		).Only(context)

	if err == nil {
		return a, nil
	}

	if _, ok := err.(*ent.NotFoundError); ok {
		a, err := client.Album.
			Create().
			SetName(albumName).
			SetURLName(slug.Make(albumName)).
			SetImage(albumImage.Data).
			SetImageMimeType(albumImage.MimeType).
			SetArtist(albumArtist).
			Save(context)

		if err == nil {
			return a, nil
		}
	}

	return nil, errors.New("failed to find or create album")
}

func GetOrCreateArtist(name string, context context.Context, client *ent.Tx) (*ent.Artist, error) {
	a, err := client.
		Artist.
		Query().
		Where(
			artist.NameEqualFold(name),
		).
		Only(context)

	if err == nil {
		return a, nil
	}

	if _, ok := err.(*ent.NotFoundError); ok {
		a, err := client.Artist.
			Create().
			SetName(name).
			SetURLName(slug.Make(name)).
			Save(context)

		if err == nil {
			return a, nil
		}
	}

	return nil, errors.New("failed to find or create artist")
}
