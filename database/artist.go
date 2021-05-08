package database

import (
	"goreact/models"

	"github.com/jmoiron/sqlx"
)

type ArtistAlbum struct {
	Id      int            `json:"id"`
	Name    string         `json:"name"`
	UrlName string         `json:"urlName"`
	Tracks  []models.Track `json:"tracks"`
}

type Artist struct {
	Name   string
	Albums []ArtistAlbum
}

func GetArtistById(artistId int, db *sqlx.DB) (*Artist, error) {
	var artistName string

	err := db.Get(
		&artistName,
		"SELECT name FROM artists WHERE id = ?",
		artistId,
	)

	if err != nil {
		return nil, err
	}

	albums := []ArtistAlbum{}

	err = db.Select(
		&albums,
		`SELECT
				id,
				name,
				urlname
			FROM
				albums
			WHERE
				artistid = ?
			`,
		artistId,
	)

	if err != nil {
		return nil, err
	}

	return &Artist{
		Name:   artistName,
		Albums: albums,
	}, nil
}
