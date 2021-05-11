package database

import (
	"goreact/models"
	"sort"

	"github.com/jmoiron/sqlx"
)

type album struct {
	Id      int            `db:"id"      json:"id"`
	Name    string         `db:"name"    json:"name"`
	UrlName string         `db:"urlName" json:"urlName"`
	Tracks  []models.Track `db:"tracks"  json:"tracks"`
}

func GetAlbum(albumId int, db *sqlx.DB) (*album, error) {
	album := album{}

	err := db.Get(
		&album,
		`
			SELECT
				id,
				name,
				urlname
			FROM 
				albums
			WHERE
				id = ?
			`,
		albumId,
	)

	if err != nil {
		return nil, err
	}

	trackIds := []int{}

	err = db.Select(
		&trackIds,
		`SELECT
				id
			FROM
				tracks
			WHERE
				albumid = ?
			`,
		albumId,
	)

	if err != nil {
		return nil, err
	}

	tracks, err := GetTracksByIds(trackIds, db)
	if err != nil {
		return nil, err
	}

	// Sort by track number ascending
	sort.SliceStable(tracks, func(i int, j int) bool { return tracks[i].TrackNumber < tracks[j].TrackNumber })

	album.Tracks = tracks

	return &album, nil
}

type albumImage struct {
	Image    []byte `db:"image"`
	MimeType string `db:"imagemimetype"`
}

func GetAlbumCover(albumId int, db *sqlx.DB) (*albumImage, error) {
	image := albumImage{}

	err := db.Get(
		&image,
		`SELECT
			image,
			imagemimetype
		FROM
			albums
		WHERE
			id = ?`,
		albumId,
	)
	if err != nil {
		return nil, err
	}

	return &image, nil
}
