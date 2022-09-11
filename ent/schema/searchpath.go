package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// SearchPath holds the schema definition for the SearchPath entity.
type SearchPath struct {
	ent.Schema
}

// Fields of the SearchPath.
func (SearchPath) Fields() []ent.Field {
	return []ent.Field{
		field.String("path").Immutable(),
	}
}

// Edges of the SearchPath.
func (SearchPath) Edges() []ent.Edge {
	return nil
}
