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
		field.String("name").Unique(),
		field.Time("birthday").Optional(),
	}
}

// Edges of the Artist.
func (Artist) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("profile_picture", Picture.Type).Unique(),
		edge.To("pictures", Picture.Type),

		edge.From("directed", Movie.Type).Ref("directors"),
		edge.From("acted", Movie.Type).Ref("cast"),
		edge.From("wrote", Movie.Type).Ref("writers"),
		edge.From("countries", Country.Type).Ref("artists"),
	}
}
