package handlers

import (
	"context"
	"goreact/ent"
	"goreact/ent/album"
	"goreact/ent/artist"
	"goreact/ent/track"
	"goreact/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetArtist(client *ent.Client, context context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		artist_info, err := client.
			Artist.
			Query().
			Where(artist.ID(id)).
			Only(context)

		if err != nil {
			return err
		}

		dbAlbums, err := client.
			Album.
			Query().
			Where(album.Or(
				album.HasArtistWith(artist.ID(id)),
				album.HasTracksWith(track.HasArtistsWith(artist.ID(id))),
			)).
			Select(album.FieldID, album.FieldName, album.FieldURLName).
			All(context)

		if err != nil {
			return nil
		}

		albums := []models.Album{}

		for _, album := range dbAlbums {
			albums = append(albums, models.Album{
				Id:      album.ID,
				Name:    album.Name,
				UrlName: album.URLName,
			})
		}

		return c.JSON(&fiber.Map{
			"id":      artist_info.ID,
			"name":    artist_info.Name,
			"urlName": artist_info.URLName,
			"albums":  albums,
		})
	}
}
