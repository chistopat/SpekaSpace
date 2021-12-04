package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Specification holds the schema definition for the Specification entity.
type Specification struct {
	ent.Schema
}

// Fields of the Specification.
func (Specification) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.String("name").Unique(),
		field.String("author"),
		field.String("status"),
	}

}

// Edges of the Specification.
func (Specification) Edges() []ent.Edge {
	return nil
}
