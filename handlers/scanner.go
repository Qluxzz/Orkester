package handlers

import (
	"context"
	"fmt"
	"log"
	"orkester/ent"
	"orkester/ent/album"
	"orkester/ent/track"
	"orkester/indexFiles"
	"orkester/models"
	"orkester/repositories"
	"os"

	"github.com/gofiber/fiber/v2"
)

func AddSearchPath(client *ent.Client, context context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		type Body struct {
			Path string `json:"path"`
		}

		body := new(Body)

		err := c.BodyParser(&body)

		if err != nil {
			return err
		}

		_, err = os.Stat(body.Path)
		if os.IsNotExist(err) {
			return c.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("Folder %s does not exist", body.Path))
		}

		_, err = client.SearchPath.Create().SetPath(body.Path).Save(context)
		if err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}

func UpdateLibrary(client *ent.Client, context context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		searchPaths, err := client.SearchPath.Query().All(context)
		if err != nil {
			return err
		}

		if len(searchPaths) == 0 {
			return c.Status(fiber.StatusBadRequest).SendString("No search paths added, please add a search path before calling this endpoint")
		}

		uniqueTracks := make(map[string]*indexFiles.IndexedTrack)

		failedTracks := []*indexFiles.FailedAudioFile{}

		for _, path := range searchPaths {
			tracksOnDisk, failed_tracks, err := indexFiles.ScanPathForMusicFiles(path.Path)
			if err != nil {
				return err
			}

			log.Printf("%d tracks failed to be indexed", len(failed_tracks))

			log.Printf("%d tracks found", len(tracksOnDisk))

			for _, track := range tracksOnDisk {
				uniqueTracks[track.Path] = track
			}

			failedTracks = append(failedTracks, failed_tracks...)
		}

		uniqueTracksList := []*indexFiles.IndexedTrack{}
		for _, track := range uniqueTracks {
			uniqueTracksList = append(uniqueTracksList, track)
		}

		removed_db_tracks, err := repositories.RemoveDeletedEntities(uniqueTracksList, client, context)
		if err != nil {
			return err
		}

		log.Printf("Removed %d tracks", len(removed_db_tracks))

		addedTrackIds, err := repositories.AddTracks(uniqueTracksList, client, context)
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

		addedTracks := []models.TrackWithPath{}

		for _, track := range tracks {
			addedTracks = append(addedTracks, models.FromEntTrackWithPath(track))
		}

		removedTracks := []models.TrackWithPath{}

		for _, track := range removed_db_tracks {
			removedTracks = append(removedTracks, models.FromEntTrackWithPath(track))
		}

		return c.JSON(&fiber.Map{
			"tracksRemoved": removedTracks,
			"tracksAdded":   addedTracks,
			"failedTracks":  failedTracks,
		})
	}
}
