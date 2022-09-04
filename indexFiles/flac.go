package indexFiles

import (
	"log"

	"github.com/mewkiz/flac"
	"github.com/mewkiz/flac/meta"
)

func FlacTryGetEmbeddedImage(path string) *Image {
	f, err := flac.ParseFile(path)
	if err != nil {
		return nil
	}
	defer f.Close()

	for _, block := range f.Blocks {
		switch block.Type {
		case meta.TypePicture:
			data, valid := block.Body.(*meta.Picture)
			if !valid {
				log.Fatalln("Block said it was TypePicture but could not be cast to it!")
			}

			const coverFront = uint32(3)

			if data.Type == coverFront {
				return &Image{
					Data:     data.Data,
					MimeType: data.MIME,
				}
			}
		}
	}

	return nil
}
