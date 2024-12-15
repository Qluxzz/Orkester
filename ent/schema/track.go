package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
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
		field.Int("length").Immutable(),
		field.String("mimetype").Immutable(),
	}
}

// Edges of the Track.
func (Track) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("artists", Artist.Type).Required(),
		edge.From("album", Album.Type).Ref("tracks").Unique().Required(),
		edge.To("liked", LikedTrack.Type).Unique().Annotations(entsql.Annotation{
			OnDelete: entsql.Cascade,
		}),
	}
}

func (Track) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("track_number", "title").Edges("album").Unique(),
	}
}
