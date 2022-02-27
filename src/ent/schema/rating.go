package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Rating holds the schema definition for the Rating entity.
type Rating struct {
	ent.Schema
}

// Fields of the Rating.
func (Rating) Fields() []ent.Field {
	return []ent.Field{
		field.String("origin"),
		field.String("original_rating"),
		field.Int("normalized_rating").Max(5).NonNegative(),
	}
}

// Edges of the Rating.
func (Rating) Edges() []ent.Edge {
	return nil
}
