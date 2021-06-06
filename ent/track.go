// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"goreact/ent/album"
	"goreact/ent/track"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Track is the model entity for the Track schema.
type Track struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// TrackNumber holds the value of the "track_number" field.
	TrackNumber int `json:"track_number,omitempty"`
	// Path holds the value of the "path" field.
	Path string `json:"path,omitempty"`
	// Date holds the value of the "date" field.
	Date time.Time `json:"date,omitempty"`
	// Length holds the value of the "length" field.
	Length int `json:"length,omitempty"`
	// Mimetype holds the value of the "mimetype" field.
	Mimetype string `json:"mimetype,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TrackQuery when eager-loading is set.
	Edges        TrackEdges `json:"edges"`
	album_tracks *int
}

// TrackEdges holds the relations/edges for other nodes in the graph.
type TrackEdges struct {
	// Artists holds the value of the artists edge.
	Artists []*Artist `json:"artists,omitempty"`
	// Album holds the value of the album edge.
	Album *Album `json:"album,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ArtistsOrErr returns the Artists value or an error if the edge
// was not loaded in eager-loading.
func (e TrackEdges) ArtistsOrErr() ([]*Artist, error) {
	if e.loadedTypes[0] {
		return e.Artists, nil
	}
	return nil, &NotLoadedError{edge: "artists"}
}

// AlbumOrErr returns the Album value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TrackEdges) AlbumOrErr() (*Album, error) {
	if e.loadedTypes[1] {
		if e.Album == nil {
			// The edge album was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: album.Label}
		}
		return e.Album, nil
	}
	return nil, &NotLoadedError{edge: "album"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Track) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case track.FieldID, track.FieldTrackNumber, track.FieldLength:
			values[i] = new(sql.NullInt64)
		case track.FieldTitle, track.FieldPath, track.FieldMimetype:
			values[i] = new(sql.NullString)
		case track.FieldDate:
			values[i] = new(sql.NullTime)
		case track.ForeignKeys[0]: // album_tracks
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Track", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Track fields.
func (t *Track) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case track.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case track.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				t.Title = value.String
			}
		case track.FieldTrackNumber:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field track_number", values[i])
			} else if value.Valid {
				t.TrackNumber = int(value.Int64)
			}
		case track.FieldPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field path", values[i])
			} else if value.Valid {
				t.Path = value.String
			}
		case track.FieldDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field date", values[i])
			} else if value.Valid {
				t.Date = value.Time
			}
		case track.FieldLength:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field length", values[i])
			} else if value.Valid {
				t.Length = int(value.Int64)
			}
		case track.FieldMimetype:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mimetype", values[i])
			} else if value.Valid {
				t.Mimetype = value.String
			}
		case track.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field album_tracks", value)
			} else if value.Valid {
				t.album_tracks = new(int)
				*t.album_tracks = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryArtists queries the "artists" edge of the Track entity.
func (t *Track) QueryArtists() *ArtistQuery {
	return (&TrackClient{config: t.config}).QueryArtists(t)
}

// QueryAlbum queries the "album" edge of the Track entity.
func (t *Track) QueryAlbum() *AlbumQuery {
	return (&TrackClient{config: t.config}).QueryAlbum(t)
}

// Update returns a builder for updating this Track.
// Note that you need to call Track.Unwrap() before calling this method if this Track
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Track) Update() *TrackUpdateOne {
	return (&TrackClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Track entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Track) Unwrap() *Track {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Track is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Track) String() string {
	var builder strings.Builder
	builder.WriteString("Track(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", title=")
	builder.WriteString(t.Title)
	builder.WriteString(", track_number=")
	builder.WriteString(fmt.Sprintf("%v", t.TrackNumber))
	builder.WriteString(", path=")
	builder.WriteString(t.Path)
	builder.WriteString(", date=")
	builder.WriteString(t.Date.Format(time.ANSIC))
	builder.WriteString(", length=")
	builder.WriteString(fmt.Sprintf("%v", t.Length))
	builder.WriteString(", mimetype=")
	builder.WriteString(t.Mimetype)
	builder.WriteByte(')')
	return builder.String()
}

// Tracks is a parsable slice of Track.
type Tracks []*Track

func (t Tracks) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
