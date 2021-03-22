package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func createSchemas(db *sqlx.DB) {
	trackSchema := `CREATE TABLE IF NOT EXISTS track(
		id INTEGER PRIMARY KEY,
		path TEXT NOT NULL
	);`

	_, err := db.Exec(trackSchema)
	if err != nil {
		log.Fatalln(err)
	}
}

type AddTrackRequest struct {
	Path string `db:"path"`
}

func indexFolder(path string) []AddTrackRequest {
	path, err := filepath.EvalSymlinks(path)
	if err != nil {
		log.Fatalln(err)
	}

	tracks := []AddTrackRequest{}

	filepath.Walk(path, func(path string, fileInfo os.FileInfo, err error) error {
		filename := strings.ToLower(fileInfo.Name())
		ext := filepath.Ext(filename)

		if isMusicFile(ext) {
			tracks = append(
				tracks,
				AddTrackRequest{
					Path: path,
				},
			)

			return nil
		}

		if isCoverImage(filename) {
			return nil
		}

		return nil
	})

	return tracks
}

func isMusicFile(extension string) bool {
	validFileExtensions := []string{".ogg", ".flac", ".mp3"}

	for _, validExtension := range validFileExtensions {
		if extension == validExtension {
			return true
		}
	}

	return false
}

func isCoverImage(filename string) bool {
	hasValidFileName := func() bool {
		validFilenames := []string{"cover", "folder"}

		for _, validFilename := range validFilenames {
			if strings.HasPrefix(filename, validFilename) {
				return true
			}
		}

		return false
	}()

	hasValidExtension := func() bool {
		validFileExtensions := []string{".jpg", ".jpeg", ".png"}

		for _, validFileExtension := range validFileExtensions {
			if strings.HasSuffix(filename, validFileExtension) {
				return true
			}
		}

		return false
	}()

	return hasValidFileName && hasValidExtension
}

type DBTrack struct {
	Id   int    `db:"id"`
	Path string `db:"path"`
}

func addTracks(tracks []AddTrackRequest, db *sqlx.DB) {
	log.Printf("Tracks found: %d", len(tracks))

	for _, track := range tracks {
		res := db.MustExec("INSERT INTO track (path) VALUES (?)", track.Path)
		last, err := res.LastInsertId()
		if err != nil {
			log.Fatalln(err)
		}

		log.Print(last)
	}
}

func main() {
	db, err := sqlx.Connect("sqlite3", ":memory:")
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	createSchemas(db)

	tracks := indexFolder("./content")
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

		track := DBTrack{}
		err = db.Get(&track, "SELECT * FROM track WHERE id=:id", id)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		c.Response().Header.Add("content-type", "audio/flac")

		stream, err := os.Open(track.Path)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.SendStream(stream)
	})

	app.Listen(":3000")
}
