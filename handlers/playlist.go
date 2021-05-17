package handlers

import (
	"goreact/database"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func GetLikedTracks(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		likedTracks, err := database.GetLikedTracks(db)
		if err != nil {
			return err
		}

		return c.JSON(likedTracks)
	}
}
