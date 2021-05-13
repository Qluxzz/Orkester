package handlers

import (
	"goreact/database"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func GetArtist(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		artist, err := database.GetArtistById(id, db)

		if err != nil {
			return err
		}

		return c.JSON(artist)
	}
}
