package repositories

import (
	"goreact/indexFiles"
	"goreact/models"

	"github.com/jmoiron/sqlx"
)

func AddTracks(tracks []indexFiles.IndexedTrack, db *sqlx.DB) error {
	insertArtistStmt := `
		INSERT OR IGNORE INTO artists (name) VALUES (?)
	`

	insertAlbumStmt := `
		INSERT OR IGNORE INTO albums (name, image, imagemimetype) VALUES (?, ?, ?)
	`

	insertGenreStmt := `
		INSERT OR IGNORE INTO genres (name) VALUES (?)
	`

	insertTrackStmt := `
		INSERT INTO tracks (
			title,
			tracknumber,
			path,
			date,
			albumid,
			artistid,
			genreid
		) VALUES(
			?,
			?,
			?,
			?,
			(SELECT id FROM albums WHERE name = ?),
			(SELECT id FROM artists WHERE name = ?),
			(SELECT id from genres WHERE name = ?)
		)
	`

	tx := db.MustBegin()

	for _, track := range tracks {
		if track.Artist != "" {
			tx.MustExec(insertArtistStmt, track.Artist)
		}

		if track.Album.Name != "" {
			tx.MustExec(insertAlbumStmt, track.Album.Name, track.Album.Image.Data, track.Album.Image.MimeType)
		}

		if track.Genre != "" {
			tx.MustExec(insertGenreStmt, track.Genre)
		}

		tx.MustExec(
			insertTrackStmt,
			track.Title,
			track.TrackNumber,
			track.Path,
			track.Date,
			track.Album.Name,
			track.Artist,
			track.Genre,
		)
	}

	err := tx.Commit()
	return err
}

func GetTracksByIds(ids []int, db *sqlx.DB) ([]models.Track, error) {
	query, args, err := sqlx.In(`
			SELECT
				t.id,
				t.title,
				t.tracknumber,
				t.date,
				albums.name album,
				artists.name artist,
				genres.name genre
			FROM
				tracks t
			LEFT JOIN artists
				ON artists.id = t.artistid
			LEFT JOIN albums
				ON albums.id = t.albumid
			LEFT JOIN genres
				ON genres.id = t.genreid
			WHERE t.id IN (?)
		`,
		ids,
	)

	if err != nil {
		return nil, err
	}

	dbTracks := []models.DBTrack{}
	err = db.Select(&dbTracks, query, args...)
	if err != nil {
		return nil, err
	}

	tracks := []models.Track{}

	for _, dbTrack := range dbTracks {
		tracks = append(tracks, dbTrack.ToDomain())
	}

	return tracks, nil
}
