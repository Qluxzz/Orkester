package models

import (
	"orkester/ent"
	"time"
)

type Track struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	TrackNumber int       `json:"trackNumber"`
	Length      int       `json:"length"`
	Album       *Album    `json:"album,omitempty"`
	Artists     []*Artist `json:"artists"`
	Liked       bool      `json:"liked"`
}

type TrackWithPath struct {
	Track
	Path string `json:"path"`
}

type TrackWithDate struct {
	Track
	DateAdded time.Time `json:"dateAdded"`
}

type IdNameAndUrlName struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	UrlName string `json:"urlName"`
}

type Album = IdNameAndUrlName
type Artist = IdNameAndUrlName

func FromEntTracks(dbTracks []*ent.Track) []Track {
	tracks := []Track{}

	for _, track := range dbTracks {
		tracks = append(tracks, FromEntTrack(track))
	}

	return tracks
}

func FromEntTrack(dbTrack *ent.Track) Track {
	track := Track{
		Id:          dbTrack.ID,
		TrackNumber: dbTrack.TrackNumber,
		Title:       dbTrack.Title,
		Length:      dbTrack.Length,
		Liked:       dbTrack.Edges.Liked != nil,
		Album: &Album{
			Id:      dbTrack.Edges.Album.ID,
			Name:    dbTrack.Edges.Album.Name,
			UrlName: dbTrack.Edges.Album.URLName,
		},
	}

	artists := []*Artist{}

	for _, artist := range dbTrack.Edges.Artists {
		a := &Artist{
			Id:      artist.ID,
			Name:    artist.Name,
			UrlName: artist.URLName,
		}

		artists = append(artists, a)
	}

	track.Artists = artists

	return track
}

func FromEntTrackWithPath(dbTrack *ent.Track) TrackWithPath {
	return TrackWithPath{
		Track: FromEntTrack(dbTrack),
		Path:  dbTrack.Path,
	}
}
