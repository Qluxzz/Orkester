package repositories

import (
	"context"
	"errors"
	"orkester/ent"
	"orkester/ent/album"
	"orkester/ent/artist"
	"orkester/indexFiles"

	"github.com/gosimple/slug"
)

type TrackIds = []int

func AddTracks(tracks []*indexFiles.IndexedTrack, client *ent.Client, context context.Context) (TrackIds, error) {
	tx, err := client.Tx(context)
	if err != nil {
		return nil, err
	}

	addedTrackIds := TrackIds{}

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
			album, err = GetOrCreateAlbum(
				track_on_disk.AlbumName,
				track_on_disk.Date,
				track_on_disk.Image,
				albumArtist,
				context,
				tx,
			)
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
			addedTrackIds = append(addedTrackIds, added_track.ID)
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

	return addedTrackIds, nil
}

func GetOrCreateAlbum(albumName string, released *indexFiles.ReleaseDate, albumImage *indexFiles.Image, albumArtist *ent.Artist, context context.Context, client *ent.Tx) (*ent.Album, error) {
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
		image, err := client.
			AlbumImage.
			Create().
			SetImage(albumImage.Data).
			SetImageMimeType(albumImage.MimeType).
			Save(context)

		if err != nil {
			return nil, err
		}

		a, err := client.Album.
			Create().
			SetName(albumName).
			SetURLName(slug.Make(albumName)).
			SetCover(image).
			SetArtist(albumArtist).
			SetReleased(released).
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
