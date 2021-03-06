// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/joelschutz/gomecoma/src/ent/rating"
)

// RatingCreate is the builder for creating a Rating entity.
type RatingCreate struct {
	config
	mutation *RatingMutation
	hooks    []Hook
}

// SetOrigin sets the "origin" field.
func (rc *RatingCreate) SetOrigin(s string) *RatingCreate {
	rc.mutation.SetOrigin(s)
	return rc
}

// SetOriginalRating sets the "original_rating" field.
func (rc *RatingCreate) SetOriginalRating(s string) *RatingCreate {
	rc.mutation.SetOriginalRating(s)
	return rc
}

// SetNormalizedRating sets the "normalized_rating" field.
func (rc *RatingCreate) SetNormalizedRating(i int) *RatingCreate {
	rc.mutation.SetNormalizedRating(i)
	return rc
}

// Mutation returns the RatingMutation object of the builder.
func (rc *RatingCreate) Mutation() *RatingMutation {
	return rc.mutation
}

// Save creates the Rating in the database.
func (rc *RatingCreate) Save(ctx context.Context) (*Rating, error) {
	var (
		err  error
		node *Rating
	)
	if len(rc.hooks) == 0 {
		if err = rc.check(); err != nil {
			return nil, err
		}
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RatingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rc.check(); err != nil {
				return nil, err
			}
			rc.mutation = mutation
			if node, err = rc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			if rc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RatingCreate) SaveX(ctx context.Context) *Rating {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RatingCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RatingCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RatingCreate) check() error {
	if _, ok := rc.mutation.Origin(); !ok {
		return &ValidationError{Name: "origin", err: errors.New(`ent: missing required field "Rating.origin"`)}
	}
	if _, ok := rc.mutation.OriginalRating(); !ok {
		return &ValidationError{Name: "original_rating", err: errors.New(`ent: missing required field "Rating.original_rating"`)}
	}
	if _, ok := rc.mutation.NormalizedRating(); !ok {
		return &ValidationError{Name: "normalized_rating", err: errors.New(`ent: missing required field "Rating.normalized_rating"`)}
	}
	if v, ok := rc.mutation.NormalizedRating(); ok {
		if err := rating.NormalizedRatingValidator(v); err != nil {
			return &ValidationError{Name: "normalized_rating", err: fmt.Errorf(`ent: validator failed for field "Rating.normalized_rating": %w`, err)}
		}
	}
	return nil
}

func (rc *RatingCreate) sqlSave(ctx context.Context) (*Rating, error) {
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (rc *RatingCreate) createSpec() (*Rating, *sqlgraph.CreateSpec) {
	var (
		_node = &Rating{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: rating.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: rating.FieldID,
			},
		}
	)
	if value, ok := rc.mutation.Origin(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rating.FieldOrigin,
		})
		_node.Origin = value
	}
	if value, ok := rc.mutation.OriginalRating(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rating.FieldOriginalRating,
		})
		_node.OriginalRating = value
	}
	if value, ok := rc.mutation.NormalizedRating(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: rating.FieldNormalizedRating,
		})
		_node.NormalizedRating = value
	}
	return _node, _spec
}

// RatingCreateBulk is the builder for creating many Rating entities in bulk.
type RatingCreateBulk struct {
	config
	builders []*RatingCreate
}

// Save creates the Rating entities in the database.
func (rcb *RatingCreateBulk) Save(ctx context.Context) ([]*Rating, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Rating, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RatingMutation)
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
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RatingCreateBulk) SaveX(ctx context.Context) []*Rating {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RatingCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RatingCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
