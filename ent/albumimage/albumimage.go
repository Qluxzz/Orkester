// Code generated by entc, DO NOT EDIT.

package albumimage

const (
	// Label holds the string label denoting the albumimage type in the database.
	Label = "album_image"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldImage holds the string denoting the image field in the database.
	FieldImage = "image"
	// FieldImageMimeType holds the string denoting the image_mime_type field in the database.
	FieldImageMimeType = "image_mime_type"
	// EdgeAlbum holds the string denoting the album edge name in mutations.
	EdgeAlbum = "album"
	// Table holds the table name of the albumimage in the database.
	Table = "album_images"
	// AlbumTable is the table the holds the album relation/edge.
	AlbumTable = "album_images"
	// AlbumInverseTable is the table name for the Album entity.
	// It exists in this package in order to avoid circular dependency with the "album" package.
	AlbumInverseTable = "albums"
	// AlbumColumn is the table column denoting the album relation/edge.
	AlbumColumn = "album_image_album"
)

// Columns holds all SQL columns for albumimage fields.
var Columns = []string{
	FieldID,
	FieldImage,
	FieldImageMimeType,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "album_images"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"album_image_album",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}
