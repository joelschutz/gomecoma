// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/joelschutz/gomecoma/src/ent/predicate"
	"github.com/joelschutz/gomecoma/src/ent/rating"
)

// RatingUpdate is the builder for updating Rating entities.
type RatingUpdate struct {
	config
	hooks    []Hook
	mutation *RatingMutation
}

// Where appends a list predicates to the RatingUpdate builder.
func (ru *RatingUpdate) Where(ps ...predicate.Rating) *RatingUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetOrigin sets the "origin" field.
func (ru *RatingUpdate) SetOrigin(s string) *RatingUpdate {
	ru.mutation.SetOrigin(s)
	return ru
}

// SetOriginalRating sets the "original_rating" field.
func (ru *RatingUpdate) SetOriginalRating(s string) *RatingUpdate {
	ru.mutation.SetOriginalRating(s)
	return ru
}

// SetNormalizedRating sets the "normalized_rating" field.
func (ru *RatingUpdate) SetNormalizedRating(i int) *RatingUpdate {
	ru.mutation.ResetNormalizedRating()
	ru.mutation.SetNormalizedRating(i)
	return ru
}

// AddNormalizedRating adds i to the "normalized_rating" field.
func (ru *RatingUpdate) AddNormalizedRating(i int) *RatingUpdate {
	ru.mutation.AddNormalizedRating(i)
	return ru
}

// Mutation returns the RatingMutation object of the builder.
func (ru *RatingUpdate) Mutation() *RatingMutation {
	return ru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RatingUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ru.hooks) == 0 {
		if err = ru.check(); err != nil {
			return 0, err
		}
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RatingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ru.check(); err != nil {
				return 0, err
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			if ru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RatingUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RatingUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RatingUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *RatingUpdate) check() error {
	if v, ok := ru.mutation.NormalizedRating(); ok {
		if err := rating.NormalizedRatingValidator(v); err != nil {
			return &ValidationError{Name: "normalized_rating", err: fmt.Errorf(`ent: validator failed for field "Rating.normalized_rating": %w`, err)}
		}
	}
	return nil
}

func (ru *RatingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   rating.Table,
			Columns: rating.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: rating.FieldID,
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Origin(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rating.FieldOrigin,
		})
	}
	if value, ok := ru.mutation.OriginalRating(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rating.FieldOriginalRating,
		})
	}
	if value, ok := ru.mutation.NormalizedRating(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: rating.FieldNormalizedRating,
		})
	}
	if value, ok := ru.mutation.AddedNormalizedRating(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: rating.FieldNormalizedRating,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{rating.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// RatingUpdateOne is the builder for updating a single Rating entity.
type RatingUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RatingMutation
}

// SetOrigin sets the "origin" field.
func (ruo *RatingUpdateOne) SetOrigin(s string) *RatingUpdateOne {
	ruo.mutation.SetOrigin(s)
	return ruo
}

// SetOriginalRating sets the "original_rating" field.
func (ruo *RatingUpdateOne) SetOriginalRating(s string) *RatingUpdateOne {
	ruo.mutation.SetOriginalRating(s)
	return ruo
}

// SetNormalizedRating sets the "normalized_rating" field.
func (ruo *RatingUpdateOne) SetNormalizedRating(i int) *RatingUpdateOne {
	ruo.mutation.ResetNormalizedRating()
	ruo.mutation.SetNormalizedRating(i)
	return ruo
}

// AddNormalizedRating adds i to the "normalized_rating" field.
func (ruo *RatingUpdateOne) AddNormalizedRating(i int) *RatingUpdateOne {
	ruo.mutation.AddNormalizedRating(i)
	return ruo
}

// Mutation returns the RatingMutation object of the builder.
func (ruo *RatingUpdateOne) Mutation() *RatingMutation {
	return ruo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RatingUpdateOne) Select(field string, fields ...string) *RatingUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Rating entity.
func (ruo *RatingUpdateOne) Save(ctx context.Context) (*Rating, error) {
	var (
		err  error
		node *Rating
	)
	if len(ruo.hooks) == 0 {
		if err = ruo.check(); err != nil {
			return nil, err
		}
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RatingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ruo.check(); err != nil {
				return nil, err
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			if ruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RatingUpdateOne) SaveX(ctx context.Context) *Rating {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RatingUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RatingUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *RatingUpdateOne) check() error {
	if v, ok := ruo.mutation.NormalizedRating(); ok {
		if err := rating.NormalizedRatingValidator(v); err != nil {
			return &ValidationError{Name: "normalized_rating", err: fmt.Errorf(`ent: validator failed for field "Rating.normalized_rating": %w`, err)}
		}
	}
	return nil
}

func (ruo *RatingUpdateOne) sqlSave(ctx context.Context) (_node *Rating, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   rating.Table,
			Columns: rating.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: rating.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Rating.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, rating.FieldID)
		for _, f := range fields {
			if !rating.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != rating.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.Origin(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rating.FieldOrigin,
		})
	}
	if value, ok := ruo.mutation.OriginalRating(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rating.FieldOriginalRating,
		})
	}
	if value, ok := ruo.mutation.NormalizedRating(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: rating.FieldNormalizedRating,
		})
	}
	if value, ok := ruo.mutation.AddedNormalizedRating(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: rating.FieldNormalizedRating,
		})
	}
	_node = &Rating{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{rating.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}