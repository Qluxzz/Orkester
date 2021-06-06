package handlers

import (
	"context"
	"goreact/ent"
	"goreact/ent/album"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAlbum(ctx context.Context, client *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		album, err := client.
			Album.
			Query().
			Where(album.ID(id)).
			WithArtist().
			WithTracks(func(q *ent.TrackQuery) {
				q.WithArtists()
			}).
			Only(ctx)
		if err != nil {
			return err
		}

		return c.JSON(&fiber.Map{
			"id":      album.ID,
			"name":    album.Name,
			"urlName": album.URLName,
			"tracks":  album.Edges.Tracks,
			"artist":  album.Edges.Artist,
		})
	}
}

func GetAlbumCover(ctx context.Context, client *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		albumCover, err := client.Album.Query().Where(album.ID(id)).Only(ctx)
		if err != nil {
			return err
		}

		c.Response().Header.Add("Content-Type", albumCover.ImageMimeType)
		c.Response().Header.Add("Cache-Control", "max-age=31536000")
		return c.Send(albumCover.Image)
	}
}
