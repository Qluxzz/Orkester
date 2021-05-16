package database

import (
	"net/url"
	"strings"

	"github.com/jmoiron/sqlx"
)

// Takes HTML encoded query and returns formatted query to be used as search query
// Example: Tom%20%PETTY -> tompetty
func formatQuery(query string) (string, error) {
	query, err := url.QueryUnescape(query)
	if err != nil {
		return "", err
	}
	query = strings.TrimSpace(strings.ToLower(query))
	return query, nil
}

type idNameAndUrlName struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Urlname string `json:"urlName"`
}

type SearchResults struct {
	Tracks  []idNameAndUrlName
	Albums  []idNameAndUrlName
	Artists []idNameAndUrlName
}

func Search(query string, db *sqlx.DB) (*SearchResults, error) {
	query, err := formatQuery(query)
	if err != nil {
		return nil, err
	}

	wildcardQuery := "%" + query + "%"

	tracks := []idNameAndUrlName{}
	err = db.Select(
		&tracks,
		`SELECT
				id,
				title as name
			FROM
				tracks
			WHERE
				LOWER(title) LIKE $1
				OR EXISTS(SELECT * FROM albums WHERE id = albumid AND LOWER(REPLACE(name, ' ', '')) LIKE $1)
				OR EXISTS(
					SELECT
						*
					FROM
						artists
					WHERE
						id IN (
							SELECT
								artistid
							FROM
								trackArtists
							WHERE
								trackid = tracks.id
						) AND LOWER(REPLACE(name, ' ', '')) LIKE $1
				)
			`,
		wildcardQuery,
	)
	if err != nil {
		return nil, err
	}

	albums := []idNameAndUrlName{}
	err = db.Select(
		&albums,
		`SELECT
				id,
				name,
				urlname
			FROM
				albums
			WHERE
				LOWER(REPLACE(name, ' ', '')) LIKE ?
			ORDER BY
				name
			`,
		wildcardQuery,
	)

	if err != nil {
		return nil, err
	}

	artists := []idNameAndUrlName{}
	err = db.Select(
		&artists,
		`SELECT
				id,
				name,
				urlname
			FROM
				artists
			WHERE
				LOWER(REPLACE(name, ' ', '')) LIKE ?
			ORDER BY
				name
			`,
		wildcardQuery,
	)
	if err != nil {
		return nil, err
	}

	return &SearchResults{
		Tracks:  tracks,
		Albums:  albums,
		Artists: artists,
	}, nil
}
