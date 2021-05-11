package models

type AlbumImage struct {
	Image    []byte `db:"image"`
	MimeType string `db:"imagemimetype"`
}
