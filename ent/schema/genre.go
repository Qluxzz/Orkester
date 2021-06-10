package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Genre holds the schema definition for the Genre entity.
type Genre struct {
	ent.Schema
}

// Fields of the Genre.
func (Genre) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Immutable(),
		field.String("url_name").Immutable(),
	}
}

// Edges of the Genre.
func (Genre) Edges() []ent.Edge {
	return nil
}
