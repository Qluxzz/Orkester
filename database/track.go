package database

import (
	"goreact/models"

	"github.com/jmoiron/sqlx"
)

func GetTracksByIds(ids []int, db *sqlx.DB) ([]models.Track, error) {
	if len(ids) == 0 {
		return []models.Track{}, nil
	}

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
				genres.id genreid,
				genres.name genrename,
				genres.urlname genreurlname,
				IIF(EXISTS(SELECT * FROM likedTracks WHERE trackid = t.id), 1, 0) likeStatus
			FROM
				tracks t
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
		dbArtists := []models.DBArtist{}
		err = db.Select(&dbArtists, "SELECT id, name, urlname FROM artists WHERE id IN (SELECT DISTINCT artistid FROM trackArtists WHERE trackid = ?)", dbTrack.Id)
		if err != nil {
			continue
		}

		tracks = append(tracks, dbTrack.ToDomain(dbArtists))
	}

	return tracks, nil
}

func GetTrackById(id int, db *sqlx.DB) (*models.Track, error) {
	tracks, err := GetTracksByIds([]int{id}, db)
	if err != nil {
		return nil, err
	}

	if len(tracks) == 0 {
		return nil, nil
	}

	return &tracks[0], nil
}

type pathAndMimeType struct {
	Path     string `db:"path"`
	MimeType string `db:"mimetype"`
}

func GetTrackPath(id int, db *sqlx.DB) (*pathAndMimeType, error) {
	data := pathAndMimeType{}

	err := db.Get(&data, "SELECT path, mimetype FROM tracks WHERE id=?", id)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func LikeTrack(id int, db *sqlx.DB) {
	db.Exec("INSERT INTO likedTracks (trackid) VALUES (?)", id)
}

func UnlikeTrack(id int, db *sqlx.DB) {
	db.Exec("DELETE FROM likedTracks WHERE trackid = ?", id)
}
