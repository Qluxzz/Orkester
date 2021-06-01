package handlers

import (
	"goreact/database"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func TracksInfo(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ids := new([]int)

		if err := c.BodyParser(ids); err != nil {
			return err
		}

		tracks, err := database.GetTracksByIds(*ids, db)
		if err != nil {
			return nil
		}

		return c.JSON(tracks)
	}
}

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

		pathAndMimeType, err := database.GetTrackPath(id, db)
		if err != nil {
			return err
		}

		stream, err := os.Open(pathAndMimeType.Path)
		if err != nil {
			return err
		}

		c.Response().Header.Add("content-type", pathAndMimeType.MimeType)
		return c.SendStream(stream)
	}
}

func LikeTrack(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		database.LikeTrack(id, db)

		return c.SendStatus(fiber.StatusOK)
	}
}

func UnLikeTrack(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		database.UnlikeTrack(id, db)

		return c.SendStatus(fiber.StatusOK)
	}
}
