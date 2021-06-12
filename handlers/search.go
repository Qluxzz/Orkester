package handlers

import (
	"context"
	"goreact/ent"
	"goreact/ent/album"
	"goreact/ent/artist"
	"goreact/ent/track"

	"github.com/gofiber/fiber/v2"
)

func Search(client *ent.Client, context context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		query := c.Params("query")

		tracks := []struct {
			Id    int    `json:"id"`
			Title string `json:"title"`
		}{}

		err := client.
			Track.
			Query().
			Where(track.TitleContains(query)).
			Select(track.FieldID, track.FieldTitle).
			Scan(context, &tracks)

		if err != nil {
			return err
		}

		albums := []struct {
			Id      int    `json:"id"`
			Name    string `json:"name"`
			URLName string `json:"urlName" sql:"url_name"`
		}{}

		err = client.
			Album.
			Query().
			Where(album.NameContains(query)).
			Select(album.FieldID, album.FieldName, album.FieldURLName).
			Scan(context, &albums)

		if err != nil {
			return err
		}

		artists := []struct {
			Id      int    `json:"id"`
			Name    string `json:"name"`
			URLName string `json:"urlName"  sql:"url_name"`
		}{}

		err = client.
			Artist.
			Query().
			Where(artist.NameContains(query)).
			Select(artist.FieldID, artist.FieldName, artist.FieldURLName).
			Scan(context, &artists)

		if err != nil {
			return err
		}

		return c.JSON(&fiber.Map{
			"tracks":  tracks,
			"albums":  albums,
			"artists": artists,
		})
	}
}
