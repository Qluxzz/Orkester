package handlers

import (
	"context"
	"goreact/ent"
	"goreact/ent/artist"
	"goreact/indexFiles"
	"sort"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AlbumWithReleaseDate struct {
	Id       int                     `json:"id"`
	Name     string                  `json:"name"`
	UrlName  string                  `json:"urlName"`
	Released *indexFiles.ReleaseDate `json:"released"`
}

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

		deduplicatedAlbums := make(map[int]AlbumWithReleaseDate)

		for _, album := range artistInfo.Edges.Albums {

			deduplicatedAlbums[album.ID] = AlbumWithReleaseDate{
				Id:       album.ID,
				Name:     album.Name,
				UrlName:  album.URLName,
				Released: album.Released,
			}
		}

		for _, track := range artistInfo.Edges.Tracks {
			album := track.Edges.Album

			deduplicatedAlbums[album.ID] = AlbumWithReleaseDate{
				Id:       album.ID,
				Name:     album.Name,
				UrlName:  album.URLName,
				Released: album.Released,
			}
		}

		albums := []AlbumWithReleaseDate{}

		for _, album := range deduplicatedAlbums {
			albums = append(albums, album)
		}

		// Maps don't have a specified sort order
		// So we sort the array by id so it's the same between requests
		sort.SliceStable(albums, func(i, j int) bool {
			return albums[i].Released.After(albums[j].Released)
		})

		return c.JSON(&fiber.Map{
			"id":      artistInfo.ID,
			"name":    artistInfo.Name,
			"urlName": artistInfo.URLName,
			"albums":  albums,
		})
	}
}
