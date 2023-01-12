package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Movie holds the schema definition for the Movie entity.
type Movie struct {
	ent.Schema
}

// Fields of the Movie.
func (Movie) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("original_title").Optional(),
		// field.Strings("languages").Optional(),
		field.Time("release_date").Optional(),
		field.String("plot").Optional(),
		field.Int("duration").Optional(),
		field.Bool("watched").Default(false),
	}
}

// Edges of the Movie.
func (Movie) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("file", File.Type).Unique(),
		edge.To("ratings", Rating.Type),
		edge.To("poster", Picture.Type).Unique(),
		edge.To("fanart", Picture.Type),
		edge.To("cast", Artist.Type),
		edge.To("directors", Artist.Type),
		edge.To("writers", Artist.Type),

		edge.From("genres", MovieGenre.Type).Ref("movies"),
		edge.From("countries", Country.Type).Ref("movies"),
	}
}
