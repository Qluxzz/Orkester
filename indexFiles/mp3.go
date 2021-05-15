package indexFiles

import (
	"errors"
	"strconv"
	"strings"

	"github.com/mikkyang/id3-go"
	v2 "github.com/mikkyang/id3-go/v2"
)

// Info on frames and fields can be found here
// https://id3.org/id3v2.3.0 (2021-05-04)

func ParseMp3File(path string) (*IndexedTrack, error) {
	mp3File, err := id3.Open(path)
	if err != nil {
		return nil, err
	}

	defer mp3File.Close()

	track := new(IndexedTrack)
	track.Title = CreateValidNullString(trimNullFromString(mp3File.Title()))
	track.Path = CreateValidNullString(path)
	track.MimeType = CreateValidNullString("audio/mpeg")

	trackNumberFrame, valid := mp3File.Frame("TRCK").(*v2.TextFrame)
	if valid {
		trackNumber, err := strconv.Atoi(trimNullFromString(trackNumberFrame.Text()))
		if err == nil {
			track.TrackNumber = CreateValidNullInt(trackNumber)
		}
	}

	lengthFrame, valid := mp3File.Frame("TLEN").(*v2.TextFrame)
	if valid {
		lengthMs, err := strconv.Atoi(trimNullFromString(lengthFrame.Text()))
		if err == nil {
			track.Length = CreateValidNullInt(lengthMs / 1000)
		}
	}

	track.AlbumName = CreateValidNullString(trimNullFromString(mp3File.Album()))

	imageFrame, valid := mp3File.Frame("APIC").(*v2.ImageFrame)
	if valid {
		track.Image = &Image{
			Data:     imageFrame.Data(),
			MimeType: CreateValidNullString(imageFrame.MIMEType()),
		}
	}
	track.ArtistName = CreateValidNullString(trimNullFromString(mp3File.Artist()))
	track.Genre = CreateValidNullString(trimNullFromString(mp3File.Genre()))
	track.Date = CreateValidNullString(trimNullFromString(mp3File.Year()))

	if !track.ArtistName.Valid {
		return nil, errors.New("track was missing artist")
	}

	return track, nil
}

func trimNullFromString(s string) string {
	return strings.Trim(s, "\x00")
}
