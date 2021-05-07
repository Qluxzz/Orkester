package handlers

import (
	"goreact/repositories"
	"sort"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func GetAlbum(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		var albumName string

		err := db.Get(
			&albumName,
			`
			SELECT
				name
			FROM 
				albums
			WHERE
				id = ?
			`,
			id,
		)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		trackIds := []int{}

		err = db.Select(
			&trackIds,
			`SELECT
				id
			FROM
				tracks
			WHERE
				albumid = ?
			`,
			id,
		)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		tracks, err := repositories.GetTracksByIds(trackIds, db)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		// Sort by track number ascending
		sort.SliceStable(tracks, func(i int, j int) bool { return tracks[i].TrackNumber < tracks[j].TrackNumber })

		return c.JSON(&fiber.Map{
			"name":   albumName,
			"tracks": tracks,
		})
	}
}

func GetAlbumCover(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		type albumImage struct {
			Image    []byte `db:"image"`
			MimeType string `db:"imagemimetype"`
		}

		image := albumImage{}

		err = db.Get(
			&image,
			`SELECT
				image,
				imagemimetype
			FROM
				albums
			WHERE
				id = ?`,
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
