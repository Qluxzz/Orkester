package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/mewkiz/flac"
	"github.com/mewkiz/flac/meta"
)

type Image struct {
	Data     []byte
	MimeType string
}

type Album struct {
	Name  string
	Image Image
}

type AddTrackRequest struct {
	Path        string
	Title       string
	Artist      string
	Album       Album
	AlbumArtist string
	TrackNumber string
	Genre       string
	Length      string
	Date        string
}

func IndexFolder(path string) []AddTrackRequest {
	path, err := filepath.EvalSymlinks(path)
	if err != nil {
		log.Fatalln(err)
	}

	tracks := []AddTrackRequest{}

	filepath.Walk(path, func(path string, fileInfo os.FileInfo, err error) error {
		filename := strings.ToLower(fileInfo.Name())
		ext := filepath.Ext(filename)

		if isMusicFile(ext) {
			f, err := flac.ParseFile(path)
			if err != nil {
				return nil
			}
			defer f.Close()

			track := AddTrackRequest{
				Path: path,
			}

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
							track.Artist = value
						case "album":
							track.Album = Album{
								Name: value,
							}
						case "albumartist":
							track.AlbumArtist = value
						case "tracknumber":
							track.TrackNumber = value
						case "genre":
							track.Genre = value
						case "date":
							track.Date = value
						case "length":
							track.Length = value
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
							MimeType: data.MIME,
							Data:     data.Data,
						}
					}
				}
			}

			tracks = append(
				tracks,
				track,
			)

			return nil
		}

		if isCoverImage(filename) {
			return nil
		}

		return nil
	})

	return tracks
}

func isMusicFile(extension string) bool {
	validFileExtensions := []string{".ogg", ".flac", ".mp3"}

	for _, validExtension := range validFileExtensions {
		if extension == validExtension {
			return true
		}
	}

	return false
}

func isCoverImage(filename string) bool {
	hasValidFileName := func() bool {
		validFilenames := []string{"cover", "folder"}

		for _, validFilename := range validFilenames {
			if strings.HasPrefix(filename, validFilename) {
				return true
			}
		}

		return false
	}()

	hasValidExtension := func() bool {
		validFileExtensions := []string{".jpg", ".jpeg", ".png"}

		for _, validFileExtension := range validFileExtensions {
			if strings.HasSuffix(filename, validFileExtension) {
				return true
			}
		}

		return false
	}()

	return hasValidFileName && hasValidExtension
}
