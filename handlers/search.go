package handlers

import (
	"goreact/models"
	"net/url"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

// Takes HTML encoded query and returns formatted query to be used as search query
// Example: Tom%20%PETTY -> tompetty
func formatQuery(query string) (string, error) {
	query, err := url.QueryUnescape(query)
	if err != nil {
		return "", err
	}
	query = strings.ReplaceAll(
		strings.TrimSpace(strings.ToLower(query)),
		" ",
		"",
	)
	return query, nil
}
func Search(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		query, err := formatQuery(c.Params("query"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		wildcardQuery := "%" + query + "%"

		tracks := []models.IdNameAndUrlName{}
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
				OR EXISTS(SELECT * FROM artists WHERE id = artistid AND LOWER(REPLACE(name, ' ', '')) LIKE $1)
			`,
			wildcardQuery,
		)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		albums := []models.IdNameAndUrlName{}
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
			`,
			wildcardQuery,
		)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		artists := []models.IdNameAndUrlName{}
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
			`,
			wildcardQuery,
		)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.JSON(&fiber.Map{
			"tracks":  tracks,
			"albums":  albums,
			"artists": artists,
		})
	}
}
