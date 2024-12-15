package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Image holds the schema definition for the Image entity.
type Image struct {
	ent.Schema
}

// Fields of the Image.
func (Image) Fields() []ent.Field {
	return []ent.Field{
		field.Bytes("image").Immutable(),
		field.String("image_mime_type").Immutable(),
		field.String("hash").Immutable(),
	}
}
