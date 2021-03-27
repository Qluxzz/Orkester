package models

import "database/sql"

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
