package handlers

import (
	"goreact/models"
	"goreact/repositories"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func BrowseArtists(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		artists := []models.IdNameAndUrlName{}

		err := db.Select(
			&artists,
			"SELECT id, name, urlname FROM artists",
		)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.JSON(artists)
	}
}

func BrowseArtist(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		trackIds := []int{}

		err := db.Select(
			&trackIds,
			`SELECT
				id
			FROM 
				tracks
			WHERE
				artistid = (SELECT id FROM artists WHERE urlname = ?)
		`, c.Params("name"))

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		tracks, err := repositories.GetTracksByIds(trackIds, db)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.JSON(tracks)
	}
}

func BrowseGenres(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		genres := []models.IdNameAndUrlName{}

		err := db.Select(
			&genres,
			"SELECT name, urlname FROM genres",
		)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.JSON(genres)
	}
}

func BrowseGenre(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		trackIds := []int{}

		err := db.Select(
			&trackIds,
			`SELECT
				id
			FROM
				tracks
			WHERE
				genreid = (SELECT id FROM genres WHERE urlname = ?)
		`, c.Params("name"))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		tracks, err := repositories.GetTracksByIds(trackIds, db)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.JSON(tracks)
	}
}
