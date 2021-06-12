package models

type Track struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	TrackNumber int       `json:"trackNumber"`
	Date        string    `json:"date"`
	Length      int       `json:"length"`
	Album       *Album    `json:"album"`
	Artists     []*Artist `json:"artists"`
	Genre       *Genre    `json:"genre"`
	Liked       bool      `json:"liked"`
}

type IdNameAndUrlName struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	UrlName string `json:"urlName"`
}

type Album = IdNameAndUrlName
type Artist = IdNameAndUrlName
type Genre = IdNameAndUrlName
