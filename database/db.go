package database

import (
	"github.com/jmoiron/sqlx"
)

func createSchemas(db *sqlx.DB) error {
	artistSchema := `CREATE TABLE IF NOT EXISTS artists(
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL UNIQUE
	);`

	albumSchema := `CREATE TABLE IF NOT EXISTS albums(
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL UNIQUE,
		image BLOB,
		imagemimetype TEXT
	);`

	genreSchema := `CREATE TABLE IF NOT EXISTS genres(
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL UNIQUE
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

var db *sqlx.DB

func GetInstance() (*sqlx.DB, error) {
	if db != nil {
		return db, nil
	}

	db, err := sqlx.Connect("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	err = createSchemas(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}
