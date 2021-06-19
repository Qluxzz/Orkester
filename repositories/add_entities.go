package repositories

import (
	"context"
	"errors"
	"goreact/ent"
	"goreact/ent/album"
	"goreact/ent/artist"
	"goreact/indexFiles"
	"time"

	"github.com/gosimple/slug"
)

func AddTracks(tracks []*indexFiles.IndexedTrack, client *ent.Client, context context.Context) (int, error) {
	tx, err := client.Tx(context)
	if err != nil {
		return 0, err
	}

	tracks_added := 0

	for _, track := range tracks {
		artists := []*ent.Artist{}

		for _, artist := range track.Artists {
			a, err := GetOrCreateArtist(artist.String, context, tx)
			if err != nil {
				return 0, err
			}

			artists = append(artists, a)
		}

		var albumArtist *ent.Artist

		if track.AlbumArtist.Valid {
			albumArtist, err = GetOrCreateArtist(track.AlbumArtist.String, context, tx)
			if err != nil {
				return 0, err
			}

		} else {
			albumArtist = artists[0]
		}

		var album *ent.Album

		if track.AlbumName.Valid {
			album, err = GetOrCreateAlbum(track, albumArtist, context, tx)
			if err != nil {
				return 0, err
			}
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
				return 0, err
			}
		} else {
			tracks_added += 1
		}
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return tracks_added, nil
}

func GetOrCreateAlbum(track *indexFiles.IndexedTrack, albumArtist *ent.Artist, context context.Context, client *ent.Tx) (*ent.Album, error) {
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
		return a, nil
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
			artist.NameEQ(name),
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
