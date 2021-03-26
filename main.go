package main

import (
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sqlx.Connect("sqlite3", ":memory:")
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	err = createSchemas(db)
	if err != nil {
		log.Fatalln(err)
	}

	tracks, err := ScanPathForMusicFiles("./content")
	if err != nil {
		log.Fatalln(err)
	}

	err = addTracks(tracks, db)
	if err != nil {
		log.Fatalln(err)
	}

	app := fiber.New()

	app.Static("/", "web/build")

	app.Get("/api/track/:id/image", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		type TrackImage struct {
			Image    []byte `db:"image"`
			MimeType string `db:"imagemimetype"`
		}

		image := TrackImage{}

		err = db.Get(
			&image,
			`SELECT 
				image, 
				imagemimetype 
			FROM 
				albums 
			WHERE 
				id = (
					SELECT 
						albumid 
					FROM 
						tracks 
					WHERE 
						id = ?
				)`,
			id,
		)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		c.Response().Header.Add("Content-Type", image.MimeType)
		return c.Send(image.Image)
	})

	app.Get("/api/track/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		tracks, err := getTrackByIds([]int{id}, db)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		if len(tracks) == 0 {
			return c.SendStatus(404)
		}

		return c.JSON(tracks[0])
	})

	app.Get("/api/track/:id/stream", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		var path string
		err = db.Get(&path, "SELECT path FROM tracks WHERE id=?", id)
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

	app.Get("/api/browse/genres", func(c *fiber.Ctx) error {
		genres := []string{}

		err = db.Select(&genres, "SELECT name FROM genres")
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.JSON(genres)
	})

	app.Get("/api/browse/artists", func(c *fiber.Ctx) error {
		artists := []string{}

		err = db.Select(&artists, "SELECT name FROM artists")
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.JSON(artists)
	})

	app.Get("/api/browse/artist/:artist", func(c *fiber.Ctx) error {
		trackIds := []int{}

		err = db.Select(
			&trackIds,
			`SELECT
				tracks.id
			 FROM tracks
			 INNER JOIN artists
			 	ON artists.id = tracks.artistid
			 WHERE
			 	artists.urlname = ?
		`, c.Params("artist"))
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		tracks, err := getTrackByIds(trackIds, db)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.JSON(tracks)
	})

	app.Get("/api/browse/genre/:genre", func(c *fiber.Ctx) error {
		trackIds := []int{}

		err = db.Select(
			&trackIds,
			`SELECT
				tracks.id
			 FROM tracks
			 INNER JOIN genres
			 	ON genres.id = tracks.genreid
			 WHERE
			 	genres.urlname = ?
		`, c.Params("genre"))
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		tracks, err := getTrackByIds(trackIds, db)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.JSON(tracks)
	})

	app.Get("/*", func(c *fiber.Ctx) error {
		return c.SendFile("./web/build/index.html")
	})

	log.Fatalln(app.Listen(":3001"))
}

func createSchemas(db *sqlx.DB) error {
	artistSchema := `CREATE TABLE IF NOT EXISTS artists(
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL UNIQUE,
		urlname TEXT NOT NULL UNIQUE
	);`

	albumSchema := `CREATE TABLE IF NOT EXISTS albums(
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL UNIQUE,
		image BLOB,
		imagemimetype TEXT
	);`

	genreSchema := `CREATE TABLE IF NOT EXISTS genres(
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL UNIQUE,
		urlname TEXT NOT NULL UNIQUE
	);`

	trackSchema := `CREATE TABLE IF NOT EXISTS tracks(
		id INTEGER PRIMARY KEY,
		title TEXT NOT NULL,
		tracknumber TEXT NOT NULL,
		path TEXT NOT NULL,
		date TEXT NOT NULL,
		albumid INTEGER,
		artistid INTEGER,
		genreid INTEGER,
		FOREIGN KEY (albumid) REFERENCES albums(id),
		FOREIGN KEY (artistid) REFERENCES artists(id),
		FOREIGN KEY (genreid) REFERENCES genres(id)
	);`

	tx := db.MustBegin()

	tx.MustExec(artistSchema)
	tx.MustExec(albumSchema)
	tx.MustExec(genreSchema)
	tx.MustExec(trackSchema)

	err := tx.Commit()
	return err
}
