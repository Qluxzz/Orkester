package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Track holds the schema definition for the Track entity.
type Track struct {
	ent.Schema
}

// Fields of the Track.
func (Track) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.Int("track_number"),
		field.String("path"),
		field.Time("date"),
		field.Int("length"),
		field.String("mimetype"),
	}
}

// Edges of the Track.
func (Track) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("artists", Artist.Type).Required(),
		edge.From("album", Album.Type).Ref("tracks").Unique(),
	}
}
