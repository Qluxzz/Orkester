package main

import (
	"context"
	"log"

	"goreact/ent"
	"goreact/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	_ "github.com/mattn/go-sqlite3"
)

var mode string

func main() {
	client, err := ent.Open("sqlite3", "sqlite.db?cache=shared&_fk=1")

	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	ctx := context.Background()

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	app := fiber.New()

	// Middlewares

	app.Use(logger.New())

	app.Use(cors.New())

	// Routes

	v1 := app.Group("/api/v1")

	track := v1.Group("/track")
	track.Get("/:id/stream", handlers.TrackStream(client, ctx))
	track.Get("/:id", handlers.TrackInfo(client, ctx))
	track.Put("/:id/like", handlers.LikeTrack(client, ctx))
	track.Delete("/:id/like", handlers.UnLikeTrack(client, ctx))

	album := v1.Group("/album")
	album.Get("/:id", handlers.GetAlbum(client, ctx))
	album.Get("/:id/image", handlers.GetAlbumCover(client, ctx))

	artist := v1.Group("/artist")
	artist.Get("/:id", handlers.GetArtist(client, ctx))

	search := v1.Group("/search")
	search.Get("/:query", handlers.Search(client, ctx))

	playlist := v1.Group("/playlist")
	playlist.Get("/liked", handlers.GetLikedTracks(client, ctx))

	v1.Put("/scan", handlers.UpdateLibrary(client, ctx))

	if mode == "production" {
		app.Static("/", "client/")

		app.Get("/*", func(c *fiber.Ctx) error {
			log.Print("Tried to access /")
			return c.SendFile("client/index.html")
		})
	} else {
		// Used for end-to-end testing
		v1.Put("/scan/fake", handlers.AddFakeTracks(client, ctx))
	}

	// Start app

	log.Fatalln(app.Listen(":42000"))
}
