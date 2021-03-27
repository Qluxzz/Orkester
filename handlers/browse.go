package handlers

import (
	"goreact/repositories"

	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
	"github.com/jmoiron/sqlx"
)

type NameAndUrlName struct {
	name    string
	urlname string
}

func BrowseArtists(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		artists := []string{}

		err := db.Select(
			&artists,
			"SELECT name FROM artists",
		)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		artistsAndUrlnames := []NameAndUrlName{}

		for _, artist := range artists {
			artistsAndUrlnames = append(
				artistsAndUrlnames,
				NameAndUrlName{
					name:    artist,
					urlname: slug.Make(artist),
				},
			)
		}

		return c.JSON(artistsAndUrlnames)
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
		genres := []string{}

		err := db.Select(
			&genres,
			"SELECT name FROM genres",
		)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		genresAndUrlnames := []NameAndUrlName{}

		for _, genre := range genres {
			genresAndUrlnames = append(
				genresAndUrlnames,
				NameAndUrlName{
					name:    genre,
					urlname: slug.Make(genre),
				},
			)
		}

		return c.JSON(genresAndUrlnames)
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
				genreid = (SELECT id FROM genre WHERE urlname = ?)
		`, c.Params("genre-url-name"))
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
