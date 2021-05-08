package handlers

import (
	"goreact/database"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Search(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		searchResults, err := database.Search(c.Params("query"), db)
		if err != nil {
			return err
		}

		return c.JSON(&fiber.Map{
			"tracks":  searchResults.Tracks,
			"albums":  searchResults.Albums,
			"artists": searchResults.Artists,
		})
	}
}
