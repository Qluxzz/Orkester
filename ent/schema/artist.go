package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Artist holds the schema definition for the Artist entity.
type Artist struct {
	ent.Schema
}

// Fields of the Artist.
func (Artist) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique().Immutable(),
		field.String("url_name").Immutable(),
	}
}

// Edges of the Artist.
func (Artist) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("albums", Album.Type),
		edge.From("track", Track.Type).Ref("artists"),
	}
}
