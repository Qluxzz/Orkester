package database

import (
	"goreact/models"

	"github.com/jmoiron/sqlx"
)

func GetLikedTracks(db *sqlx.DB) ([]models.Track, error) {
	trackIds := []int{}

	err := db.Select(
		&trackIds,
		"SELECT trackid FROM likedTracks",
	)

	if err != nil {
		return nil, err
	}

	tracks, err := GetTracksByIds(trackIds, db)
	if err != nil {
		return nil, err
	}

	return tracks, nil
}
