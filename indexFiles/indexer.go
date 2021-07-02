package indexFiles

import (
	"errors"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/gabriel-vasile/mimetype"
)

type FailedAudioFile struct {
	Path  string
	Error error
}

func ScanPathForMusicFiles(path string) ([]*IndexedTrack, []*FailedAudioFile, error) {
	path, err := filepath.EvalSymlinks(path)
	if err != nil {
		return nil, nil, err
	}

	successfully_parsed_audio_files := []*IndexedTrack{}
	failed_audio_files := []*FailedAudioFile{}

	filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		if err != nil {
			return err
		}

		track, err := parseAudioFile(path)
		if track != nil {
			successfully_parsed_audio_files = append(successfully_parsed_audio_files, track)
		}

		if err != nil {
			failed_audio_files = append(failed_audio_files, &FailedAudioFile{
				Path:  path,
				Error: err,
			})
		}

		return nil
	})

	return successfully_parsed_audio_files, failed_audio_files, nil
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
		} else {
			return nil, errors.New("failed to find image, it's required for tracks to be added to the database")
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
			MimeType: mime.String(),
		}, nil
	}

	return nil, errors.New("failed to find cover image")
}

type Image struct {
	Data     []byte
	MimeType string
}

type IndexedTrack struct {
	Path        string
	Title       string
	Artists     []string
	AlbumName   string
	Image       *Image
	AlbumArtist string
	TrackNumber int
	Genre       string
	Length      int
	Date        string
	MimeType    string
}
