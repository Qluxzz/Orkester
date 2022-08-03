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

	track.Path = path
	track.Length = int(f.Info.NSamples) / int(f.Info.SampleRate)
	track.MimeType = "audio/flac"

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
					track.Title = value
				case "artist":
					track.Artists = append(track.Artists, value)
				case "album":
					track.AlbumName = value
				case "albumartist":
					track.AlbumArtist = value
				case "tracknumber":
					// Some tracknumbers are formatted as (tracknumber)/(amount of tracks)
					slices := strings.Split(value, "/")
					if trackNumber, err := strconv.Atoi(slices[0]); err == nil {
						track.TrackNumber = trackNumber
					}
				case "date":
					// The date must be in ISO 8601 format
					// but may be followed by a space character
					// then any text you wish

					var dateString string

					// If value contains optional text
					if strings.Contains(value, " ") {
						parts := strings.SplitN(value, " ", 2)
						dateString = parts[0]
					} else {
						dateString = value
					}

					date, err := ParseDate(dateString)

					if err != nil {
						log.Printf("%s invalid date format", value)
					} else {
						track.Date = date
					}
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
					MimeType: data.MIME,
				}
			}
		}
	}

	// Ignore tracks without any artist
	if len(track.Artists) == 0 {
		return nil, errors.New("track must have artist")
	}

	return track, nil
}
