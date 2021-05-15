package indexFiles

import (
	"database/sql"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/gabriel-vasile/mimetype"
)

func ScanPathForMusicFiles(path string) ([]*IndexedTrack, error) {
	path, err := filepath.EvalSymlinks(path)
	if err != nil {
		return nil, err
	}

	tracks := []*IndexedTrack{}

	filepath.Walk(path, func(path string, fileInfo os.FileInfo, err error) error {
		track, _ := parseAudioFile(path)
		if track != nil {
			tracks = append(tracks, track)
		}

		return nil
	})

	return tracks, nil
}

func parseAudioFile(path string) (*IndexedTrack, error) {
	var track *IndexedTrack
	var err error

	switch filepath.Ext(path) {
	case ".flac":
		track, err = ParseFlacFile(path)
	case ".mp3":
		track, err = ParseMp3File(path)
	default:
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	if track.Image == nil {
		image, err := scanFolderForCoverImage(filepath.Dir(path))
		if err == nil {
			track.Image = image
		}
	}

	return track, nil
}

func scanFolderForCoverImage(path string) (*Image, error) {
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

type Image struct {
	Data     []byte
	MimeType sql.NullString
}

type IndexedTrack struct {
	Path        sql.NullString
	Title       sql.NullString
	ArtistName  sql.NullString
	AlbumName   sql.NullString
	Image       *Image
	AlbumArtist sql.NullString
	TrackNumber sql.NullInt32
	Genre       sql.NullString
	Length      sql.NullInt32
	Date        sql.NullString
	MimeType    sql.NullString
}
