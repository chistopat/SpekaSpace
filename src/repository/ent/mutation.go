// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"
	"time"

	"git.redmadrobot.com/internship/backend/lim-ext/src/repository/ent/predicate"
	"git.redmadrobot.com/internship/backend/lim-ext/src/repository/ent/specification"
	"github.com/google/uuid"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeSpecification = "Specification"
)

// SpecificationMutation represents an operation that mutates the Specification nodes in the graph.
type SpecificationMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	created_at    *time.Time
	name          *string
	description   *string
	author        *string
	status        *specification.Status
	spec          *[]string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Specification, error)
	predicates    []predicate.Specification
}

var _ ent.Mutation = (*SpecificationMutation)(nil)

// specificationOption allows management of the mutation configuration using functional options.
type specificationOption func(*SpecificationMutation)

// newSpecificationMutation creates new mutation for the Specification entity.
func newSpecificationMutation(c config, op Op, opts ...specificationOption) *SpecificationMutation {
	m := &SpecificationMutation{
		config:        c,
		op:            op,
		typ:           TypeSpecification,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withSpecificationID sets the ID field of the mutation.
func withSpecificationID(id uuid.UUID) specificationOption {
	return func(m *SpecificationMutation) {
		var (
			err   error
			once  sync.Once
			value *Specification
		)
		m.oldValue = func(ctx context.Context) (*Specification, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Specification.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withSpecification sets the old Specification of the mutation.
func withSpecification(node *Specification) specificationOption {
	return func(m *SpecificationMutation) {
		m.oldValue = func(context.Context) (*Specification, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m SpecificationMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m SpecificationMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Specification entities.
func (m *SpecificationMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *SpecificationMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetCreatedAt sets the "created_at" field.
func (m *SpecificationMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *SpecificationMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Specification entity.
// If the Specification object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SpecificationMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *SpecificationMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetName sets the "name" field.
func (m *SpecificationMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *SpecificationMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Specification entity.
// If the Specification object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SpecificationMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *SpecificationMutation) ResetName() {
	m.name = nil
}

// SetDescription sets the "description" field.
func (m *SpecificationMutation) SetDescription(s string) {
	m.description = &s
}

// Description returns the value of the "description" field in the mutation.
func (m *SpecificationMutation) Description() (r string, exists bool) {
	v := m.description
	if v == nil {
		return
	}
	return *v, true
}

// OldDescription returns the old "description" field's value of the Specification entity.
// If the Specification object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SpecificationMutation) OldDescription(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldDescription is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldDescription requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDescription: %w", err)
	}
	return oldValue.Description, nil
}

// ResetDescription resets all changes to the "description" field.
func (m *SpecificationMutation) ResetDescription() {
	m.description = nil
}

// SetAuthor sets the "author" field.
func (m *SpecificationMutation) SetAuthor(s string) {
	m.author = &s
}

// Author returns the value of the "author" field in the mutation.
func (m *SpecificationMutation) Author() (r string, exists bool) {
	v := m.author
	if v == nil {
		return
	}
	return *v, true
}

// OldAuthor returns the old "author" field's value of the Specification entity.
// If the Specification object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SpecificationMutation) OldAuthor(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldAuthor is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldAuthor requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAuthor: %w", err)
	}
	return oldValue.Author, nil
}

// ClearAuthor clears the value of the "author" field.
func (m *SpecificationMutation) ClearAuthor() {
	m.author = nil
	m.clearedFields[specification.FieldAuthor] = struct{}{}
}

// AuthorCleared returns if the "author" field was cleared in this mutation.
func (m *SpecificationMutation) AuthorCleared() bool {
	_, ok := m.clearedFields[specification.FieldAuthor]
	return ok
}

// ResetAuthor resets all changes to the "author" field.
func (m *SpecificationMutation) ResetAuthor() {
	m.author = nil
	delete(m.clearedFields, specification.FieldAuthor)
}

// SetStatus sets the "status" field.
func (m *SpecificationMutation) SetStatus(s specification.Status) {
	m.status = &s
}

// Status returns the value of the "status" field in the mutation.
func (m *SpecificationMutation) Status() (r specification.Status, exists bool) {
	v := m.status
	if v == nil {
		return
	}
	return *v, true
}

// OldStatus returns the old "status" field's value of the Specification entity.
// If the Specification object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SpecificationMutation) OldStatus(ctx context.Context) (v specification.Status, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldStatus is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldStatus requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStatus: %w", err)
	}
	return oldValue.Status, nil
}

// ResetStatus resets all changes to the "status" field.
func (m *SpecificationMutation) ResetStatus() {
	m.status = nil
}

// SetSpec sets the "spec" field.
func (m *SpecificationMutation) SetSpec(s []string) {
	m.spec = &s
}

// Spec returns the value of the "spec" field in the mutation.
func (m *SpecificationMutation) Spec() (r []string, exists bool) {
	v := m.spec
	if v == nil {
		return
	}
	return *v, true
}

// OldSpec returns the old "spec" field's value of the Specification entity.
// If the Specification object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SpecificationMutation) OldSpec(ctx context.Context) (v []string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldSpec is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldSpec requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSpec: %w", err)
	}
	return oldValue.Spec, nil
}

// ClearSpec clears the value of the "spec" field.
func (m *SpecificationMutation) ClearSpec() {
	m.spec = nil
	m.clearedFields[specification.FieldSpec] = struct{}{}
}

// SpecCleared returns if the "spec" field was cleared in this mutation.
func (m *SpecificationMutation) SpecCleared() bool {
	_, ok := m.clearedFields[specification.FieldSpec]
	return ok
}

// ResetSpec resets all changes to the "spec" field.
func (m *SpecificationMutation) ResetSpec() {
	m.spec = nil
	delete(m.clearedFields, specification.FieldSpec)
}

// Where appends a list predicates to the SpecificationMutation builder.
func (m *SpecificationMutation) Where(ps ...predicate.Specification) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *SpecificationMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Specification).
func (m *SpecificationMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *SpecificationMutation) Fields() []string {
	fields := make([]string, 0, 6)
	if m.created_at != nil {
		fields = append(fields, specification.FieldCreatedAt)
	}
	if m.name != nil {
		fields = append(fields, specification.FieldName)
	}
	if m.description != nil {
		fields = append(fields, specification.FieldDescription)
	}
	if m.author != nil {
		fields = append(fields, specification.FieldAuthor)
	}
	if m.status != nil {
		fields = append(fields, specification.FieldStatus)
	}
	if m.spec != nil {
		fields = append(fields, specification.FieldSpec)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *SpecificationMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case specification.FieldCreatedAt:
		return m.CreatedAt()
	case specification.FieldName:
		return m.Name()
	case specification.FieldDescription:
		return m.Description()
	case specification.FieldAuthor:
		return m.Author()
	case specification.FieldStatus:
		return m.Status()
	case specification.FieldSpec:
		return m.Spec()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *SpecificationMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case specification.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case specification.FieldName:
		return m.OldName(ctx)
	case specification.FieldDescription:
		return m.OldDescription(ctx)
	case specification.FieldAuthor:
		return m.OldAuthor(ctx)
	case specification.FieldStatus:
		return m.OldStatus(ctx)
	case specification.FieldSpec:
		return m.OldSpec(ctx)
	}
	return nil, fmt.Errorf("unknown Specification field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *SpecificationMutation) SetField(name string, value ent.Value) error {
	switch name {
	case specification.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case specification.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case specification.FieldDescription:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDescription(v)
		return nil
	case specification.FieldAuthor:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAuthor(v)
		return nil
	case specification.FieldStatus:
		v, ok := value.(specification.Status)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStatus(v)
		return nil
	case specification.FieldSpec:
		v, ok := value.([]string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSpec(v)
		return nil
	}
	return fmt.Errorf("unknown Specification field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *SpecificationMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *SpecificationMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *SpecificationMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Specification numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *SpecificationMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(specification.FieldAuthor) {
		fields = append(fields, specification.FieldAuthor)
	}
	if m.FieldCleared(specification.FieldSpec) {
		fields = append(fields, specification.FieldSpec)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *SpecificationMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *SpecificationMutation) ClearField(name string) error {
	switch name {
	case specification.FieldAuthor:
		m.ClearAuthor()
		return nil
	case specification.FieldSpec:
		m.ClearSpec()
		return nil
	}
	return fmt.Errorf("unknown Specification nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *SpecificationMutation) ResetField(name string) error {
	switch name {
	case specification.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case specification.FieldName:
		m.ResetName()
		return nil
	case specification.FieldDescription:
		m.ResetDescription()
		return nil
	case specification.FieldAuthor:
		m.ResetAuthor()
		return nil
	case specification.FieldStatus:
		m.ResetStatus()
		return nil
	case specification.FieldSpec:
		m.ResetSpec()
		return nil
	}
	return fmt.Errorf("unknown Specification field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *SpecificationMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *SpecificationMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *SpecificationMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *SpecificationMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *SpecificationMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *SpecificationMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *SpecificationMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Specification unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *SpecificationMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Specification edge %s", name)
}
