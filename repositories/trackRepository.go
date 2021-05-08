package repositories

import (
	"goreact/indexFiles"

	"github.com/gosimple/slug"
	"github.com/jmoiron/sqlx"
)

func AddTracks(tracks []*indexFiles.IndexedTrack, db *sqlx.DB) error {
	insertArtistStmt := `
		INSERT INTO artists (name, urlname) VALUES (?, ?) ON CONFLICT DO NOTHING
	`

	insertAlbumStmt := `
		INSERT INTO albums (
			name, 
			urlname, 
			image, 
			imagemimetype, 
			artistid
		) VALUES (
			?,
			?, 
			?, 
			?, 
			(SELECT id FROM artists WHERE name = ?)
		) ON CONFLICT DO NOTHING
	`

	insertGenreStmt := `
		INSERT INTO genres (name, urlname) VALUES (?, ?) ON CONFLICT DO NOTHING
	`

	insertTrackStmt := `
		INSERT INTO tracks (
			title,
			tracknumber,
			path,
			date,
			length,
			albumid,
			artistid,
			genreid
		) VALUES(
			$1,
			$2,
			$3,
			$4,
			$5,
			(
				SELECT 
					id 
				FROM 
					albums 
				WHERE 
					name = $6
					AND artistid = (
						SELECT id FROM artists WHERE name = $7
					)
			),
			(SELECT id FROM artists WHERE name = $7),
			(SELECT id from genres WHERE name = $8)
		) ON CONFLICT DO NOTHING
	`

	tx := db.MustBegin()

	for _, track := range tracks {
		tx.MustExec(insertArtistStmt, track.Artist, slug.Make(track.Artist.String))

		if track.Album.Name.Valid {
			tx.MustExec(
				insertAlbumStmt,
				track.Album.Name,
				slug.Make(track.Album.Name.String),
				track.Album.Image.Data,
				track.Album.Image.MimeType,
				track.Artist,
			)
		}

		if track.Genre.Valid {
			tx.MustExec(insertGenreStmt, track.Genre, slug.Make(track.Genre.String))
		}

		tx.MustExec(
			insertTrackStmt,
			track.Title,
			track.TrackNumber,
			track.Path,
			track.Date,
			track.Length,
			track.Album.Name,
			track.Artist,
			track.Artist,
			track.Genre,
		)
	}

	err := tx.Commit()
	return err
}

type NotFoundError string

func (e *NotFoundError) Error() string { return "Not Found" }
