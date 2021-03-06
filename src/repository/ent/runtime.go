// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"git.redmadrobot.com/internship/backend/lim-ext/src/repository/ent/specification"
	"git.redmadrobot.com/internship/backend/lim-ext/src/repository/schema"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	specificationFields := schema.Specification{}.Fields()
	_ = specificationFields
	// specificationDescCreatedAt is the schema descriptor for created_at field.
	specificationDescCreatedAt := specificationFields[1].Descriptor()
	// specification.DefaultCreatedAt holds the default value on creation for the created_at field.
	specification.DefaultCreatedAt = specificationDescCreatedAt.Default.(func() time.Time)
	// specificationDescUpdatedAt is the schema descriptor for updated_at field.
	specificationDescUpdatedAt := specificationFields[2].Descriptor()
	// specification.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	specification.DefaultUpdatedAt = specificationDescUpdatedAt.Default.(func() time.Time)
	// specificationDescID is the schema descriptor for id field.
	specificationDescID := specificationFields[0].Descriptor()
	// specification.DefaultID holds the default value on creation for the id field.
	specification.DefaultID = specificationDescID.Default.(func() uuid.UUID)
}
