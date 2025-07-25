// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/h1rono/gql-tutor/internal/ent/repository"
	"github.com/h1rono/gql-tutor/internal/ent/user"
)

// Repository is the model entity for the Repository schema.
type Repository struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// OwnerID holds the value of the "owner_id" field.
	OwnerID uuid.UUID `json:"owner_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RepositoryQuery when eager-loading is set.
	Edges        RepositoryEdges `json:"edges"`
	selectValues sql.SelectValues
}

// RepositoryEdges holds the relations/edges for other nodes in the graph.
type RepositoryEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RepositoryEdges) OwnerOrErr() (*User, error) {
	if e.Owner != nil {
		return e.Owner, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Repository) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case repository.FieldName:
			values[i] = new(sql.NullString)
		case repository.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case repository.FieldID, repository.FieldOwnerID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Repository fields.
func (r *Repository) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case repository.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				r.ID = *value
			}
		case repository.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				r.Name = value.String
			}
		case repository.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				r.CreatedAt = value.Time
			}
		case repository.FieldOwnerID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field owner_id", values[i])
			} else if value != nil {
				r.OwnerID = *value
			}
		default:
			r.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Repository.
// This includes values selected through modifiers, order, etc.
func (r *Repository) Value(name string) (ent.Value, error) {
	return r.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the Repository entity.
func (r *Repository) QueryOwner() *UserQuery {
	return NewRepositoryClient(r.config).QueryOwner(r)
}

// Update returns a builder for updating this Repository.
// Note that you need to call Repository.Unwrap() before calling this method if this Repository
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Repository) Update() *RepositoryUpdateOne {
	return NewRepositoryClient(r.config).UpdateOne(r)
}

// Unwrap unwraps the Repository entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Repository) Unwrap() *Repository {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Repository is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Repository) String() string {
	var builder strings.Builder
	builder.WriteString("Repository(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("name=")
	builder.WriteString(r.Name)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(r.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("owner_id=")
	builder.WriteString(fmt.Sprintf("%v", r.OwnerID))
	builder.WriteByte(')')
	return builder.String()
}

// Repositories is a parsable slice of Repository.
type Repositories []*Repository
