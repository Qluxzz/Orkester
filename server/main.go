package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func createSchemas(db *sqlx.DB) {
	artistSchema := `CREATE TABLE IF NOT EXISTS artists(
		id INTEGER PRIMARY KEY,
		name TEXT UNQIUE NOT NULL 
	);`

	albumSchema := `CREATE TABLE IF NOT EXISTS albums(
		id INTEGER PRIMARY KEY,
		name TEXT UNIQUE NOT NULL
	);`

	trackSchema := `CREATE TABLE IF NOT EXISTS tracks(
		id INTEGER PRIMARY KEY,
		title TEXT NOT NULL,
		tracknumber TEXT NOT NULL,
		path TEXT NOT NULL,
		date TEXT NOT NULL,
		albumid INTEGER,
		artistid INTEGER,
		FOREIGN KEY (albumid) REFERENCES albums(id),
		FOREIGN KEY (artistid) REFERENCES artists(id)
	);`

	tx := db.MustBegin()

	tx.MustExec(artistSchema)
	tx.MustExec(albumSchema)
	tx.MustExec(trackSchema)

	err := tx.Commit()
	if err != nil {
		log.Fatalln(err)
	}
}

type Track struct {
	Id          int
	Title       string
	TrackNumber string
	Path        string
	Date        string
	Album       string
	Artist      string
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

	return Track{
		Id:          track.Id,
		Title:       track.Title,
		TrackNumber: track.TrackNumber,
		Path:        track.Path,
		Date:        track.Date,
		Album:       album,
		Artist:      artist,
	}
}

type DBTrack struct {
	Id          int            `db:"id"`
	Title       string         `db:"title"`
	TrackNumber string         `db:"tracknumber"`
	Path        string         `db:"path"`
	Date        string         `db:"date"`
	Album       sql.NullString `db:"album"`
	Artist      sql.NullString `db:"artist"`
}

func addTracks(tracks []AddTrackRequest, db *sqlx.DB) {
	log.Printf("Tracks found: %d", len(tracks))

	tx := db.MustBegin()

	insertArtistStmt := `
		INSERT OR IGNORE INTO artists (name) VALUES (?)
	`

	insertAlbumStmt := `
		INSERT OR IGNORE INTO albums (name) VALUES (?)
	`

	insertTrackStmt := `
		INSERT INTO tracks (
			title,
			tracknumber,
			path,
			date,
			albumid,
			artistid
		) VALUES(
			?,
			?,
			?,
			?,
			(SELECT id FROM albums WHERE name = ?),
			(SELECT id FROM artists WHERE name = ?)
		)
	`

	for _, track := range tracks {
		tx.MustExec(insertArtistStmt, track.Artist)
		tx.MustExec(insertAlbumStmt, track.Album)

		tx.MustExec(
			insertTrackStmt,
			track.Title,
			track.TrackNumber,
			track.Path,
			track.Date,
			track.Album,
			track.Artist,
		)
	}

	err := tx.Commit()
	if err != nil {
		log.Fatalln(err)
	}
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
		dbTracks := []DBTrack{}

		getTracksStmt := `
			SELECT
				t.id,
				t.title,
				t.tracknumber,
				t.path,
				t.date,
				albums.name album,
				artists.name artist
			FROM 
				tracks t
			LEFT JOIN artists
				ON artists.id = t.artistid
			LEFT JOIN albums
				ON albums.id = t.albumid
		`

		err := db.Select(&dbTracks, getTracksStmt)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		tracks := []Track{}

		for _, dbTrack := range dbTracks {
			tracks = append(tracks, dbTrack.ToDomain())
		}

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
