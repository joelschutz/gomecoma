package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// MovieGenre holds the schema definition for the MovieGenre entity.
type MovieGenre struct {
	ent.Schema
}

// Fields of the MovieGenre.
func (MovieGenre) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the MovieGenre.
func (MovieGenre) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("movies", Movie.Type),
	}
}
