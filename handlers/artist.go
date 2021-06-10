package handlers

import (
	"context"
	"goreact/ent"
	"goreact/ent/album"
	"goreact/ent/artist"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetArtist(client *ent.Client, context context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		artist, err := client.
			Artist.
			Query().
			Where(artist.ID(id)).
			WithAlbums(func(aq *ent.AlbumQuery) {
				aq.Select(album.FieldID, album.FieldName, album.FieldURLName)
			}).
			Only(context)

		if err != nil {
			return err
		}

		return c.JSON(&fiber.Map{
			"id":      artist.ID,
			"name":    artist.Name,
			"urlName": artist.URLName,
			"albums":  artist.Edges.Albums,
		})
	}
}
