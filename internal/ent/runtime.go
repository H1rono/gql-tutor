// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/h1rono/gql-tutor/internal/ent/repository"
	"github.com/h1rono/gql-tutor/internal/ent/schema"
	"github.com/h1rono/gql-tutor/internal/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	repositoryFields := schema.Repository{}.Fields()
	_ = repositoryFields
	// repositoryDescName is the schema descriptor for name field.
	repositoryDescName := repositoryFields[1].Descriptor()
	// repository.NameValidator is a validator for the "name" field. It is called by the builders before save.
	repository.NameValidator = repositoryDescName.Validators[0].(func(string) error)
	// repositoryDescCreatedAt is the schema descriptor for created_at field.
	repositoryDescCreatedAt := repositoryFields[2].Descriptor()
	// repository.DefaultCreatedAt holds the default value on creation for the created_at field.
	repository.DefaultCreatedAt = repositoryDescCreatedAt.Default.(func() time.Time)
	// repositoryDescID is the schema descriptor for id field.
	repositoryDescID := repositoryFields[0].Descriptor()
	// repository.DefaultID holds the default value on creation for the id field.
	repository.DefaultID = repositoryDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[1].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
