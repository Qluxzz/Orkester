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
			(SELECT id from genres WHERE name = $9)
		) ON CONFLICT DO NOTHING
	`

	insertTrackArtist := `
		INSERT INTO
			trackArtists (
				trackid,
				artistid
			) VALUES (
				?,
				(SELECT id FROM artists WHERE name = ?)
			)`

	for _, track := range tracks {
		tx := db.MustBegin()

		for _, artist := range track.Artists {
			tx.MustExec(insertArtistStmt, artist, slug.Make(artist.String))
		}
		if track.AlbumArtist.Valid {
			tx.MustExec(insertArtistStmt, track.AlbumArtist, slug.Make(track.AlbumArtist.String))
		}

		var albumArtist sql.NullString
		if track.AlbumArtist.Valid {
			albumArtist = track.AlbumArtist
		} else {
			albumArtist = track.Artists[0]
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

		result := tx.MustExec(
			insertTrackStmt,
			track.Title,
			track.TrackNumber,
			track.Path,
			track.Date,
			track.Length,
			track.MimeType,
			track.AlbumName,
			albumArtist,
			track.Genre,
		)

		// When SQLite is updated so I can use RETURNING
		// this can be greatly improved
		rowsAffected, err := result.RowsAffected()
		if err != nil || rowsAffected == 0 {
			tx.Rollback()
			continue
		}

		id, err := result.LastInsertId()
		if err != nil {
			tx.Rollback()
			continue
		}

		for _, artist := range track.Artists {
			tx.MustExec(
				insertTrackArtist,
				id,
				artist,
			)
		}

		err = tx.Commit()
		if err != nil {
			return err
		}
	}

	return nil
}
