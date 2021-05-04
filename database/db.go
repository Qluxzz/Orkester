package database

import (
	"github.com/jmoiron/sqlx"
)

func createSchemas(db *sqlx.DB) error {
	artistSchema := `CREATE TABLE IF NOT EXISTS artists(
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL UNIQUE,
		urlname TEXT NOT NULL
	);`

	albumSchema := `CREATE TABLE IF NOT EXISTS albums(
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL UNIQUE,
		urlname TEXT NOT NULL,
		image BLOB,
		imagemimetype TEXT
	);`

	genreSchema := `CREATE TABLE IF NOT EXISTS genres(
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL UNIQUE,
		urlname TEXT NOT NULL
	);`

	trackSchema := `CREATE TABLE IF NOT EXISTS tracks(
		id INTEGER PRIMARY KEY,
		title TEXT NOT NULL,
		tracknumber TEXT NOT NULL,
		path TEXT NOT NULL,
		date TEXT NOT NULL,
		length INT NOT NULL,
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

var db *sqlx.DB

func GetInstance() (*sqlx.DB, error) {
	if db != nil {
		return db, nil
	}

	/*
		Each connection to ":memory:" opens a brand new in-memory sql database,
		so if the stdlib's sql engine happens to open another connection and you've only specified ":memory:",
		that connection will see a brand new database.
		A workaround is to use "file::memory:?cache=shared" (or "file:foobar?mode=memory&cache=shared").
		Every connection to this string will point to the same in-memory database.

		https://github.com/mattn/go-sqlite3#faq
	*/
	db, err := sqlx.Connect("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		return nil, err
	}

	err = createSchemas(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}
