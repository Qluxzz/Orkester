package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gosimple/slug"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func createSchemas(db *sqlx.DB) {
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
	if err != nil {
		log.Fatalln(err)
	}
}

type Track struct {
	Id            int
	Title         string
	TrackNumber   string
	Date          string
	Album         string
	Artist        string
	Genre         string
	Image         []byte
	ImageMimeType string
}

func (track DBTrack) ToDomain() Track {
	artist := func() string {
		if !track.Artist.Valid {
			return ""
		}

		return track.Artist.String
	}()

	album := func() string {
		if !track.Album.Valid {
			return ""
		}

		return track.Album.String
	}()

	genre := func() string {
		if !track.Genre.Valid {
			return ""
		}

		return track.Genre.String
	}()

	return Track{
		Id:          track.Id,
		Title:       track.Title,
		TrackNumber: track.TrackNumber,
		Date:        track.Date,
		Genre:       genre,
		Album:       album,
		Artist:      artist,
	}
}

type DBTrack struct {
	Id            int            `db:"id"`
	Title         string         `db:"title"`
	TrackNumber   string         `db:"tracknumber"`
	Date          string         `db:"date"`
	Album         sql.NullString `db:"album"`
	Artist        sql.NullString `db:"artist"`
	Genre         sql.NullString `db:"genre"`
	Image         []byte         `db:"image"`
	ImageMimeType string         `db:"imagemimetype"`
}

func addTracks(tracks []AddTrackRequest, db *sqlx.DB) {
	log.Printf("Tracks found: %d", len(tracks))

	tx := db.MustBegin()

	insertArtistStmt := `
		INSERT OR IGNORE INTO artists (name, urlname) VALUES (?, ?)
	`

	insertAlbumStmt := `
		INSERT OR IGNORE INTO albums (name, image, imagemimetype) VALUES (?, ?, ?)
	`

	insertGenreStmt := `
		INSERT OR IGNORE INTO genres (name, urlname) VALUES (?, ?)
	`

	insertTrackStmt := `
		INSERT INTO tracks (
			title,
			tracknumber,
			path,
			date,
			albumid,
			artistid,
			genreid
		) VALUES(
			?,
			?,
			?,
			?,
			(SELECT id FROM albums WHERE name = ?),
			(SELECT id FROM artists WHERE name = ?),
			(SELECT id from genres WHERE name = ?)
		)
	`

	for _, track := range tracks {
		if track.Artist != "" {
			tx.MustExec(insertArtistStmt, track.Artist, slug.Make(track.Artist))
		}
		tx.MustExec(insertAlbumStmt, track.Album.Name, track.Album.Image.Data, track.Album.Image.MimeType)
		if track.Genre != "" {
			tx.MustExec(insertGenreStmt, track.Genre, slug.Make(track.Genre))
		}
		tx.MustExec(
			insertTrackStmt,
			track.Title,
			track.TrackNumber,
			track.Path,
			track.Date,
			track.Album.Name,
			track.Artist,
			track.Genre,
		)
	}

	err := tx.Commit()
	if err != nil {
		log.Fatalln(err)
	}
}

func getTrackByIds(ids []int, db *sqlx.DB) ([]Track, error) {
	query, args, err := sqlx.In(
		`SELECT
				t.id,
				t.title,
				t.tracknumber,
				t.date,
				albums.name album,
				artists.name artist,
				genres.name genre
			FROM 
				tracks t
			LEFT JOIN artists
				ON artists.id = t.artistid
			LEFT JOIN albums
				ON albums.id = t.albumid
			LEFT JOIN genres
				ON genres.id = t.genreid
			WHERE t.id IN (?)
		`,
		ids,
	)

	if err != nil {
		return nil, err
	}

	dbTracks := []DBTrack{}
	err = db.Select(&dbTracks, query, args...)
	if err != nil {
		return nil, err
	}

	tracks := []Track{}

	for _, dbTrack := range dbTracks {
		tracks = append(tracks, dbTrack.ToDomain())
	}

	return tracks, nil
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

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3002",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		getTracksStmt := `
			SELECT
				id
			FROM 
				tracks
		`
		trackIds := []int{}
		err := db.Select(&trackIds, getTracksStmt)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		tracks, err := getTrackByIds(trackIds, db)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		_json, err := json.Marshal(tracks)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		c.Response().Header.Add("Content-Type", "application/json")

		return c.Send(_json)
	})

	app.Get("/track/:id/image", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		type TrackImage struct {
			Image    []byte `db:"image"`
			MimeType string `db:"imagemimetype"`
		}

		image := TrackImage{}

		err = db.Get(&image, "SELECT image, imagemimetype FROM albums WHERE id = (SELECT albumid FROM tracks WHERE id = ?)", id)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		c.Response().Header.Add("Content-Type", image.MimeType)
		return c.Send(image.Image)
	})

	app.Get("/track/:id", func(c *fiber.Ctx) error {
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

		_json, err := json.Marshal(tracks[0])
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

	app.Get("/browse/genres", func(c *fiber.Ctx) error {
		genres := []string{}

		err = db.Select(&genres, "SELECT name FROM genres")
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		_json, err := json.Marshal(genres)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		c.Response().Header.Add("Content-Type", "application/json")
		return c.Send(_json)
	})

	app.Get("/browse/artists", func(c *fiber.Ctx) error {
		artists := []string{}

		err = db.Select(&artists, "SELECT name FROM artists")
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		_json, err := json.Marshal(artists)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		c.Response().Header.Add("Content-Type", "application/json")
		return c.Send(_json)
	})

	app.Get("/browse/artist/:artist", func(c *fiber.Ctx) error {
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

		c.Response().Header.Add("Content-Type", "application/json")

		_json, err := json.Marshal(&tracks)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.Send(_json)
	})

	app.Get("/browse/genre/:genre", func(c *fiber.Ctx) error {
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

		c.Response().Header.Add("Content-Type", "application/json")

		_json, err := json.Marshal(&tracks)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.Send(_json)
	})

	log.Fatalln(app.Listen(":3001"))
}
