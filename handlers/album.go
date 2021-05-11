package handlers

import (
	"goreact/database"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func GetAlbum(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		album, err := database.GetAlbum(id, db)
		if err != nil {
			return err
		}

		return c.JSON(album)
	}
}

func GetAlbumCover(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		albumCover, err := database.GetAlbumCover(id, db)
		if err != nil {
			return err
		}

		c.Response().Header.Add("Content-Type", albumCover.MimeType)
		c.Response().Header.Add("Cache-Control", "max-age=31536000")
		return c.Send(albumCover.Image)
	}
}
