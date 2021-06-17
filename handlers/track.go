package handlers

import (
	"context"
	"goreact/ent"
	"goreact/ent/likedtrack"
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
			WithLiked().
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
				Liked:       track.Edges.Liked != nil,
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

		pathAndMimeType, err := client.
			Track.
			Query().
			Where(track.ID(id)).
			Select(track.FieldPath, track.FieldMimetype).
			Only(context)

		if err != nil {
			return err
		}

		stream, err := os.Open(pathAndMimeType.Path)
		if err != nil {
			return err
		}
		defer stream.Close()

		c.Response().Header.Add("content-type", pathAndMimeType.Mimetype)
		return c.SendStream(stream)
	}
}

func LikeTrack(client *ent.Client, context context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		_, err = client.
			LikedTrack.
			Create().
			SetTrackID(id).
			Save(context)

		if err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusOK)
	}
}

func UnLikeTrack(client *ent.Client, context context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		_, err = client.LikedTrack.Delete().Where(likedtrack.HasTrackWith(track.ID(id))).Exec(context)

		if err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusOK)
	}
}
