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

type artist struct {
	Id      int           `db:"id" json:"id"`
	Name    string        `db:"name" json:"name"`
	UrlName string        `db:"urlname" json:"urlName"`
	Albums  []ArtistAlbum `json:"albums"`
}

func GetArtistById(artistId int, db *sqlx.DB) (*artist, error) {
	artist := artist{}

	err := db.Get(
		&artist,
		"SELECT id, name, urlname FROM artists WHERE id = ?",
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

	artist.Albums = albums

	return &artist, nil
}
