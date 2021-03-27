package main

import (
	"log"

	"goreact/database"
	"goreact/handlers"
	"goreact/indexFiles"
	"goreact/repositories"

	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	tracks, err := indexFiles.ScanPathForMusicFiles("./content")
	if err != nil {
		log.Fatalln(err)
	}

	db, err := database.GetInstance()
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	err = repositories.AddTracks(tracks, db)
	if err != nil {
		log.Fatalln(err)
	}

	app := fiber.New()

	v1 := app.Group("/api/v1")

	track := v1.Group("/track")

	track.Get("/:id/image", handlers.TrackImage(db))
	track.Get("/:id/stream", handlers.TrackStream(db))
	track.Get("/:id", handlers.TrackInfo(db))

	browse := v1.Group("/browse")

	browse.Get("/artists", handlers.BrowseArtists(db))
	browse.Get("/artists/:artist-url-name", handlers.BrowseArtist(db))

	browse.Get("/genres", handlers.BrowseGenres(db))
	browse.Get("/genres/:genre-url-name", handlers.BrowseGenre((db)))

	app.Static("/", "web/build")

	app.Get("/*", func(c *fiber.Ctx) error {
		return c.SendFile("./web/build/index.html")
	})

	log.Fatalln(app.Listen(":3000"))
}
