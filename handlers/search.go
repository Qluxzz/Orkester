package handlers

import (
	"context"
	"net/url"
	"orkester/ent"
	"orkester/ent/album"
	"orkester/ent/artist"
	"orkester/ent/track"

	"github.com/gofiber/fiber/v2"
)

func Search(client *ent.Client, context context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		query, err := url.QueryUnescape(c.Params("query"))
		if err != nil {
			return err
		}

		tracks := []struct {
			Id    int    `json:"id"`
			Title string `json:"title"`
		}{}

		err = client.
			Track.
			Query().
			Where(track.TitleContainsFold(query)).
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
			Where(album.NameContainsFold(query)).
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
			Where(artist.NameContainsFold(query)).
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
