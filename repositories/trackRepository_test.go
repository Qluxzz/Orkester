package repositories

import (
	"goreact/database"
	"goreact/indexFiles"
	"testing"
)

func TestMultipleAlbumsWithSameName(t *testing.T) {
	db, err := database.GetInstance()
	if err != nil {
		t.Error(err.Error())
	}

	track := new(indexFiles.IndexedTrack)
	track.Album.Name = indexFiles.CreateValidNullString("Foo")
	track.Artist = indexFiles.CreateValidNullString("Test")

	track2 := new(indexFiles.IndexedTrack)
	track2.Album.Name = indexFiles.CreateValidNullString("Foo")
	track2.Artist = indexFiles.CreateValidNullString("Test2")

	tracks := []*indexFiles.IndexedTrack{
		track,
		track2,
	}

	err = AddTracks(tracks, db)
	if err != nil {
		t.Error(err.Error())
	}
}
