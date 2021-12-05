package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

const (
	StatusDraft = "DRAFT"
	StatusNew   = "NEW"
)

// Specification holds the schema definition for the Specification entity.
type Specification struct {
	ent.Schema
}

// Fields of the Specification.
func (Specification) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now),
		field.String("name").Unique(),
		field.String("description"),
		field.String("author").Optional(),
		field.Enum("status").Values(StatusDraft, StatusNew).Default(StatusDraft),
		field.JSON("spec", []string{}).Optional(),
	}

}

// Edges of the Specification.
func (Specification) Edges() []ent.Edge {
	return nil
}
