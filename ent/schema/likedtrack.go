package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// LikedTrack holds the schema definition for the LikedTrack entity.
type LikedTrack struct {
	ent.Schema
}

// Fields of the LikedTrack.
func (LikedTrack) Fields() []ent.Field {
	return []ent.Field{
		field.Time("date_added").Default(time.Now().UTC).Immutable(),
	}
}

// Edges of the LikedTrack.
func (LikedTrack) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("track", Track.Type).Unique().Required(),
	}
}
