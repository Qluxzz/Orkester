package indexFiles

import (
	"github.com/mikkyang/id3-go"
	v2 "github.com/mikkyang/id3-go/v2"
)

// Info on frames and fields can be found here
// https://id3.org/id3v2.3.0 (2021-05-04)

func Mp3TryGetEmbeddedImage(path string) *Image {
	mp3File, err := id3.Open(path)
	if err != nil {
		return nil
	}

	defer mp3File.Close()

	imageFrame, valid := mp3File.Frame("APIC").(*v2.ImageFrame)
	if valid {
		return &Image{
			Data:     imageFrame.Data(),
			MimeType: imageFrame.MIMEType(),
		}
	}

	return nil
}
