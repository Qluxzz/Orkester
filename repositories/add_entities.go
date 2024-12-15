package repositories

import (
	"context"
	"errors"
	"orkester/ent"
	"orkester/ent/album"
	"orkester/ent/artist"
	"orkester/ent/image"
	"orkester/indexFiles"
	"strconv"

	"github.com/cespare/xxhash/v2"
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
		artistIds := []int{}

		for _, artist := range track_on_disk.Artists {
			artistId, err := GetOrCreateArtist(artist, context, tx)
			if err != nil {
				tx.Rollback()
				return nil, err
			}

			artistIds = append(artistIds, artistId)
		}

		var albumArtistId int

		if track_on_disk.AlbumArtist != "" {
			albumArtistId, err = GetOrCreateArtist(track_on_disk.AlbumArtist, context, tx)
			if err != nil {
				tx.Rollback()
				return nil, err
			}

		} else {
			albumArtistId = artistIds[0]
		}

		var albumId int

		if track_on_disk.AlbumName != "" {
			albumId, err = GetOrCreateAlbum(
				track_on_disk.AlbumName,
				track_on_disk.Date,
				track_on_disk.Image,
				albumArtistId,
				context,
				tx,
			)
			if err != nil {
				tx.Rollback()
				return nil, err
			}
		}

		trackImageId, err := GetOrCreateImage(track_on_disk.Image, context, tx)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		added_track, err := tx.
			Track.
			Create().
			SetTitle(track_on_disk.Title).
			SetTrackNumber(int(track_on_disk.TrackNumber)).
			SetPath(track_on_disk.Path).
			SetLength(int(track_on_disk.Length)).
			SetMimetype(track_on_disk.MimeType).
			SetImageID(trackImageId).
			SetAlbumID(albumId).
			AddArtistIDs(artistIds...).
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

func GetOrCreateImage(img *indexFiles.Image, context context.Context, client *ent.Tx) (int, error) {
	hash := strconv.FormatUint(xxhash.Sum64(img.Data), 10)

	i, err := client.
		Image.
		Query().
		Where(image.HashEQ(hash)).
		Select(image.FieldID).
		Only(context)

	if err == nil {
		return i.ID, nil
	}

	if _, ok := err.(*ent.NotFoundError); ok {
		i, err = client.
			Image.
			Create().
			SetImage(img.Data).
			SetImageMimeType(img.MimeType).
			SetHash(hash).
			Save(context)

		if err != nil {
			return 0, err
		}

		return i.ID, nil
	}

	return 0, errors.New("failed to find or create image")
}

func GetOrCreateAlbum(albumName string, released *indexFiles.ReleaseDate, albumImage *indexFiles.Image, albumArtistId int, context context.Context, client *ent.Tx) (int, error) {
	a, err := client.
		Album.
		Query().
		Where(
			album.And(
				album.NameEqualFold(albumName),
				album.HasArtistWith(artist.ID(albumArtistId)),
			),
		).
		Select(album.FieldID).
		Only(context)

	if err == nil {
		return a.ID, nil
	}

	if _, ok := err.(*ent.NotFoundError); ok {
		imageId, err := GetOrCreateImage(albumImage, context, client)

		if err != nil {
			return 0, err
		}

		a, err := client.Album.
			Create().
			SetName(albumName).
			SetURLName(slug.Make(albumName)).
			SetCoverID(imageId).
			SetArtistID(albumArtistId).
			SetReleased(released).
			Save(context)

		if err == nil {
			return a.ID, nil
		}
	}

	return 0, errors.New("failed to find or create album")
}

func GetOrCreateArtist(name string, context context.Context, client *ent.Tx) (int, error) {
	a, err := client.
		Artist.
		Query().
		Where(
			artist.NameEqualFold(name),
		).
		Select(artist.FieldID).
		Only(context)

	if err == nil {
		return a.ID, nil
	}

	if _, ok := err.(*ent.NotFoundError); ok {
		a, err := client.Artist.
			Create().
			SetName(name).
			SetURLName(slug.Make(name)).
			Save(context)

		if err == nil {
			return a.ID, nil
		}
	}

	return 0, errors.New("failed to find or create artist")
}
