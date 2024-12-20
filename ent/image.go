// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"orkester/ent/image"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Image is the model entity for the Image schema.
type Image struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Image holds the value of the "image" field.
	Image []byte `json:"image,omitempty"`
	// ImageMimeType holds the value of the "image_mime_type" field.
	ImageMimeType string `json:"image_mime_type,omitempty"`
	// Hash holds the value of the "hash" field.
	Hash         string `json:"hash,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Image) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case image.FieldImage:
			values[i] = new([]byte)
		case image.FieldID:
			values[i] = new(sql.NullInt64)
		case image.FieldImageMimeType, image.FieldHash:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Image fields.
func (i *Image) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case image.FieldID:
			value, ok := values[j].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			i.ID = int(value.Int64)
		case image.FieldImage:
			if value, ok := values[j].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field image", values[j])
			} else if value != nil {
				i.Image = *value
			}
		case image.FieldImageMimeType:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image_mime_type", values[j])
			} else if value.Valid {
				i.ImageMimeType = value.String
			}
		case image.FieldHash:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hash", values[j])
			} else if value.Valid {
				i.Hash = value.String
			}
		default:
			i.selectValues.Set(columns[j], values[j])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Image.
// This includes values selected through modifiers, order, etc.
func (i *Image) Value(name string) (ent.Value, error) {
	return i.selectValues.Get(name)
}

// Update returns a builder for updating this Image.
// Note that you need to call Image.Unwrap() before calling this method if this Image
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Image) Update() *ImageUpdateOne {
	return NewImageClient(i.config).UpdateOne(i)
}

// Unwrap unwraps the Image entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Image) Unwrap() *Image {
	_tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Image is not a transactional entity")
	}
	i.config.driver = _tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Image) String() string {
	var builder strings.Builder
	builder.WriteString("Image(")
	builder.WriteString(fmt.Sprintf("id=%v, ", i.ID))
	builder.WriteString("image=")
	builder.WriteString(fmt.Sprintf("%v", i.Image))
	builder.WriteString(", ")
	builder.WriteString("image_mime_type=")
	builder.WriteString(i.ImageMimeType)
	builder.WriteString(", ")
	builder.WriteString("hash=")
	builder.WriteString(i.Hash)
	builder.WriteByte(')')
	return builder.String()
}

// Images is a parsable slice of Image.
type Images []*Image
