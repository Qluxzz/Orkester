package handlers

import (
	"goreact/models"
	"goreact/repositories"
	"sort"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type Album struct {
	Name   string         `json:"name"`
	Year   string         `json:"year"`
	Tracks []models.Track `json:"tracks"`
}

func GetAlbum(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		album := Album{}

		err := db.Get(
			&album,
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

		album.Tracks = tracks

		return c.JSON(album)
	}
}
