package handlers

import (
	"goreact/models"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type ArtistAlbum struct {
	Id      int            `json:"id"`
	Name    string         `json:"name"`
	UrlName string         `json:"urlName"`
	Tracks  []models.Track `json:"tracks"`
}

func GetArtist(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		var artistName string

		err := db.Get(
			&artistName,
			"SELECT name FROM artists WHERE id = ?",
			id,
		)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
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
			id,
		)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.JSON(&fiber.Map{
			"name":   artistName,
			"albums": albums,
		})
	}
}
