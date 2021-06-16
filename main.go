package main

import (
	"context"
	"log"
	"time"

	"goreact/ent"
	"goreact/handlers"
	"goreact/indexFiles"
	"goreact/repositories"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/logger"

	_ "github.com/mattn/go-sqlite3"
)

func scanAndAddTracksToDb(client *ent.Client, ctx context.Context) {
	tracks, err := indexFiles.ScanPathForMusicFiles("./content")
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%d tracks found", len(tracks))

	err = repositories.RemoveDeletedEntities(tracks, client, ctx)
	if err != nil {
		log.Fatalln(err)
	}

	repositories.AddTracks(tracks, client, ctx)
}

func main() {
	client, err := ent.Open("sqlite3", "sqlite.db?cache=shared&_fk=1")

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
	app.Use(cache.New(cache.Config{
		Expiration:   10 * time.Minute,
		CacheControl: true,
	}))

	// Routes

	v1 := app.Group("/api/v1")

	track := v1.Group("/track")
	track.Get("/:id/stream", handlers.TrackStream(client, ctx))
	track.Get("/:id", handlers.TrackInfo(client, ctx))
	track.Put("/:id/like", handlers.LikeTrack(client, ctx))
	track.Delete("/:id/like", handlers.UnLikeTrack(client, ctx))

	v1.Post("/tracks/ids", handlers.TracksInfo(client, ctx))

	album := v1.Group("/album")
	album.Get("/:id", handlers.GetAlbum(client, ctx))
	album.Get("/:id/image", handlers.GetAlbumCover(client, ctx))

	artist := v1.Group("/artist")
	artist.Get("/:id", handlers.GetArtist(client, ctx))

	search := v1.Group("/search")
	search.Get("/:query", handlers.Search(client, ctx))

	playlist := v1.Group("/playlist")
	playlist.Get("/liked", handlers.GetLikedTracks(client, ctx))

	// app.Static("/", "web/build")

	// app.Get("/*", func(c *fiber.Ctx) error {
	// 	return c.SendFile("./web/build/index.html")
	// })

	// Start app

	log.Fatalln(app.Listen(":42000"))
}
