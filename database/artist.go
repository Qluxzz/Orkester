package database

import (
	"github.com/jmoiron/sqlx"
)

type artistAlbum struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	UrlName string `json:"urlName"`
}

type artist struct {
	Id      int           `db:"id" json:"id"`
	Name    string        `db:"name" json:"name"`
	UrlName string        `db:"urlname" json:"urlName"`
	Albums  []artistAlbum `json:"albums"`
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

	err = db.Select(
		&artist.Albums,
		`SELECT DISTINCT
				id,
				name,
				urlname
			FROM
				albums
			WHERE
				id IN (SELECT albumid FROM tracks WHERE artistid = ?)
		`,
		artist.Id,
	)

	if err != nil {
		return nil, err
	}

	return &artist, nil
}
