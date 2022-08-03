package handlers

import (
	"context"
	"goreact/ent"
	"goreact/ent/track"
	"goreact/models"
	"sort"

	"github.com/gofiber/fiber/v2"
)

func GetLikedTracks(client *ent.Client, context context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		dbTracks, err := client.
			Track.
			Query().
			Where(track.HasLiked()).
			WithAlbum().
			WithArtists().
			WithLiked().
			All(context)

		if err != nil {
			return err
		}

		sort.SliceStable(dbTracks, func(i, j int) bool {
			return dbTracks[i].Edges.Liked.DateAdded.After(dbTracks[j].Edges.Liked.DateAdded)
		})

		tracks := []models.Track{}

		for _, track := range dbTracks {
			t := models.Track{
				Id:          track.ID,
				TrackNumber: track.TrackNumber,
				Title:       track.Title,
				Length:      track.Length,
				Liked:       true,
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
			t.Album = &models.Album{
				Id:      track.Edges.Album.ID,
				Name:    track.Edges.Album.Name,
				UrlName: track.Edges.Album.URLName,
			}

			tracks = append(tracks, t)
		}

		return c.JSON(tracks)
	}
}
