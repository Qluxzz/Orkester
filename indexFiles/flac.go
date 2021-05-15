package indexFiles

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/mewkiz/flac"
	"github.com/mewkiz/flac/meta"
)

func ParseFlacFile(path string) (*IndexedTrack, error) {
	f, err := flac.ParseFile(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	track := new(IndexedTrack)

	track.Path = CreateValidNullString(path)
	track.Length = CreateValidNullInt(int(f.Info.NSamples) / int(f.Info.SampleRate))
	track.MimeType = CreateValidNullString("audio/flac")

	for _, block := range f.Blocks {
		switch block.Type {
		case meta.TypeVorbisComment:
			data, valid := block.Body.(*meta.VorbisComment)
			if !valid {
				log.Fatalln("Block said it was TypeVorbisComment but could not be cast to it!")
			}

			for _, tag := range data.Tags {
				tagType := strings.ToLower(tag[0])
				value := strings.TrimSpace(tag[1])

				switch tagType {
				case "title":
					track.Title = CreateValidNullString(value)
				case "artist":
					track.ArtistName = CreateValidNullString(value)
				case "album":
					track.AlbumName = CreateValidNullString(value)
				case "albumartist":
					track.AlbumArtist = CreateValidNullString(value)
				case "tracknumber":
					if trackNumber, err := strconv.Atoi(value); err == nil {
						track.TrackNumber = CreateValidNullInt(trackNumber)
					}
				case "genre":
					track.Genre = CreateValidNullString(value)
				case "date":
					track.Date = CreateValidNullString(value)
				}
			}
		case meta.TypePicture:
			data, valid := block.Body.(*meta.Picture)
			if !valid {
				log.Fatalln("Block said it was TypePicture but could not be cast to it!")
			}

			coverFront := uint32(3)

			if data.Type == coverFront {
				track.Image = &Image{
					Data:     data.Data,
					MimeType: CreateValidNullString(data.MIME),
				}
			}
		}
	}

	if !track.ArtistName.Valid {
		return nil, errors.New("track was missing artist")
	}

	return track, nil
}
