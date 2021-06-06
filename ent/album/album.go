// Code generated by entc, DO NOT EDIT.

package album

const (
	// Label holds the string label denoting the album type in the database.
	Label = "album"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldURLName holds the string denoting the url_name field in the database.
	FieldURLName = "url_name"
	// FieldImage holds the string denoting the image field in the database.
	FieldImage = "image"
	// FieldImageMimeType holds the string denoting the image_mime_type field in the database.
	FieldImageMimeType = "image_mime_type"
	// EdgeArtist holds the string denoting the artist edge name in mutations.
	EdgeArtist = "artist"
	// EdgeTracks holds the string denoting the tracks edge name in mutations.
	EdgeTracks = "tracks"
	// Table holds the table name of the album in the database.
	Table = "albums"
	// ArtistTable is the table the holds the artist relation/edge.
	ArtistTable = "albums"
	// ArtistInverseTable is the table name for the Artist entity.
	// It exists in this package in order to avoid circular dependency with the "artist" package.
	ArtistInverseTable = "artists"
	// ArtistColumn is the table column denoting the artist relation/edge.
	ArtistColumn = "artist_albums"
	// TracksTable is the table the holds the tracks relation/edge.
	TracksTable = "tracks"
	// TracksInverseTable is the table name for the Track entity.
	// It exists in this package in order to avoid circular dependency with the "track" package.
	TracksInverseTable = "tracks"
	// TracksColumn is the table column denoting the tracks relation/edge.
	TracksColumn = "album_tracks"
)

// Columns holds all SQL columns for album fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldURLName,
	FieldImage,
	FieldImageMimeType,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "albums"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"artist_albums",
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
