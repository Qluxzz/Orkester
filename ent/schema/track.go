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
		field.String("title").Immutable(),
		field.Int("track_number").Immutable(),
		field.String("path").Immutable(),
		field.Time("date").Immutable(),
		field.Int("length").Immutable(),
		field.String("mimetype").Immutable(),
	}
}

// Edges of the Track.
func (Track) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("artists", Artist.Type).Required(),
		edge.From("album", Album.Type).Ref("tracks").Unique(),
	}
}