package repositories

import (
	"context"
	"goreact/ent"
	"goreact/ent/album"
	"goreact/ent/artist"
	"goreact/ent/track"
	"goreact/indexFiles"
	"log"
)

func RemoveDeletedEntities(tracks []*indexFiles.IndexedTrack, client *ent.Client, context context.Context) error {
	db_tracks, err := client.
		Track.
		Query().
		WithAlbum(func(aq *ent.AlbumQuery) {
			aq.Select(album.FieldName)
		}).
		Select(track.FieldTrackNumber, track.FieldTitle).
		All(context)

	if err != nil {
		return err
	}

	tracks_removed_from_disk := GetTracksRemovedFromDisk(tracks, db_tracks)

	ids_to_be_removed := []int{}

	for _, removed_track := range tracks_removed_from_disk {
		ids_to_be_removed = append(ids_to_be_removed, removed_track.ID)
	}

	removed_tracks, err := client.Track.Delete().Where(track.IDIn(ids_to_be_removed...)).Exec(context)

	if err != nil {
		return err
	}

	log.Printf("Removed %d tracks", removed_tracks)

	// Remove albums without tracks
	removed_albums, err := client.Album.Delete().Where(album.Not(album.HasTracks())).Exec(context)
	if err != nil {
		return err
	}

	log.Printf("Removed %d albums", removed_albums)

	// Remove artists without albums or tracks
	removed_artists, err := client.
		Artist.
		Delete().
		Where(
			artist.And(
				artist.Not(artist.HasTracks()),
				artist.Not(artist.HasAlbums()),
			),
		).
		Exec(context)
	if err != nil {
		return err
	}

	log.Printf("Removed %d artists", removed_artists)

	return nil
}

func GetTracksRemovedFromDisk(tracks []*indexFiles.IndexedTrack, dbTracks []*ent.Track) []*ent.Track {
	removed_tracks := []*ent.Track{}

	for _, dbTrack := range dbTracks {
		if !TrackExistsOnDisk(tracks, dbTrack) {
			removed_tracks = append(removed_tracks, dbTrack)
		}
	}

	return removed_tracks
}

func TrackExistsOnDisk(tracks []*indexFiles.IndexedTrack, dbTrack *ent.Track) bool {
	for _, track := range tracks {
		if IsSameTrack(track, dbTrack) {
			return true
		}
	}

	return false
}

// This is a copy of the indexes found in ent/schema/track.go
// Which defines the unique constraint for a track
func IsSameTrack(track *indexFiles.IndexedTrack, dbTrack *ent.Track) bool {
	return track.Title.String == dbTrack.Title &&
		int(track.TrackNumber.Int32) == dbTrack.TrackNumber &&
		track.AlbumName.String == dbTrack.Edges.Album.Name
}
