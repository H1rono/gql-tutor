package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Repository holds the schema definition for the Repository entity.
type Repository struct {
	ent.Schema
}

// Fields of the Repository.
func (Repository) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Unique().Immutable().Default(uuid.New),
		field.String("name").NotEmpty(),
		field.Time("created_at").Immutable().Default(time.Now),
		field.UUID("owner_id", uuid.UUID{}).Immutable(),
	}
}

// Edges of the Repository.
func (Repository) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("repositories").Field("owner_id").Unique().Immutable().Required(),
	}
}
