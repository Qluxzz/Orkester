package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Album holds the schema definition for the Album entity.
type Album struct {
	ent.Schema
}

// Fields of the Album.
func (Album) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("url_name"),
		field.Bytes("image"),
		field.String("image_mime_type"),
	}
}

// Edges of the Album.
func (Album) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("artist", Artist.Type).Ref("albums").Unique().Required(),
		edge.To("tracks", Track.Type),
	}
}

func (Album) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").Edges("artist").Unique(),
	}
}
