package handlers

import (
	"bufio"
	"context"
	"fmt"
	"orkester/ent"
	"orkester/ent/likedtrack"
	"orkester/ent/track"
	"orkester/models"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func TrackInfo(client *ent.Client, context context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := models.FromTrackId(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		dbTrack, err := client.Track.
			Query().
			Where(track.ID(id)).
			WithAlbum().
			WithArtists().
			WithLiked().
			Only(context)

		if err != nil {
			return err
		}

		return c.JSON(models.FromEntTrack(dbTrack))
	}
}

func TracksInfo(client *ent.Client, context context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ids := c.Query("ids")

		if ids == "" {
			return c.Status(fiber.StatusBadRequest).SendString("No ids were supplied!")
		}

		var trackIds []int

		for _, id := range strings.Split(ids, ",") {
			trackId, err := models.FromTrackId(id)
			if err == nil {
				trackIds = append(trackIds, trackId)
			}
		}

		dbTracks, err := client.
			Track.
			Query().
			Where(track.IDIn(trackIds...)).
			WithAlbum().
			WithArtists().
			WithLiked().
			All(context)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		return c.JSON(models.FromEntTracks(dbTracks))

	}
}

func TrackStream(client *ent.Client, context context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := models.FromTrackId(c.Params("id"))
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

		f, err := os.Open(pathAndMimeType.Path)
		if err != nil {
			f.Close()
			return err
		}

		c.Response().Header.SetContentType(pathAndMimeType.Mimetype)
		c.Response().SetBodyStreamWriter(func(w *bufio.Writer) {
			w.ReadFrom(f)
			w.Flush()
		})

		return nil
	}
}

func TrackImage(client *ent.Client, context context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := models.FromTrackId(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		trackCover, err := client.
			Track.
			Query().
			Where(track.ID(id)).
			QueryImage().
			Only(context)

		if err != nil {
			return err
		}

		const secondsInAYear int = 3600 * 24 * 365

		c.Response().Header.Add("Content-Type", trackCover.ImageMimeType)
		c.Response().Header.Add("Cache-Control", fmt.Sprintf("max-age=%d", secondsInAYear))
		return c.Send(trackCover.Image)
	}
}

func LikeTrack(client *ent.Client, context context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := models.FromTrackId(c.Params("id"))
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
		id, err := models.FromTrackId(c.Params("id"))
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
