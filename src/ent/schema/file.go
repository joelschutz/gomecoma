package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// File holds the schema definition for the File entity.
type File struct {
	ent.Schema
}

// Fields of the File.
func (File) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("path"),
		field.Enum("type").Values("audio", "video", "image"),
		field.String("external_id").Optional(),
		field.String("external_info_provider").Optional(),
		field.String("results").Optional(),
		field.Bool("synced").Default(false),
	}
}

// Edges of the File.
func (File) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("movie", Movie.Type).Ref("file").Unique(),
	}
}
