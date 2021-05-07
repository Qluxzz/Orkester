package handlers

import (
	"goreact/repositories"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func TrackInfo(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		track, err := repositories.GetTrackById(id, db)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.JSON(track)
	}
}

func TrackImage(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		type trackImage struct {
			Image    []byte `db:"image"`
			MimeType string `db:"imagemimetype"`
		}

		image := trackImage{}

		err = db.Get(
			&image,
			`SELECT
				image,
				imagemimetype
			FROM
				albums
			WHERE
				id = (
					SELECT
						albumid
					FROM
						tracks
					WHERE
						id = ?
				)`,
			id,
		)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		c.Response().Header.Add("Content-Type", image.MimeType)
		c.Response().Header.Add("Cache-Control", "max-age=31536000")
		return c.Send(image.Image)
	}
}

func TrackStream(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		var path string
		err = db.Get(&path, "SELECT path FROM tracks WHERE id=?", id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		c.Response().Header.Add("content-type", "audio/flac")

		stream, err := os.Open(path)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.SendStream(stream)
	}
}
