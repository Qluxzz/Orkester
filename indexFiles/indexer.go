package indexFiles

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/mewkiz/flac"
	"github.com/mewkiz/flac/meta"
	"github.com/mikkyang/id3-go"
	v2 "github.com/mikkyang/id3-go/v2"
)

func ScanPathForMusicFiles(path string) ([]*IndexedTrack, error) {
	path, err := filepath.EvalSymlinks(path)
	if err != nil {
		return nil, err
	}

	tracks := []*IndexedTrack{}

	filepath.Walk(path, func(path string, fileInfo os.FileInfo, err error) error {
		filename := strings.ToLower(fileInfo.Name())
		ext := filepath.Ext(filename)

		switch ext {
		case ".flac":
			track, err := parseFlacFile(path)
			if err == nil {
				tracks = append(tracks, track)
			}
		case ".mp3":
			track, err := parseMp3File(path)
			if err == nil {
				tracks = append(tracks, track)
			}
		}

		return nil
	})

	return tracks, nil
}

func parseFlacFile(path string) (*IndexedTrack, error) {
	f, err := flac.ParseFile(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	track := new(IndexedTrack)

	track.Path = CreateValidNullString(path)
	track.Length = int(f.Info.NSamples) / int(f.Info.SampleRate)

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
					track.Artist = CreateValidNullString(value)
				case "album":
					track.Album.Name = CreateValidNullString(value)
				case "albumartist":
					track.AlbumArtist = CreateValidNullString(value)
				case "tracknumber":
					trackNumber, err := strconv.Atoi(value)
					if err == nil {
						track.TrackNumber = trackNumber
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
				track.Album.Image = Image{
					Data:     data.Data,
					MimeType: CreateValidNullString(data.MIME),
				}
			}
		}
	}

	if !track.Artist.Valid {
		return nil, errors.New("track was missing artist")
	}

	return track, nil
}

func CreateValidNullString(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

// Info on frames and fields can be found here
// https://id3.org/id3v2.3.0 (2021-05-04)

func parseMp3File(path string) (*IndexedTrack, error) {
	mp3File, err := id3.Open(path)
	if err != nil {
		return nil, err
	}

	defer mp3File.Close()

	track := new(IndexedTrack)
	track.Title = CreateValidNullString(TrimNullFromString(mp3File.Title()))
	track.Path = CreateValidNullString(path)

	trackNumberFrame, valid := mp3File.Frame("TRCK").(*v2.TextFrame)
	if valid {
		trackNumber, err := strconv.Atoi(TrimNullFromString(trackNumberFrame.Text()))
		if err == nil {
			track.TrackNumber = trackNumber
		}
	}

	lengthFrame, valid := mp3File.Frame("TLEN").(*v2.TextFrame)
	if valid {
		lengthMs, err := strconv.Atoi(TrimNullFromString(lengthFrame.Text()))
		if err == nil {
			track.Length = lengthMs / 1000
		}
	} else {
		log.Fatal("Lengthframe was not a valid cast")
	}

	track.Album.Name = CreateValidNullString(TrimNullFromString(mp3File.Album()))

	imageFrame, valid := mp3File.Frame("APIC").(*v2.ImageFrame)
	if valid {
		track.Album.Image = Image{
			Data:     imageFrame.Data(),
			MimeType: CreateValidNullString(imageFrame.MIMEType()),
		}
	}
	track.Artist = CreateValidNullString(TrimNullFromString(mp3File.Artist()))
	track.Genre = CreateValidNullString(TrimNullFromString(mp3File.Genre()))
	track.Date = CreateValidNullString(TrimNullFromString(mp3File.Year()))

	if !track.Artist.Valid {
		return nil, errors.New("track was missing artist")
	}

	return track, nil
}

func TrimNullFromString(s string) string {
	return strings.Trim(s, "\x00")
}

type Image struct {
	Data     []byte
	MimeType sql.NullString
}

type Album struct {
	Name  sql.NullString
	Image Image
}

type IndexedTrack struct {
	Path        sql.NullString
	Title       sql.NullString
	Artist      sql.NullString
	Album       Album
	AlbumArtist sql.NullString
	TrackNumber int
	Genre       sql.NullString
	Length      int
	Date        sql.NullString
}
