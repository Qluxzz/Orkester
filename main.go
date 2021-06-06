package main

import (
	"context"
	"log"

	"goreact/ent"
	"goreact/handlers"
	"goreact/indexFiles"
	"goreact/repositories"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	_ "github.com/mattn/go-sqlite3"
)

func scanAndAddTracksToDb(client *ent.Client, ctx context.Context) {
	tracks, err := indexFiles.ScanPathForMusicFiles("./content")
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%d tracks found", len(tracks))

	repositories.AddTracks(tracks, client, ctx)

	log.Print("Tracks has been added")
}

func main() {
	client, err := ent.Open("sqlite3", "file::memory:?cache=shared&_fk=1")

	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	ctx := context.Background()

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	go scanAndAddTracksToDb(client, ctx)

	app := fiber.New()

	// Middlewares

	app.Use(logger.New())

	// Routes

	v1 := app.Group("/api/v1")

	// track := v1.Group("/track")
	// track.Get("/:id/stream", handlers.TrackStream(db))
	// track.Get("/:id", handlers.TrackInfo(db))
	// track.Get("/:id/like", handlers.LikeTrack(db))
	// track.Get("/:id/unlike", handlers.UnLikeTrack(db))

	// v1.Post("/tracks/ids", handlers.TracksInfo(db))

	album := v1.Group("/album")
	album.Get("/:id", handlers.GetAlbum(ctx, client))
	album.Get("/:id/image", handlers.GetAlbumCover(ctx, client))

	// artist := v1.Group("/artist")
	// artist.Get("/:id", handlers.GetArtist(db))

	// search := v1.Group("/search")
	// search.Get("/:query", handlers.Search(db))

	// playlist := v1.Group("/playlist")
	// playlist.Get("/liked", handlers.GetLikedTracks(db))

	// app.Static("/", "web/build")

	// app.Get("/*", func(c *fiber.Ctx) error {
	// 	return c.SendFile("./web/build/index.html")
	// })

	// Start app

	log.Fatalln(app.Listen(":42000"))
}
