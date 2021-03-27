package handlers

import (
	"goreact/repositories"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func BrowseArtists(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		genres := []string{}

		err := db.Select(&genres, "SELECT name FROM genres")
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.JSON(genres)
	}
}

func BrowseArtist(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		trackIds := []int{}

		err := db.Select(
			&trackIds,
			`SELECT
			tracks.id
		 FROM tracks
		 INNER JOIN artists
			 ON artists.id = tracks.artistid
		 WHERE
			 artists.urlname = ?
		`, c.Params("artist-url-name"))

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
		trackIds := []int{}

		err := db.Select(
			&trackIds,
			`SELECT
				tracks.id
			 FROM tracks
			 INNER JOIN genres
			 	ON genres.id = tracks.genreid
			 WHERE
			 	genres.urlname = ?
		`, c.Params("genre"))
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
