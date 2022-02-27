// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/joelschutz/gomecoma/src/ent/picture"
)

// PictureCreate is the builder for creating a Picture entity.
type PictureCreate struct {
	config
	mutation *PictureMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (pc *PictureCreate) SetName(s string) *PictureCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pc *PictureCreate) SetNillableName(s *string) *PictureCreate {
	if s != nil {
		pc.SetName(*s)
	}
	return pc
}

// SetFilename sets the "filename" field.
func (pc *PictureCreate) SetFilename(s string) *PictureCreate {
	pc.mutation.SetFilename(s)
	return pc
}

// SetPath sets the "path" field.
func (pc *PictureCreate) SetPath(s string) *PictureCreate {
	pc.mutation.SetPath(s)
	return pc
}

// Mutation returns the PictureMutation object of the builder.
func (pc *PictureCreate) Mutation() *PictureMutation {
	return pc.mutation
}

// Save creates the Picture in the database.
func (pc *PictureCreate) Save(ctx context.Context) (*Picture, error) {
	var (
		err  error
		node *Picture
	)
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PictureMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			if node, err = pc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			if pc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PictureCreate) SaveX(ctx context.Context) *Picture {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PictureCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PictureCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PictureCreate) check() error {
	if _, ok := pc.mutation.Filename(); !ok {
		return &ValidationError{Name: "filename", err: errors.New(`ent: missing required field "Picture.filename"`)}
	}
	if _, ok := pc.mutation.Path(); !ok {
		return &ValidationError{Name: "path", err: errors.New(`ent: missing required field "Picture.path"`)}
	}
	return nil
}

func (pc *PictureCreate) sqlSave(ctx context.Context) (*Picture, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pc *PictureCreate) createSpec() (*Picture, *sqlgraph.CreateSpec) {
	var (
		_node = &Picture{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: picture.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: picture.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: picture.FieldName,
		})
		_node.Name = value
	}
	if value, ok := pc.mutation.Filename(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: picture.FieldFilename,
		})
		_node.Filename = value
	}
	if value, ok := pc.mutation.Path(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: picture.FieldPath,
		})
		_node.Path = value
	}
	return _node, _spec
}

// PictureCreateBulk is the builder for creating many Picture entities in bulk.
type PictureCreateBulk struct {
	config
	builders []*PictureCreate
}

// Save creates the Picture entities in the database.
func (pcb *PictureCreateBulk) Save(ctx context.Context) ([]*Picture, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Picture, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PictureMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PictureCreateBulk) SaveX(ctx context.Context) []*Picture {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PictureCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PictureCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
