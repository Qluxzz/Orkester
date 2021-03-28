package handlers

import (
	"goreact/models"
	"goreact/repositories"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type Album struct {
	Name   string         `json:"name"`
	Year   string         `json:"year"`
	Tracks []models.Track `json:"tracks"`
}

func GetAlbum(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		album := Album{}

		err := db.Get(
			&album,
			`
			SELECT
				name
			FROM 
				albums
			WHERE
				id = ?
			`,
			id,
		)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
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
			id,
		)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		tracks, err := repositories.GetTracksByIds(trackIds, db)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		album.Tracks = tracks

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.JSON(album)
	}
}
