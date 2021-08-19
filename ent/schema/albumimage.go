package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// AlbumImage holds the schema definition for the AlbumImage entity.
type AlbumImage struct {
	ent.Schema
}

// Fields of the AlbumImage.
func (AlbumImage) Fields() []ent.Field {
	return []ent.Field{
		field.Bytes("image").Immutable(),
		field.String("image_mime_type").Immutable(),
	}
}
