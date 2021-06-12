package handlers

import (
	"context"
	"fmt"
	"goreact/ent"
	"goreact/ent/album"
	"goreact/models"
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
			Select(album.FieldID, album.FieldName, album.FieldURLName).
			Where(album.ID(id)).
			WithArtist().
			WithTracks(func(q *ent.TrackQuery) {
				q.WithArtists()
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
				Liked:       track.Liked,
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

		return c.JSON(&fiber.Map{
			"id":      album.ID,
			"name":    album.Name,
			"urlName": album.URLName,
			"tracks":  tracks,
			"artist":  album.Edges.Artist,
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
			Select(album.FieldImage, album.FieldImageMimeType).
			Only(ctx)

		if err != nil {
			return err
		}

		secondsInAYear := int((24 * time.Hour * 365).Seconds())

		c.Response().Header.Add("Content-Type", albumCover.ImageMimeType)
		c.Response().Header.Add("Cache-Control", fmt.Sprintf("max-age=%d", secondsInAYear))
		return c.Send(albumCover.Image)
	}
}
