package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func createSchemas(db *sqlx.DB) {
	artistSchema := `CREATE TABLE IF NOT EXISTS artist (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL
	);`

	_, err := db.Exec(artistSchema)
	if err != nil {
		log.Fatalln(err)
	}

	genreSchema := `CREATE TABLE IF NOT EXISTS genre (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL
	);`

	_, err = db.Exec(genreSchema)
	if err != nil {
		log.Fatalln(err)
	}

	albumSchema := `CREATE TABLE IF NOT EXISTS album (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL,
		date TEXT NOT NULL
	);`

	_, err = db.Exec(albumSchema)
	if err != nil {
		log.Fatalln(err)
	}

	albumArtistsSchema := `CREATE TABLE IF NOT EXISTS albumartist (
		albumid INTEGER NOT NULL,
		artistid INTEGER NOT NULL,
		FOREIGN KEY (albumid) REFERENCES album(id),
		FOREIGN KEY (artistid) REFERENCES artist(id)
	);`

	_, err = db.Exec(albumArtistsSchema)
	if err != nil {
		log.Fatalln(err)
	}

	albumGenresSchema := `CREATE TABLE IF NOT EXISTS albumgenre (
		albumid INTEGER NOT NULL,
		genreid INTEGER NOT NULL,
		FOREIGN KEY (albumid) REFERENCES album(id),
		FOREIGN KEY (genreid) REFERENCES genre(id)
	);`

	_, err = db.Exec(albumGenresSchema)
	if err != nil {
		log.Fatalln(err)
	}

	trackSchema := `CREATE TABLE IF NOT EXISTS track (
		id INTEGER PRIMARY KEY,
		name text NOT NULL,
		path TEXT NOT NULL,
		albumid INTEGER NOT NULL,
		FOREIGN KEY (albumid) REFERENCES album(id)
	);`

	_, err = db.Exec(trackSchema)
	if err != nil {
		log.Fatalln(err)
	}

	trackArtistsSchema := `CREATE TABLE IF NOT EXISTS trackartist (
		trackid INTEGER NOT NULL,
		artistid INTEGER NOT NULL,
		FOREIGN KEY (trackid) REFERENCES track(id),
		FOREIGN KEY (artistid) REFERENCES artist(id)
	);`

	_, err = db.Exec(trackArtistsSchema)
	if err != nil {
		log.Fatalln(err)
	}

}

type Artist struct {
	Id   int
	Name string
}

type Album struct {
	Id   int
	Name string
	Date time.Time
}

func indexFolder(path string) {
	filepath.Walk(path, walkFn)
}

func walkFn(path string, fileInfo os.FileInfo, err error) error {
	if fileInfo.IsDir() {
		return nil
	}

	if !isMusicFile(fileInfo.Name()) {
		return nil
	}

	log.Printf("Found music file: %s", path)

	return nil
}

func isMusicFile(filename string) bool {
	validFileExtensions := []string{"ogg", "flac", "mp3"}

	for _, extension := range validFileExtensions {
		if strings.HasSuffix(filename, extension) {
			return true
		}
	}

	return false
}

func main() {

	indexFolder("./content")

	return

	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	createSchemas(db)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		db, err := sqlx.Open("sqlite3", ":memory:")
		if err != nil {
			log.Fatalln(err)
		}

		defer db.Close()

		db.Queryx("SELECT * FROM tracks")

		return c.SendString("Hello, World 👋!")
	})

	app.Get("/stream/:type/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			log.Fatalln("Invalid id: ", id)
		}

		_type := c.Params("type")

		switch _type {
		case "album":
			getAlbum(db, id)
		case "artist":
			getArtist(db, id)
		case "track":
			getTrack(db, id)
		}

		//c.Response().Header.Add("content-type", "audio/ogg")

		return c.SendString(fmt.Sprintf("Fetching %s with id %d", _type, id))
	})

	app.Listen(":3000")
}

func getAlbum(db *sqlx.DB, id int) {
	log.Println("Fetching album with id:", id)
}

func getArtist(db *sqlx.DB, id int) {
	log.Println("Fetching artist with id:", id)
}

func getTrack(db *sqlx.DB, id int) {
	log.Println("Fetching track with id:", id)
}
