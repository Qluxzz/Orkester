package handlers

import (
	"context"
	"goreact/ent"
	"goreact/ent/album"
	"goreact/ent/track"
	"goreact/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetLikedTracks(client *ent.Client, context context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		dbTracks, err := client.
			Track.
			Query().
			Where(track.HasLiked()).
			WithAlbum(func(aq *ent.AlbumQuery) {
				aq.Select(album.FieldID, album.FieldName, album.FieldURLName).Only(context)
			}).
			WithArtists().
			WithLiked().
			All(context)

		tracks := []models.Track{}
		for _, track := range dbTracks {
			liked, err := track.Edges.LikedOrErr()
			if err != nil {
				log.Fatalf("Liked track was not liked! %v", err)
			}

			t := models.Track{
				Id:          track.ID,
				TrackNumber: track.TrackNumber,
				Title:       track.Title,
				Length:      track.Length,
				Liked:       true,
				Date:        liked.DateAdded,
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

		if err != nil {
			return err
		}

		return c.JSON(tracks)
	}
}
