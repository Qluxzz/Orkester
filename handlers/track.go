package handlers

import (
	"context"
	"goreact/ent"
	"goreact/ent/track"
	"goreact/models"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func TracksInfo(client *ent.Client, context context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ids := new([]int)

		if err := c.BodyParser(ids); err != nil {
			return err
		}

		tracks, err := client.
			Track.
			Query().
			Where(track.IDIn(*ids...)).
			WithAlbum().
			WithArtists().
			All(context)

		if err != nil {
			return nil
		}

		t2 := []models.Track{}
		for _, track := range tracks {
			t := models.Track{
				Id:          track.ID,
				TrackNumber: track.TrackNumber,
				Title:       track.Title,
				Length:      track.Length,
				LikeStatus:  "unliked",
				Album: &models.Album{
					Id:      track.Edges.Album.ID,
					Name:    track.Edges.Album.Name,
					UrlName: track.Edges.Album.URLName,
				},
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

			t2 = append(t2, t)
		}

		return c.JSON(t2)
	}
}

func TrackInfo(client *ent.Client, context context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		track, err := client.Track.
			Query().
			Where(track.ID(id)).
			WithAlbum().
			WithArtists().
			Only(context)

		if err != nil {
			return err
		}

		return c.JSON(track)
	}
}

func TrackStream(client *ent.Client, context context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		pathAndMimeType, err := client.Track.Query().Where(track.ID(id)).Select(track.FieldPath, track.FieldMimetype).Only(context)
		if err != nil {
			return err
		}

		stream, err := os.Open(pathAndMimeType.Path)
		if err != nil {
			return err
		}

		c.Response().Header.Add("content-type", pathAndMimeType.Mimetype)
		return c.SendStream(stream)
	}
}

/*
func LikeTrack(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		database.LikeTrack(id, db)

		return c.SendStatus(fiber.StatusOK)
	}
}

func UnLikeTrack(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		database.UnlikeTrack(id, db)

		return c.SendStatus(fiber.StatusOK)
	}
}
*/
