package database

import (
	"goreact/models"
	"sort"

	"github.com/jmoiron/sqlx"
)

type Album struct {
	Id      string
	Name    string
	UrlName string
	Tracks  []models.Track
}

func GetAlbum(albumId int, db *sqlx.DB) (*Album, error) {

	type NameAndUrlName struct {
		Id      int    `db:"id"`
		Name    string `db:"name"`
		UrlName string `db:"urlname"`
	}

	nameAndUrlName := NameAndUrlName{}

	err := db.Get(
		&nameAndUrlName,
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

	return &Album{
		Name:    nameAndUrlName.Name,
		UrlName: nameAndUrlName.UrlName,
		Tracks:  tracks,
	}, nil
}

func GetAlbumCover(albumId int, db *sqlx.DB) (*models.AlbumImage, error) {
	image := models.AlbumImage{}

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
