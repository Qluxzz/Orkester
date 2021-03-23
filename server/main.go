package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func createSchemas(db *sqlx.DB) {
	trackSchema := `CREATE TABLE IF NOT EXISTS track(
		id INTEGER PRIMARY KEY,
		title TEXT NOT NULL,
		tracknumber TEXT NOT NULL,
		path TEXT NOT NULL,
		date TEXT NOT NULL
	);`

	_, err := db.Exec(trackSchema)
	if err != nil {
		log.Fatalln(err)
	}
}

type DBTrack struct {
	Id          int    `db:"id"`
	Title       string `db:"title"`
	TrackNumber string `db:"tracknumber"`
	Path        string `db:"path"`
	Date        string `db:"date"`
}

func addTracks(tracks []AddTrackRequest, db *sqlx.DB) {
	log.Printf("Tracks found: %d", len(tracks))

	insertTrackStmt := `
		INSERT INTO track (
			title,
			tracknumber,
			path,
			date
		) VALUES(
			?,
			?,
			?,
			?
		)
	`

	tx := db.MustBegin()
	for _, track := range tracks {
		tx.MustExec(
			insertTrackStmt,
			track.Title,
			track.TrackNumber,
			track.Path,
			track.Date,
		)
	}
	tx.Commit()
}

func main() {
	db, err := sqlx.Connect("sqlite3", ":memory:")
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	createSchemas(db)

	tracks := IndexFolder("./content")
	addTracks(tracks, db)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		tracks := []DBTrack{}
		db.Select(&tracks, "SELECT * FROM track")

		_json, err := json.Marshal(tracks)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		c.Response().Header.Add("Content-Type", "application/json")

		return c.Send(_json)
	})

	app.Get("/stream/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		var path string
		err = db.Get(&path, "SELECT path FROM track WHERE id=:id", id)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		c.Response().Header.Add("content-type", "audio/flac")

		stream, err := os.Open(path)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.SendStream(stream)
	})

	log.Fatalln(app.Listen(":3001"))
}
