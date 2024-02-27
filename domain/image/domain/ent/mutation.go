// Code generated by ent, DO NOT EDIT.

package ent

import (
	"clove-api/domain/image/domain/ent/image"
	"clove-api/domain/image/domain/ent/predicate"
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeImage = "Image"
)

// ImageMutation represents an operation that mutates the Image nodes in the graph.
type ImageMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	description   *string
	fileName      *string
	url           *string
	size          *string
	instance      *[]byte
	createdAt     *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Image, error)
	predicates    []predicate.Image
}

var _ ent.Mutation = (*ImageMutation)(nil)

// imageOption allows management of the mutation configuration using functional options.
type imageOption func(*ImageMutation)

// newImageMutation creates new mutation for the Image entity.
func newImageMutation(c config, op Op, opts ...imageOption) *ImageMutation {
	m := &ImageMutation{
		config:        c,
		op:            op,
		typ:           TypeImage,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withImageID sets the ID field of the mutation.
func withImageID(id int) imageOption {
	return func(m *ImageMutation) {
		var (
			err   error
			once  sync.Once
			value *Image
		)
		m.oldValue = func(ctx context.Context) (*Image, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Image.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withImage sets the old Image of the mutation.
func withImage(node *Image) imageOption {
	return func(m *ImageMutation) {
		m.oldValue = func(context.Context) (*Image, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ImageMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ImageMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *ImageMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *ImageMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Image.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *ImageMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *ImageMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Image entity.
// If the Image object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ImageMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *ImageMutation) ResetName() {
	m.name = nil
}

// SetDescription sets the "description" field.
func (m *ImageMutation) SetDescription(s string) {
	m.description = &s
}

// Description returns the value of the "description" field in the mutation.
func (m *ImageMutation) Description() (r string, exists bool) {
	v := m.description
	if v == nil {
		return
	}
	return *v, true
}

// OldDescription returns the old "description" field's value of the Image entity.
// If the Image object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ImageMutation) OldDescription(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDescription is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDescription requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDescription: %w", err)
	}
	return oldValue.Description, nil
}

// ResetDescription resets all changes to the "description" field.
func (m *ImageMutation) ResetDescription() {
	m.description = nil
}

// SetFileName sets the "fileName" field.
func (m *ImageMutation) SetFileName(s string) {
	m.fileName = &s
}

// FileName returns the value of the "fileName" field in the mutation.
func (m *ImageMutation) FileName() (r string, exists bool) {
	v := m.fileName
	if v == nil {
		return
	}
	return *v, true
}

// OldFileName returns the old "fileName" field's value of the Image entity.
// If the Image object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ImageMutation) OldFileName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldFileName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldFileName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldFileName: %w", err)
	}
	return oldValue.FileName, nil
}

// ResetFileName resets all changes to the "fileName" field.
func (m *ImageMutation) ResetFileName() {
	m.fileName = nil
}

// SetURL sets the "url" field.
func (m *ImageMutation) SetURL(s string) {
	m.url = &s
}

// URL returns the value of the "url" field in the mutation.
func (m *ImageMutation) URL() (r string, exists bool) {
	v := m.url
	if v == nil {
		return
	}
	return *v, true
}

// OldURL returns the old "url" field's value of the Image entity.
// If the Image object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ImageMutation) OldURL(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldURL is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldURL requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldURL: %w", err)
	}
	return oldValue.URL, nil
}

// ResetURL resets all changes to the "url" field.
func (m *ImageMutation) ResetURL() {
	m.url = nil
}

// SetSize sets the "size" field.
func (m *ImageMutation) SetSize(s string) {
	m.size = &s
}

// Size returns the value of the "size" field in the mutation.
func (m *ImageMutation) Size() (r string, exists bool) {
	v := m.size
	if v == nil {
		return
	}
	return *v, true
}

// OldSize returns the old "size" field's value of the Image entity.
// If the Image object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ImageMutation) OldSize(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldSize is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldSize requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSize: %w", err)
	}
	return oldValue.Size, nil
}

