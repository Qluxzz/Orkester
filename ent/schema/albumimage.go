package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
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

// Edges of the AlbumImage.
func (AlbumImage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("album", Album.Type).Unique(),
	}
}
