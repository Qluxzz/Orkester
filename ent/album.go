// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"goreact/ent/album"
	"goreact/ent/albumimage"
	"goreact/ent/artist"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Album is the model entity for the Album schema.
type Album struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// URLName holds the value of the "url_name" field.
	URLName string `json:"url_name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AlbumQuery when eager-loading is set.
	Edges             AlbumEdges `json:"edges"`
	album_album_image *int
	artist_albums     *int
}

// AlbumEdges holds the relations/edges for other nodes in the graph.
type AlbumEdges struct {
	// Artist holds the value of the artist edge.
	Artist *Artist `json:"artist,omitempty"`
	// Tracks holds the value of the tracks edge.
	Tracks []*Track `json:"tracks,omitempty"`
	// AlbumImage holds the value of the album_image edge.
	AlbumImage *AlbumImage `json:"album_image,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ArtistOrErr returns the Artist value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AlbumEdges) ArtistOrErr() (*Artist, error) {
	if e.loadedTypes[0] {
		if e.Artist == nil {
			// The edge artist was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: artist.Label}
		}
		return e.Artist, nil
	}
	return nil, &NotLoadedError{edge: "artist"}
}

// TracksOrErr returns the Tracks value or an error if the edge
// was not loaded in eager-loading.
func (e AlbumEdges) TracksOrErr() ([]*Track, error) {
	if e.loadedTypes[1] {
		return e.Tracks, nil
	}
	return nil, &NotLoadedError{edge: "tracks"}
}

// AlbumImageOrErr returns the AlbumImage value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AlbumEdges) AlbumImageOrErr() (*AlbumImage, error) {
	if e.loadedTypes[2] {
		if e.AlbumImage == nil {
			// The edge album_image was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: albumimage.Label}
		}
		return e.AlbumImage, nil
	}
	return nil, &NotLoadedError{edge: "album_image"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Album) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case album.FieldID:
			values[i] = new(sql.NullInt64)
		case album.FieldName, album.FieldURLName:
			values[i] = new(sql.NullString)
		case album.ForeignKeys[0]: // album_album_image
			values[i] = new(sql.NullInt64)
		case album.ForeignKeys[1]: // artist_albums
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Album", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Album fields.
func (a *Album) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case album.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case album.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				a.Name = value.String
			}
		case album.FieldURLName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url_name", values[i])
			} else if value.Valid {
				a.URLName = value.String
			}
		case album.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field album_album_image", value)
			} else if value.Valid {
				a.album_album_image = new(int)
				*a.album_album_image = int(value.Int64)
			}
		case album.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field artist_albums", value)
			} else if value.Valid {
				a.artist_albums = new(int)
				*a.artist_albums = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryArtist queries the "artist" edge of the Album entity.
func (a *Album) QueryArtist() *ArtistQuery {
	return (&AlbumClient{config: a.config}).QueryArtist(a)
}

// QueryTracks queries the "tracks" edge of the Album entity.
func (a *Album) QueryTracks() *TrackQuery {
	return (&AlbumClient{config: a.config}).QueryTracks(a)
}

// QueryAlbumImage queries the "album_image" edge of the Album entity.
func (a *Album) QueryAlbumImage() *AlbumImageQuery {
	return (&AlbumClient{config: a.config}).QueryAlbumImage(a)
}

// Update returns a builder for updating this Album.
// Note that you need to call Album.Unwrap() before calling this method if this Album
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Album) Update() *AlbumUpdateOne {
	return (&AlbumClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Album entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Album) Unwrap() *Album {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Album is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Album) String() string {
	var builder strings.Builder
	builder.WriteString("Album(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", name=")
	builder.WriteString(a.Name)
	builder.WriteString(", url_name=")
	builder.WriteString(a.URLName)
	builder.WriteByte(')')
	return builder.String()
}

// Albums is a parsable slice of Album.
type Albums []*Album

func (a Albums) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
