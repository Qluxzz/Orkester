package handlers

import (
	"context"
	"goreact/ent"
	"goreact/indexFiles"
	"goreact/models"
	"goreact/repositories"
	"log"

	"github.com/gofiber/fiber/v2"
)

func UpdateLibrary(client *ent.Client, context context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tracks, err := indexFiles.ScanPathForMusicFiles("/home/qluxzz/Music")
		if err != nil {
			return err
		}

		log.Printf("%d tracks found", len(tracks))

		removed_db_tracks, err := repositories.RemoveDeletedEntities(tracks, client, context)
		if err != nil {
			return err
		}

		log.Printf("Removed %d tracks", len(removed_db_tracks))

		added_db_tracks, err := repositories.AddTracks(tracks, client, context)
		if err != nil {
			return err
		}

		log.Printf("Added %d tracks", len(added_db_tracks))

		added_tracks := []models.TrackWithPath{}

		for _, track := range added_db_tracks {
			added_tracks = append(added_tracks, models.FromEntTrackWithPath(track))
		}

		removed_tracks := []models.TrackWithPath{}

		for _, track := range removed_db_tracks {
			removed_tracks = append(removed_tracks, models.FromEntTrackWithPath(track))
		}

		return c.JSON(&fiber.Map{
			"tracksRemoved": removed_tracks,
			"tracksAdded":   added_tracks,
		})
	}
}
