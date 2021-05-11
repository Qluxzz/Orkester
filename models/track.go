package models

import "database/sql"

type DBTrack struct {
	Id            int            `db:"id"`
	Title         string         `db:"title"`
	TrackNumber   int            `db:"tracknumber"`
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
	TrackNumber int     `json:"trackNumber"`
	Date        string  `json:"date"`
	Length      int     `json:"length"`
	Album       *Album  `json:"album"`
	Artist      *Artist `json:"artist"`
	Genre       *Genre  `json:"genre"`
}

type IdNameAndUrlName struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	UrlName string `json:"urlName"`
}

type Album = IdNameAndUrlName
type Artist = IdNameAndUrlName
type Genre = IdNameAndUrlName

func (track DBTrack) ToDomain() Track {
	artist := func() *Artist {
		if !track.ArtistId.Valid {
			return nil
		}

		return &Artist{
			Id:      int(track.ArtistId.Int32),
			Name:    track.ArtistName.String,
			UrlName: track.ArtistUrlName.String,
		}
	}()

	album := func() *Album {
		if !track.AlbumId.Valid {
			return nil
		}

		return &Album{
			Id:      int(track.AlbumId.Int32),
			Name:    track.AlbumName.String,
			UrlName: track.AlbumUrlName.String,
		}
	}()

	genre := func() *Genre {
		if !track.GenreId.Valid {
			return nil
		}

		return &Genre{
			Id:      int(track.GenreId.Int32),
			Name:    track.GenreName.String,
			UrlName: track.GenreUrlName.String,
		}
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
