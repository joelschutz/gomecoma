package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Picture holds the schema definition for the Picture entity.
type Picture struct {
	ent.Schema
}

// Fields of the Picture.
func (Picture) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional(),
		field.String("filename"),
		field.String("path"),
	}
}

// Edges of the Picture.
func (Picture) Edges() []ent.Edge {
	return nil
}
