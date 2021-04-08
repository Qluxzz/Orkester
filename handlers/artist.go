package handlers

import (
	"goreact/models"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type Artist struct {
	Name   string                    `json:"name"`
	Albums []models.IdNameAndUrlName `json:"albums"`
}

func GetArtist(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		artist := Artist{}

		err := db.Get(
			&artist,
			`
				SELECT
					name
				FROM
					artists
				WHERE
					id = ?
			`,
			id,
		)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		albums := []models.IdNameAndUrlName{}

		err = db.Select(
			&albums,
			`SELECT id, name, urlname FROM albums WHERE id IN (SELECT albumid FROM tracks WHERE artistid = ?)`,
			id,
		)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		artist.Albums = albums

		return c.JSON(artist)
	}
}
