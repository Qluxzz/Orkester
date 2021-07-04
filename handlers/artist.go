package handlers

import (
	"context"
	"goreact/ent"
	"goreact/ent/album"
	"goreact/ent/artist"
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
			WithAlbums(func(aq *ent.AlbumQuery) {
				aq.Select(album.FieldID, album.FieldName, album.FieldName)
			}).
			WithTracks(func(tq *ent.TrackQuery) {
				tq.WithAlbum(func(aq *ent.AlbumQuery) {
					aq.Select(album.FieldID, album.FieldName, album.FieldName)
				})
			}).
			Only(context)

		if err != nil {
			return err
		}

		albums_map := make(map[int]models.Album)

		for _, album := range artist_info.Edges.Albums {

			albums_map[album.ID] = models.Album{
				Id:      album.ID,
				Name:    album.Name,
				UrlName: album.URLName,
			}
		}

		for _, track := range artist_info.Edges.Tracks {
			album := track.Edges.Album
			albums_map[album.ID] = models.Album{
				Id:      album.ID,
				Name:    album.Name,
				UrlName: album.URLName,
			}
		}

		albums := []models.Album{}

		for _, album := range albums_map {
			albums = append(albums, album)
		}

		return c.JSON(&fiber.Map{
			"id":      artist_info.ID,
			"name":    artist_info.Name,
			"urlName": artist_info.URLName,
			"albums":  albums,
		})
	}
}
