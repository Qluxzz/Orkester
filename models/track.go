package models

import "database/sql"

type DBTrack struct {
	Id            int            `db:"id"`
	Title         string         `db:"title"`
	TrackNumber   string         `db:"tracknumber"`
	Date          string         `db:"date"`
	Length        int            `db:"length"`
	AlbumId       sql.NullInt32  `db:"albumid"`
	AlbumName     sql.NullString `db:"albumname"`
	AlbumUrlName  sql.NullString `db:"albumurlname"`
	ArtistId      sql.NullInt32  `db:"artistid"`
	ArtistName    sql.NullString `db:"artistname"`
	ArtistUrlName sql.NullString `db:"artisturlname"`
	GenreId       sql.NullInt32  `db:"genreid"`
	GenreName     sql.NullString `db:"genrename"`
	GenreUrlName  sql.NullString `db:"genreurlname"`
	Image         []byte         `db:"image"`
	ImageMimeType string         `db:"imagemimetype"`
}

type Track struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	TrackNumber string  `json:"trackNumber"`
	Date        string  `json:"date"`
	Length      int     `json:"length"`
	Album       *Album  `json:"album"`
	Artist      *Artist `json:"artist"`
	Genre       *Genre  `json:"genre"`
}

type Album struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	UrlName string `json:"urlName"`
}

type Artist struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	UrlName string `json:"urlName"`
}

type Genre struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	UrlName string `json:"urlName"`
}

func (track DBTrack) ToDomain() Track {
	artist := func() *Artist {
		if !track.ArtistId.Valid {
			return nil
		}

		a := new(Artist)
		a.Id = int(track.ArtistId.Int32)
		a.Name = track.ArtistName.String
		a.UrlName = track.ArtistUrlName.String

		return a
	}()

	album := func() *Album {
		if !track.AlbumId.Valid {
			return nil
		}

		a := new(Album)
		a.Id = int(track.AlbumId.Int32)
		a.Name = track.AlbumName.String
		a.UrlName = track.AlbumUrlName.String

		return a
	}()

	genre := func() *Genre {
		if !track.GenreId.Valid {
			return nil
		}

		g := new(Genre)
		g.Id = int(track.GenreId.Int32)
		g.Name = track.GenreName.String
		g.UrlName = track.GenreUrlName.String

		return g
	}()

	return Track{
		Id:          track.Id,
		Title:       track.Title,
		TrackNumber: track.TrackNumber,
		Date:        track.Date,
		Length:      track.Length,
		Genre:       genre,
		Album:       album,
		Artist:      artist,
	}
}
