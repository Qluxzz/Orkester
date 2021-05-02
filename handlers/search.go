package handlers

import (
	"goreact/models"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Search(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		query := strings.ToLower(strings.TrimSpace(c.Params("query")))
		if len(query) < 4 {
			return c.Status(fiber.StatusBadRequest).SendString("Query has to be atleast 4 characters")
		}

		wildcardQuery := "%" + query + "%"

		tracks := []models.IdNameAndUrlName{}
		err := db.Select(
			&tracks,
			`SELECT id, title as name FROM tracks WHERE LOWER(title) LIKE ?`,
			wildcardQuery,
		)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		albums := []models.IdNameAndUrlName{}
		err = db.Select(
			&albums,
			`SELECT id, name, urlname FROM albums WHERE LOWER(name) LIKE ?`,
			wildcardQuery,
		)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		artists := []models.IdNameAndUrlName{}
		err = db.Select(
			&artists,
			`SELECT id, name, urlname FROM artists WHERE LOWER(name) LIKE ?`,
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
