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
	Name    string `db:"name"`
	UrlName string `db:"urlname"`
	Albums  []ArtistAlbum
}

func GetArtistById(artistId int, db *sqlx.DB) (*Artist, error) {
	artist := Artist{}

	err := db.Get(
		&artist,
		"SELECT name, urlname FROM artists WHERE id = ?",
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
		Name:    artist.Name,
		UrlName: artist.UrlName,
		Albums:  albums,
	}, nil
}
