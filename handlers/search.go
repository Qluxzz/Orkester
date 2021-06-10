package handlers

import (
	"context"
	"goreact/ent"
	"goreact/ent/album"
	"goreact/ent/artist"
	"goreact/ent/track"

	"github.com/gofiber/fiber/v2"
)

func Search(client *ent.Client, context context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		query := c.Params("query")

		tracks, err := client.Track.Query().Where(track.TitleContains(query)).All(context)
		if err != nil {
			return err
		}

		albums, err := client.Album.Query().Where(album.NameContains(query)).All(context)
		if err != nil {
			return err
		}

		artists, err := client.Artist.Query().Where(artist.NameContains(query)).All(context)
		if err != nil {
			return err
		}

		return c.JSON(&fiber.Map{
			"tracks":  tracks,
			"albums":  albums,
			"artists": artists,
		})
	}
}
