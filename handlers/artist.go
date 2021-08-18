package handlers

import (
	"context"
	"goreact/ent"
	"goreact/ent/artist"
	"goreact/models"
	"sort"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetArtist(client *ent.Client, context context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		artistInfo, err := client.
			Artist.
			Query().
			Where(artist.ID(id)).
			WithAlbums().                         // Albums
			WithTracks(func(tq *ent.TrackQuery) { // Appears on
				tq.WithAlbum()
			}).
			Only(context)

		if err != nil {
			return err
		}

		deduplicatedAlbums := make(map[int]models.Album)

		for _, album := range artistInfo.Edges.Albums {
			deduplicatedAlbums[album.ID] = models.Album{
				Id:      album.ID,
				Name:    album.Name,
				UrlName: album.URLName,
			}
		}

		for _, track := range artistInfo.Edges.Tracks {
			album := track.Edges.Album

			deduplicatedAlbums[album.ID] = models.Album{
				Id:      album.ID,
				Name:    album.Name,
				UrlName: album.URLName,
			}
		}

		albums := []models.Album{}

		for _, album := range deduplicatedAlbums {
			albums = append(albums, album)
		}

		sort.SliceStable(albums, func(i, j int) bool {
			return albums[i].Id < albums[j].Id
		})

		return c.JSON(&fiber.Map{
			"id":      artistInfo.ID,
			"name":    artistInfo.Name,
			"urlName": artistInfo.URLName,
			"albums":  albums,
		})
	}
}
