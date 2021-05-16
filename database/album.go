package database

import (
	"goreact/models"
	"sort"

	"github.com/jmoiron/sqlx"
)

type album struct {
	Id      int            `json:"id"`
	Name    string         `json:"name"`
	UrlName string         `json:"urlName"`
	Date    string         `json:"date"`
	Tracks  []models.Track `json:"tracks"`
	Artist  albumArtist    `json:"artist"`
}

type albumArtist struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	UrlName string `json:"urlName"`
}

type dbAlbum struct {
	Id       int    `db:"id"`
	Name     string `db:"name"`
	UrlName  string `db:"urlname"`
	Date     string `db:"date"`
	ArtistId int    `db:"artistid"`
}

func GetAlbum(albumId int, db *sqlx.DB) (*album, error) {
	dbAlbum := dbAlbum{}

	err := db.Get(
		&dbAlbum,
		`
			SELECT
				id,
				name,
				urlname,
				(SELECT date FROM tracks WHERE albumid = $1 LIMIT 1) date,
				artistid
			FROM 
				albums
			WHERE
				id = $1
			`,
		albumId,
	)

	if err != nil {
		return nil, err
	}

	tracks, err := getAlbumTracks(dbAlbum.Id, db)
	if err != nil {
		return nil, err
	}

	artist, err := GetArtistById(dbAlbum.ArtistId, db)
	if err != nil {
		return nil, err
	}

	return toDomain(&dbAlbum, tracks, artist), nil
}

func toDomain(dbAlbum *dbAlbum, tracks []models.Track, artist *artist) *album {
	return &album{
		Id:      dbAlbum.Id,
		Name:    dbAlbum.Name,
		UrlName: dbAlbum.UrlName,
		Date:    dbAlbum.Date,
		Tracks:  tracks,
		Artist: albumArtist{
			Id:      artist.Id,
			Name:    artist.Name,
			UrlName: artist.UrlName,
		},
	}
}

type albumImage struct {
	Image    []byte `db:"image"`
	MimeType string `db:"imagemimetype"`
}

func GetAlbumCover(albumId int, db *sqlx.DB) (*albumImage, error) {
	image := albumImage{}

	err := db.Get(
		&image,
		`SELECT
			image,
			imagemimetype
		FROM
			albums
		WHERE
			id = ?`,
		albumId,
	)
	if err != nil {
		return nil, err
	}

	return &image, nil
}

func getAlbumTracks(albumId int, db *sqlx.DB) ([]models.Track, error) {
	trackIds := []int{}

	err := db.Select(
		&trackIds,
		`SELECT
				id
			FROM
				tracks
			WHERE
				albumid = ?
			`,
		albumId,
	)

	if err != nil {
		return nil, err
	}

	tracks, err := GetTracksByIds(trackIds, db)
	if err != nil {
		return nil, err
	}

	// Sort by track number ascending
	sort.SliceStable(tracks, func(i int, j int) bool { return tracks[i].TrackNumber < tracks[j].TrackNumber })

	return tracks, nil
}
