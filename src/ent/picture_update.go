// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/joelschutz/gomecoma/src/ent/picture"
	"github.com/joelschutz/gomecoma/src/ent/predicate"
)

// PictureUpdate is the builder for updating Picture entities.
type PictureUpdate struct {
	config
	hooks    []Hook
	mutation *PictureMutation
}

// Where appends a list predicates to the PictureUpdate builder.
func (pu *PictureUpdate) Where(ps ...predicate.Picture) *PictureUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetName sets the "name" field.
func (pu *PictureUpdate) SetName(s string) *PictureUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pu *PictureUpdate) SetNillableName(s *string) *PictureUpdate {
	if s != nil {
		pu.SetName(*s)
	}
	return pu
}

// ClearName clears the value of the "name" field.
func (pu *PictureUpdate) ClearName() *PictureUpdate {
	pu.mutation.ClearName()
	return pu
}

// SetFilename sets the "filename" field.
func (pu *PictureUpdate) SetFilename(s string) *PictureUpdate {
	pu.mutation.SetFilename(s)
	return pu
}

// SetPath sets the "path" field.
func (pu *PictureUpdate) SetPath(s string) *PictureUpdate {
	pu.mutation.SetPath(s)
	return pu
}

// Mutation returns the PictureMutation object of the builder.
func (pu *PictureUpdate) Mutation() *PictureMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PictureUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PictureMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			if pu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PictureUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PictureUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PictureUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PictureUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   picture.Table,
			Columns: picture.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: picture.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: picture.FieldName,
		})
	}
	if pu.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: picture.FieldName,
		})
	}
	if value, ok := pu.mutation.Filename(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: picture.FieldFilename,
		})
	}
	if value, ok := pu.mutation.Path(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: picture.FieldPath,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{picture.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// PictureUpdateOne is the builder for updating a single Picture entity.
type PictureUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PictureMutation
}

// SetName sets the "name" field.
func (puo *PictureUpdateOne) SetName(s string) *PictureUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (puo *PictureUpdateOne) SetNillableName(s *string) *PictureUpdateOne {
	if s != nil {
		puo.SetName(*s)
	}
	return puo
}

// ClearName clears the value of the "name" field.
func (puo *PictureUpdateOne) ClearName() *PictureUpdateOne {
	puo.mutation.ClearName()
	return puo
}

// SetFilename sets the "filename" field.
func (puo *PictureUpdateOne) SetFilename(s string) *PictureUpdateOne {
	puo.mutation.SetFilename(s)
	return puo
}

// SetPath sets the "path" field.
func (puo *PictureUpdateOne) SetPath(s string) *PictureUpdateOne {
	puo.mutation.SetPath(s)
	return puo
}

// Mutation returns the PictureMutation object of the builder.
func (puo *PictureUpdateOne) Mutation() *PictureMutation {
	return puo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PictureUpdateOne) Select(field string, fields ...string) *PictureUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Picture entity.
func (puo *PictureUpdateOne) Save(ctx context.Context) (*Picture, error) {
	var (
		err  error
		node *Picture
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PictureMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			if puo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PictureUpdateOne) SaveX(ctx context.Context) *Picture {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PictureUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PictureUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PictureUpdateOne) sqlSave(ctx context.Context) (_node *Picture, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   picture.Table,
			Columns: picture.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: picture.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Picture.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, picture.FieldID)
		for _, f := range fields {
			if !picture.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != picture.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: picture.FieldName,
		})
	}
	if puo.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: picture.FieldName,
		})
	}
	if value, ok := puo.mutation.Filename(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: picture.FieldFilename,
		})
	}
	if value, ok := puo.mutation.Path(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: picture.FieldPath,
		})
	}
	_node = &Picture{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{picture.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