// ResetSize resets all changes to the "size" field.
func (m *ImageMutation) ResetSize() {
	m.size = nil
}

// SetInstance sets the "instance" field.
func (m *ImageMutation) SetInstance(b []byte) {
	m.instance = &b
}

// Instance returns the value of the "instance" field in the mutation.
func (m *ImageMutation) Instance() (r []byte, exists bool) {
	v := m.instance
	if v == nil {
		return
	}
	return *v, true
}

// OldInstance returns the old "instance" field's value of the Image entity.
// If the Image object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ImageMutation) OldInstance(ctx context.Context) (v []byte, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldInstance is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldInstance requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldInstance: %w", err)
	}
	return oldValue.Instance, nil
}

// ResetInstance resets all changes to the "instance" field.
func (m *ImageMutation) ResetInstance() {
	m.instance = nil
}

// SetCreatedAt sets the "createdAt" field.
func (m *ImageMutation) SetCreatedAt(t time.Time) {
	m.createdAt = &t
}

// CreatedAt returns the value of the "createdAt" field in the mutation.
func (m *ImageMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.createdAt
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "createdAt" field's value of the Image entity.
// If the Image object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ImageMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "createdAt" field.
func (m *ImageMutation) ResetCreatedAt() {
	m.createdAt = nil
}

// Where appends a list predicates to the ImageMutation builder.
func (m *ImageMutation) Where(ps ...predicate.Image) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the ImageMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *ImageMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Image, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *ImageMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *ImageMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Image).
func (m *ImageMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ImageMutation) Fields() []string {
	fields := make([]string, 0, 7)
	if m.name != nil {
		fields = append(fields, image.FieldName)
	}
	if m.description != nil {
		fields = append(fields, image.FieldDescription)
	}
	if m.fileName != nil {
		fields = append(fields, image.FieldFileName)
	}
	if m.url != nil {
		fields = append(fields, image.FieldURL)
	}
	if m.size != nil {
		fields = append(fields, image.FieldSize)
	}
	if m.instance != nil {
		fields = append(fields, image.FieldInstance)
	}
	if m.createdAt != nil {
		fields = append(fields, image.FieldCreatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ImageMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case image.FieldName:
		return m.Name()
	case image.FieldDescription:
		return m.Description()
	case image.FieldFileName:
		return m.FileName()
	case image.FieldURL:
		return m.URL()
	case image.FieldSize:
		return m.Size()
	case image.FieldInstance:
		return m.Instance()
	case image.FieldCreatedAt:
		return m.CreatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ImageMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case image.FieldName:
		return m.OldName(ctx)
	case image.FieldDescription:
		return m.OldDescription(ctx)
	case image.FieldFileName:
		return m.OldFileName(ctx)
	case image.FieldURL:
		return m.OldURL(ctx)
	case image.FieldSize:
		return m.OldSize(ctx)
	case image.FieldInstance:
		return m.OldInstance(ctx)
	case image.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Image field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ImageMutation) SetField(name string, value ent.Value) error {
	switch name {
	case image.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case image.FieldDescription:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDescription(v)
		return nil
	case image.FieldFileName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetFileName(v)
		return nil
	case image.FieldURL:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetURL(v)
		return nil
	case image.FieldSize:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSize(v)
		return nil
	case image.FieldInstance:
		v, ok := value.([]byte)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetInstance(v)
		return nil
	case image.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Image field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ImageMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ImageMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ImageMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Image numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ImageMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ImageMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ImageMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Image nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ImageMutation) ResetField(name string) error {
	switch name {
	case image.FieldName:
		m.ResetName()
		return nil
	case image.FieldDescription:
		m.ResetDescription()
		return nil
	case image.FieldFileName:
		m.ResetFileName()
		return nil
	case image.FieldURL:
		m.ResetURL()
		return nil
	case image.FieldSize:
		m.ResetSize()
		return nil
	case image.FieldInstance:
		m.ResetInstance()
		return nil
	case image.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	}
	return fmt.Errorf("unknown Image field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ImageMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ImageMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ImageMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ImageMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ImageMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ImageMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ImageMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Image unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ImageMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Image edge %s", name)
}
