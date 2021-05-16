package models

import "database/sql"

type DBArtist struct {
	Id      sql.NullInt32  `db:"id"`
	Name    sql.NullString `db:"name"`
	UrlName sql.NullString `db:"urlname"`
}

type DBTrack struct {
	Id            int            `db:"id"`
	Title         string         `db:"title"`
	TrackNumber   int            `db:"tracknumber"`
	Date          string         `db:"date"`
	Length        int            `db:"length"`
	AlbumId       sql.NullInt32  `db:"albumid"`
	AlbumName     sql.NullString `db:"albumname"`
	AlbumUrlName  sql.NullString `db:"albumurlname"`
	GenreId       sql.NullInt32  `db:"genreid"`
	GenreName     sql.NullString `db:"genrename"`
	GenreUrlName  sql.NullString `db:"genreurlname"`
	Image         []byte         `db:"image"`
	ImageMimeType string         `db:"imagemimetype"`
}

type Track struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	TrackNumber int       `json:"trackNumber"`
	Date        string    `json:"date"`
	Length      int       `json:"length"`
	Album       *Album    `json:"album"`
	Artists     []*Artist `json:"artists"`
	Genre       *Genre    `json:"genre"`
}

type IdNameAndUrlName struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	UrlName string `json:"urlName"`
}

type Album = IdNameAndUrlName
type Artist = IdNameAndUrlName
type Genre = IdNameAndUrlName

func (track DBTrack) ToDomain(dbArtists []DBArtist) Track {
	artists := func() []*Artist {
		artists := []*Artist{}

		for _, dbArtist := range dbArtists {
			artists = append(artists, &Artist{
				Id:      int(dbArtist.Id.Int32),
				Name:    dbArtist.Name.String,
				UrlName: dbArtist.UrlName.String,
			})
		}

		return artists
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
		Artists:     artists,
	}
}
