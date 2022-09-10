package handlers

import (
	"context"
	"log"
	"orkester/ent"
	"orkester/ent/album"
	"orkester/ent/track"
	"orkester/indexFiles"
	"orkester/models"
	"orkester/repositories"

	"github.com/gofiber/fiber/v2"
)

func UpdateLibrary(client *ent.Client, context context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tracksOnDisk, failed_tracks, err := indexFiles.ScanPathForMusicFiles("/home/qluxzz/Music")
		if err != nil {
			return err
		}

		log.Printf("%d tracks failed to be indexed", len(failed_tracks))

		log.Printf("%d tracks found", len(tracksOnDisk))

		removed_db_tracks, err := repositories.RemoveDeletedEntities(tracksOnDisk, client, context)
		if err != nil {
			return err
		}

		log.Printf("Removed %d tracks", len(removed_db_tracks))

		addedTrackIds, err := repositories.AddTracks(tracksOnDisk, client, context)
		if err != nil {
			return err
		}

		log.Printf("Added %d tracks", len(addedTrackIds))

		tracks, err := client.
			Track.
			Query().
			Where(track.IDIn(addedTrackIds...)).
			WithAlbum(func(aq *ent.AlbumQuery) {
				aq.Select(album.FieldName)
			}).
			WithArtists().
			All(context)

		if err != nil {
			return err
		}

		added_tracks := []models.TrackWithPath{}

		for _, track := range tracks {
			added_tracks = append(added_tracks, models.FromEntTrackWithPath(track))
		}

		removed_tracks := []models.TrackWithPath{}

		for _, track := range removed_db_tracks {
			removed_tracks = append(removed_tracks, models.FromEntTrackWithPath(track))
		}

		return c.JSON(&fiber.Map{
			"tracksRemoved": removed_tracks,
			"tracksAdded":   added_tracks,
			"failedTracks":  failed_tracks,
		})
	}
}
