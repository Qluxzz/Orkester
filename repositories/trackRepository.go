package repositories

import (
	"goreact/indexFiles"
	"goreact/models"

	"github.com/gosimple/slug"
	"github.com/jmoiron/sqlx"
)

func AddTracks(tracks []*indexFiles.IndexedTrack, db *sqlx.DB) error {
	insertArtistStmt := `
		INSERT INTO artists (name, urlname) VALUES (?, ?) ON CONFLICT(name) DO NOTHING
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
		) ON CONFLICT (name, artistid) DO NOTHING
	`

	insertGenreStmt := `
		INSERT INTO genres (name, urlname) VALUES (?, ?) ON CONFLICT(name) DO NOTHING
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
		)
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

func GetTracksByIds(ids []int, db *sqlx.DB) ([]models.Track, error) {
	query, args, err := sqlx.In(`
			SELECT
				t.id,
				t.title,
				t.tracknumber,
				t.date,
				t.length,
				albums.id albumid,
				albums.name albumname,
				albums.urlname albumurlname,
				artists.id artistid,
				artists.name artistname,
				artists.urlname artisturlname,
				genres.id genreid,
				genres.name genrename,
				genres.urlname genreurlname
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

func GetTrackById(id int, db *sqlx.DB) (models.Track, error) {
	tracks, err := GetTracksByIds([]int{id}, db)
	if err != nil {
		return models.Track{}, err
	}

	return tracks[0], nil
}
