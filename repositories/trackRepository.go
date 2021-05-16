package repositories

import (
	"database/sql"
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
			mimetype,
			albumid,
			artistid,
			genreid
		) VALUES(
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			(
				SELECT
					id
				FROM
					albums
				WHERE
					name = $7
					AND artistid = (
						SELECT id FROM artists WHERE name = $8
					)
			),
			(SELECT id FROM artists WHERE name = $9),
			(SELECT id from genres WHERE name = $10)
		) ON CONFLICT DO NOTHING
	`

	tx := db.MustBegin()

	for _, track := range tracks {
		tx.MustExec(insertArtistStmt, track.ArtistName, slug.Make(track.ArtistName.String))

		if track.AlbumArtist.Valid {
			tx.MustExec(insertArtistStmt, track.AlbumArtist, slug.Make(track.AlbumArtist.String))
		}

		var albumArtist sql.NullString

		if track.AlbumArtist.Valid {
			albumArtist = track.AlbumArtist
		} else {
			albumArtist = track.ArtistName
		}

		if track.AlbumName.Valid {
			tx.MustExec(
				insertAlbumStmt,
				track.AlbumName,
				slug.Make(track.AlbumName.String),
				track.Image.Data,
				track.Image.MimeType,
				albumArtist,
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
			track.MimeType,
			track.AlbumName,
			albumArtist,
			track.ArtistName,
			track.Genre,
		)
	}

	err := tx.Commit()
	return err
}
