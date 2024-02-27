// Code generated by ent, DO NOT EDIT.

package ent

import (
	"clove-api/domain/image/domain/ent/image"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ImageCreate is the builder for creating a Image entity.
type ImageCreate struct {
	config
	mutation *ImageMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ic *ImageCreate) SetName(s string) *ImageCreate {
	ic.mutation.SetName(s)
	return ic
}

// SetDescription sets the "description" field.
func (ic *ImageCreate) SetDescription(s string) *ImageCreate {
	ic.mutation.SetDescription(s)
	return ic
}

// SetFileName sets the "fileName" field.
func (ic *ImageCreate) SetFileName(s string) *ImageCreate {
	ic.mutation.SetFileName(s)
	return ic
}

// SetURL sets the "url" field.
func (ic *ImageCreate) SetURL(s string) *ImageCreate {
	ic.mutation.SetURL(s)
	return ic
}

// SetSize sets the "size" field.
func (ic *ImageCreate) SetSize(s string) *ImageCreate {
	ic.mutation.SetSize(s)
	return ic
}

// SetInstance sets the "instance" field.
func (ic *ImageCreate) SetInstance(b []byte) *ImageCreate {
	ic.mutation.SetInstance(b)
	return ic
}

// SetCreatedAt sets the "createdAt" field.
func (ic *ImageCreate) SetCreatedAt(t time.Time) *ImageCreate {
	ic.mutation.SetCreatedAt(t)
	return ic
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (ic *ImageCreate) SetNillableCreatedAt(t *time.Time) *ImageCreate {
	if t != nil {
		ic.SetCreatedAt(*t)
	}
	return ic
}

// Mutation returns the ImageMutation object of the builder.
func (ic *ImageCreate) Mutation() *ImageMutation {
	return ic.mutation
}

// Save creates the Image in the database.
func (ic *ImageCreate) Save(ctx context.Context) (*Image, error) {
	ic.defaults()
	return withHooks(ctx, ic.sqlSave, ic.mutation, ic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ic *ImageCreate) SaveX(ctx context.Context) *Image {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *ImageCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *ImageCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ic *ImageCreate) defaults() {
	if _, ok := ic.mutation.CreatedAt(); !ok {
		v := image.DefaultCreatedAt()
		ic.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *ImageCreate) check() error {
	if _, ok := ic.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Image.name"`)}
	}
	if _, ok := ic.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Image.description"`)}
	}
	if _, ok := ic.mutation.FileName(); !ok {
		return &ValidationError{Name: "fileName", err: errors.New(`ent: missing required field "Image.fileName"`)}
	}
	if _, ok := ic.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`ent: missing required field "Image.url"`)}
	}
	if _, ok := ic.mutation.Size(); !ok {
		return &ValidationError{Name: "size", err: errors.New(`ent: missing required field "Image.size"`)}
	}
	if _, ok := ic.mutation.Instance(); !ok {
		return &ValidationError{Name: "instance", err: errors.New(`ent: missing required field "Image.instance"`)}
	}
	if v, ok := ic.mutation.Instance(); ok {
		if err := image.InstanceValidator(v); err != nil {
			return &ValidationError{Name: "instance", err: fmt.Errorf(`ent: validator failed for field "Image.instance": %w`, err)}
		}
	}
	if _, ok := ic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "Image.createdAt"`)}
	}
	return nil
}

func (ic *ImageCreate) sqlSave(ctx context.Context) (*Image, error) {
	if err := ic.check(); err != nil {
		return nil, err
	}
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ic.mutation.id = &_node.ID
	ic.mutation.done = true
	return _node, nil
}

func (ic *ImageCreate) createSpec() (*Image, *sqlgraph.CreateSpec) {
	var (
		_node = &Image{config: ic.config}
		_spec = sqlgraph.NewCreateSpec(image.Table, sqlgraph.NewFieldSpec(image.FieldID, field.TypeInt))
	)
	if value, ok := ic.mutation.Name(); ok {
		_spec.SetField(image.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ic.mutation.Description(); ok {
		_spec.SetField(image.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := ic.mutation.FileName(); ok {
		_spec.SetField(image.FieldFileName, field.TypeString, value)
		_node.FileName = value
	}
	if value, ok := ic.mutation.URL(); ok {
		_spec.SetField(image.FieldURL, field.TypeString, value)
		_node.URL = value
	}
	if value, ok := ic.mutation.Size(); ok {
		_spec.SetField(image.FieldSize, field.TypeString, value)
		_node.Size = value
	}
	if value, ok := ic.mutation.Instance(); ok {
		_spec.SetField(image.FieldInstance, field.TypeBytes, value)
		_node.Instance = value
	}
	if value, ok := ic.mutation.CreatedAt(); ok {
		_spec.SetField(image.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	return _node, _spec
}

// ImageCreateBulk is the builder for creating many Image entities in bulk.
type ImageCreateBulk struct {
	config
	err      error
	builders []*ImageCreate
}

// Save creates the Image entities in the database.
func (icb *ImageCreateBulk) Save(ctx context.Context) ([]*Image, error) {
	if icb.err != nil {
		return nil, icb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Image, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ImageMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *ImageCreateBulk) SaveX(ctx context.Context) []*Image {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *ImageCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *ImageCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}
