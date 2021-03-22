package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

type AddTrackRequest struct {
	Path string `db:"path"`
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
			// f, err := os.Open(path)
			// if err != nil {
			// 	return nil
			// }

			// for {
			// 	block, err := meta.New(f)

			// 	if err != nil {
			// 		break
			// 	}

			// 	switch block.Type {
			// 	case meta.TypeVorbisComment:
			// 		err := block.Parse()
			// 		if err != nil {
			// 			log.Print(err)
			// 			break
			// 		}

			// 		data, valid := block.Body.(meta.VorbisComment)
			// 		if !valid {
			// 			log.Fatalln("Block said it was TypeVorbisComment but could not be cast to it!")
			// 		}

			// 		for _, tag := range data.Tags {
			// 			log.Print(tag)
			// 		}
			// 	case meta.TypeCueSheet:
			// 	case meta.TypePicture:
			// 	default:
			// 		block.Skip()
			// 	}

			// 	if block.IsLast {
			// 		break
			// 	}
			// }

			tracks = append(
				tracks,
				AddTrackRequest{
					Path: path,
				},
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
