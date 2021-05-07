package indexFiles

import (
	"database/sql"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gabriel-vasile/mimetype"
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
	track.Length = CreateValidNullInt(int(f.Info.NSamples) / int(f.Info.SampleRate))

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
				track.Album.Image = Image{
					Data:     data.Data,
					MimeType: CreateValidNullString(data.MIME),
				}
			}
		}
	}

	if !track.Album.Image.MimeType.Valid {
		image, err := ScanFolderForCoverImage(filepath.Dir(path))
		if err == nil {
			track.Album.Image = *image
		}
	}

	if !track.Artist.Valid {
		return nil, errors.New("track was missing artist")
	}

	return track, nil
}

func ScanFolderForCoverImage(path string) (*Image, error) {
	validImages := []string{}

	filepath.Walk(path, func(currentPath string, fileInfo os.FileInfo, err error) error {
		if fileInfo.IsDir() {
			return nil
		}

		ext := filepath.Ext(currentPath)

		hasValidExtension := func(ext string) bool {
			validExtensions := []string{
				".png",
				".jpg",
			}

			lowerExt := strings.ToLower(ext)

			for _, validExtension := range validExtensions {
				if lowerExt == validExtension {
					return true
				}
			}

			return false
		}(ext)

		fileName := fileInfo.Name()

		filenameWithoutExtension := fileName[0 : len(fileName)-len(ext)]

		hasValidFileName := func(filename string) bool {
			validFileNames := []string{
				"cover",
				"folder",
			}

			loweredFileName := strings.ToLower(filename)

			for _, validFileName := range validFileNames {
				if loweredFileName == validFileName {
					return true
				}
			}

			return false
		}(filenameWithoutExtension)

		if hasValidExtension && hasValidFileName {
			validImages = append(validImages, currentPath)
		}

		return nil
	})

	if len(validImages) > 0 {
		data, err := ioutil.ReadFile(validImages[0])
		if err != nil {
			return nil, err
		}

		mime := mimetype.Detect(data)
		if mime == nil {
			return nil, errors.New("failed to get image mimetype")
		}

		return &Image{
			Data:     data,
			MimeType: CreateValidNullString(mime.String()),
		}, nil
	}

	return nil, errors.New("failed to find cover image")
}

func CreateValidNullString(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func CreateValidNullInt(n int) sql.NullInt32 {
	return sql.NullInt32{
		Int32: int32(n),
		Valid: true,
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
			track.TrackNumber = CreateValidNullInt(trackNumber)
		}
	}

	lengthFrame, valid := mp3File.Frame("TLEN").(*v2.TextFrame)
	if valid {
		lengthMs, err := strconv.Atoi(TrimNullFromString(lengthFrame.Text()))
		if err == nil {
			track.Length = CreateValidNullInt(lengthMs / 1000)
		}
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
	TrackNumber sql.NullInt32
	Genre       sql.NullString
	Length      sql.NullInt32
	Date        sql.NullString
}
