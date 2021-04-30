package main

import (
	"log"

	"goreact/database"
	"goreact/handlers"
	"goreact/indexFiles"
	"goreact/repositories"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func scanAndAddTracksToDb(db *sqlx.DB) {
	tracks, err := indexFiles.ScanPathForMusicFiles("./content")
	if err != nil {
		log.Fatalln(err)
	}

	err = repositories.AddTracks(tracks, db)
	if err != nil {
		log.Fatalln(err)
	}

	log.Print("Tracks has been added")
}

func main() {
	db, err := database.GetInstance()
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	go scanAndAddTracksToDb(db)

	app := fiber.New()

	// Middlewares

	app.Use(logger.New())

	// Routes

	v1 := app.Group("/api/v1")

	track := v1.Group("/track")

	track.Get("/:id/image", handlers.TrackImage(db))
	track.Get("/:id/stream", handlers.TrackStream(db))
	track.Get("/:id", handlers.TrackInfo(db))

	album := v1.Group("/album")

	album.Get("/:id", handlers.GetAlbum(db))

	artist := v1.Group("/artist")
	artist.Get("/:id", handlers.GetArtist(db))

	// app.Static("/", "web/build")

	// app.Get("/*", func(c *fiber.Ctx) error {
	// 	return c.SendFile("./web/build/index.html")
	// })

	// Start app
	log.Fatalln(app.Listen(":42000"))
}
