// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/joelschutz/gomecoma/src/ent/movie"
	"github.com/joelschutz/gomecoma/src/ent/moviegenre"
)

// MovieGenreCreate is the builder for creating a MovieGenre entity.
type MovieGenreCreate struct {
	config
	mutation *MovieGenreMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (mgc *MovieGenreCreate) SetName(s string) *MovieGenreCreate {
	mgc.mutation.SetName(s)
	return mgc
}

// AddMovieIDs adds the "movies" edge to the Movie entity by IDs.
func (mgc *MovieGenreCreate) AddMovieIDs(ids ...int) *MovieGenreCreate {
	mgc.mutation.AddMovieIDs(ids...)
	return mgc
}

// AddMovies adds the "movies" edges to the Movie entity.
func (mgc *MovieGenreCreate) AddMovies(m ...*Movie) *MovieGenreCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mgc.AddMovieIDs(ids...)
}

// Mutation returns the MovieGenreMutation object of the builder.
func (mgc *MovieGenreCreate) Mutation() *MovieGenreMutation {
	return mgc.mutation
}

// Save creates the MovieGenre in the database.
func (mgc *MovieGenreCreate) Save(ctx context.Context) (*MovieGenre, error) {
	var (
		err  error
		node *MovieGenre
	)
	if len(mgc.hooks) == 0 {
		if err = mgc.check(); err != nil {
			return nil, err
		}
		node, err = mgc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MovieGenreMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mgc.check(); err != nil {
				return nil, err
			}
			mgc.mutation = mutation
			if node, err = mgc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mgc.hooks) - 1; i >= 0; i-- {
			if mgc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mgc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mgc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mgc *MovieGenreCreate) SaveX(ctx context.Context) *MovieGenre {
	v, err := mgc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mgc *MovieGenreCreate) Exec(ctx context.Context) error {
	_, err := mgc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mgc *MovieGenreCreate) ExecX(ctx context.Context) {
	if err := mgc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mgc *MovieGenreCreate) check() error {
	if _, ok := mgc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "MovieGenre.name"`)}
	}
	return nil
}

func (mgc *MovieGenreCreate) sqlSave(ctx context.Context) (*MovieGenre, error) {
	_node, _spec := mgc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mgc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (mgc *MovieGenreCreate) createSpec() (*MovieGenre, *sqlgraph.CreateSpec) {
	var (
		_node = &MovieGenre{config: mgc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: moviegenre.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: moviegenre.FieldID,
			},
		}
	)
	if value, ok := mgc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: moviegenre.FieldName,
		})
		_node.Name = value
	}
	if nodes := mgc.mutation.MoviesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   moviegenre.MoviesTable,
			Columns: moviegenre.MoviesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: movie.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MovieGenreCreateBulk is the builder for creating many MovieGenre entities in bulk.
type MovieGenreCreateBulk struct {
	config
	builders []*MovieGenreCreate
}

// Save creates the MovieGenre entities in the database.
func (mgcb *MovieGenreCreateBulk) Save(ctx context.Context) ([]*MovieGenre, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mgcb.builders))
	nodes := make([]*MovieGenre, len(mgcb.builders))
	mutators := make([]Mutator, len(mgcb.builders))
	for i := range mgcb.builders {
		func(i int, root context.Context) {
			builder := mgcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MovieGenreMutation)
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
					_, err = mutators[i+1].Mutate(root, mgcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mgcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mgcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mgcb *MovieGenreCreateBulk) SaveX(ctx context.Context) []*MovieGenre {
	v, err := mgcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mgcb *MovieGenreCreateBulk) Exec(ctx context.Context) error {
	_, err := mgcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mgcb *MovieGenreCreateBulk) ExecX(ctx context.Context) {
	if err := mgcb.Exec(ctx); err != nil {
		panic(err)
	}
}