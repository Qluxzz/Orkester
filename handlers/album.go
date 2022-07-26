package handlers

import (
	"context"
	"fmt"
	"goreact/ent"
	"goreact/ent/album"
	"goreact/models"
	"sort"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetAlbum(client *ent.Client, ctx context.Context) fiber.Handler {
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
				q.WithLiked()
			}).
			Only(ctx)
		if err != nil {
			return err
		}

		tracks := []models.Track{}
		for _, track := range album.Edges.Tracks {
			t := models.Track{
				Id:          track.ID,
				TrackNumber: track.TrackNumber,
				Title:       track.Title,
				Length:      track.Length,
				Liked:       track.Edges.Liked != nil,
			}

			artists := []*models.Artist{}

			for _, artist := range track.Edges.Artists {
				a := &models.Artist{
					Id:      artist.ID,
					Name:    artist.Name,
					UrlName: artist.URLName,
				}

				artists = append(artists, a)
			}

			t.Artists = artists

			tracks = append(tracks, t)
		}

		sort.SliceStable(tracks, func(i, j int) bool {
			return tracks[i].TrackNumber < tracks[j].TrackNumber
		})

		return c.JSON(&fiber.Map{
			"id":       album.ID,
			"name":     album.Name,
			"urlName":  album.URLName,
			"tracks":   tracks,
			"released": album.Released,
			"artist": &models.Artist{
				Id:      album.Edges.Artist.ID,
				Name:    album.Edges.Artist.Name,
				UrlName: album.Edges.Artist.URLName,
			},
		})
	}
}

func GetAlbumCover(client *ent.Client, ctx context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		albumCover, err := client.
			Album.
			Query().
			Where(album.ID(id)).
			QueryCover().
			Only(ctx)

		if err != nil {
			return err
		}

		const secondsInAYear int = 3600 * 24 * 365

		c.Response().Header.Add("Content-Type", albumCover.ImageMimeType)
		c.Response().Header.Add("Cache-Control", fmt.Sprintf("max-age=%d", secondsInAYear))
		return c.Send(albumCover.Image)
	}
}
