package database

import (
	"goreact/models"
	"sort"

	"github.com/jmoiron/sqlx"
)

type Album struct {
	Name   string
	Tracks []models.Track
}

func GetAlbum(albumId int, db *sqlx.DB) (*Album, error) {
	var albumName string

	err := db.Get(
		&albumName,
		`
			SELECT
				name
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
		Name:   albumName,
		Tracks: tracks,
	}, nil
}

type AlbumImage struct {
	Image    []byte `db:"image"`
	MimeType string `db:"imagemimetype"`
}

func GetAlbumCover(albumId int, db *sqlx.DB) (*AlbumImage, error) {
	image := AlbumImage{}

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
