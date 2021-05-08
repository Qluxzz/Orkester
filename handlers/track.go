package handlers

import (
	"goreact/database"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func TrackInfo(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		track, err := database.GetTrackById(id, db)
		if err != nil {
			return err
		}

		return c.JSON(track)
	}
}

func TrackStream(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		path, err := database.GetTrackPath(id, db)
		if err != nil {
			return err
		}

		stream, err := os.Open(*path)
		if err != nil {
			return err
		}

		c.Response().Header.Add("content-type", "audio/flac")
		return c.SendStream(stream)
	}
}
