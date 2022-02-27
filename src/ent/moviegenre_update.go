// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/joelschutz/gomecoma/src/ent/movie"
	"github.com/joelschutz/gomecoma/src/ent/moviegenre"
	"github.com/joelschutz/gomecoma/src/ent/predicate"
)

// MovieGenreUpdate is the builder for updating MovieGenre entities.
type MovieGenreUpdate struct {
	config
	hooks    []Hook
	mutation *MovieGenreMutation
}

// Where appends a list predicates to the MovieGenreUpdate builder.
func (mgu *MovieGenreUpdate) Where(ps ...predicate.MovieGenre) *MovieGenreUpdate {
	mgu.mutation.Where(ps...)
	return mgu
}

// SetName sets the "name" field.
func (mgu *MovieGenreUpdate) SetName(s string) *MovieGenreUpdate {
	mgu.mutation.SetName(s)
	return mgu
}

// AddMovieIDs adds the "movies" edge to the Movie entity by IDs.
func (mgu *MovieGenreUpdate) AddMovieIDs(ids ...int) *MovieGenreUpdate {
	mgu.mutation.AddMovieIDs(ids...)
	return mgu
}

// AddMovies adds the "movies" edges to the Movie entity.
func (mgu *MovieGenreUpdate) AddMovies(m ...*Movie) *MovieGenreUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mgu.AddMovieIDs(ids...)
}

// Mutation returns the MovieGenreMutation object of the builder.
func (mgu *MovieGenreUpdate) Mutation() *MovieGenreMutation {
	return mgu.mutation
}

// ClearMovies clears all "movies" edges to the Movie entity.
func (mgu *MovieGenreUpdate) ClearMovies() *MovieGenreUpdate {
	mgu.mutation.ClearMovies()
	return mgu
}

// RemoveMovieIDs removes the "movies" edge to Movie entities by IDs.
func (mgu *MovieGenreUpdate) RemoveMovieIDs(ids ...int) *MovieGenreUpdate {
	mgu.mutation.RemoveMovieIDs(ids...)
	return mgu
}

// RemoveMovies removes "movies" edges to Movie entities.
func (mgu *MovieGenreUpdate) RemoveMovies(m ...*Movie) *MovieGenreUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mgu.RemoveMovieIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mgu *MovieGenreUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mgu.hooks) == 0 {
		affected, err = mgu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MovieGenreMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mgu.mutation = mutation
			affected, err = mgu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mgu.hooks) - 1; i >= 0; i-- {
			if mgu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mgu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mgu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mgu *MovieGenreUpdate) SaveX(ctx context.Context) int {
	affected, err := mgu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mgu *MovieGenreUpdate) Exec(ctx context.Context) error {
	_, err := mgu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mgu *MovieGenreUpdate) ExecX(ctx context.Context) {
	if err := mgu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mgu *MovieGenreUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   moviegenre.Table,
			Columns: moviegenre.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: moviegenre.FieldID,
			},
		},
	}
	if ps := mgu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mgu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: moviegenre.FieldName,
		})
	}
	if mgu.mutation.MoviesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mgu.mutation.RemovedMoviesIDs(); len(nodes) > 0 && !mgu.mutation.MoviesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mgu.mutation.MoviesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mgu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{moviegenre.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// MovieGenreUpdateOne is the builder for updating a single MovieGenre entity.
type MovieGenreUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MovieGenreMutation
}

// SetName sets the "name" field.
func (mguo *MovieGenreUpdateOne) SetName(s string) *MovieGenreUpdateOne {
	mguo.mutation.SetName(s)
	return mguo
}

// AddMovieIDs adds the "movies" edge to the Movie entity by IDs.
func (mguo *MovieGenreUpdateOne) AddMovieIDs(ids ...int) *MovieGenreUpdateOne {
	mguo.mutation.AddMovieIDs(ids...)
	return mguo
}

// AddMovies adds the "movies" edges to the Movie entity.
func (mguo *MovieGenreUpdateOne) AddMovies(m ...*Movie) *MovieGenreUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mguo.AddMovieIDs(ids...)
}

// Mutation returns the MovieGenreMutation object of the builder.
func (mguo *MovieGenreUpdateOne) Mutation() *MovieGenreMutation {
	return mguo.mutation
}

// ClearMovies clears all "movies" edges to the Movie entity.
func (mguo *MovieGenreUpdateOne) ClearMovies() *MovieGenreUpdateOne {
	mguo.mutation.ClearMovies()
	return mguo
}

// RemoveMovieIDs removes the "movies" edge to Movie entities by IDs.
func (mguo *MovieGenreUpdateOne) RemoveMovieIDs(ids ...int) *MovieGenreUpdateOne {
	mguo.mutation.RemoveMovieIDs(ids...)
	return mguo
}

// RemoveMovies removes "movies" edges to Movie entities.
func (mguo *MovieGenreUpdateOne) RemoveMovies(m ...*Movie) *MovieGenreUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mguo.RemoveMovieIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mguo *MovieGenreUpdateOne) Select(field string, fields ...string) *MovieGenreUpdateOne {
	mguo.fields = append([]string{field}, fields...)
	return mguo
}

// Save executes the query and returns the updated MovieGenre entity.
func (mguo *MovieGenreUpdateOne) Save(ctx context.Context) (*MovieGenre, error) {
	var (
		err  error
		node *MovieGenre
	)
	if len(mguo.hooks) == 0 {
		node, err = mguo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MovieGenreMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mguo.mutation = mutation
			node, err = mguo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(mguo.hooks) - 1; i >= 0; i-- {
			if mguo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mguo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mguo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (mguo *MovieGenreUpdateOne) SaveX(ctx context.Context) *MovieGenre {
	node, err := mguo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mguo *MovieGenreUpdateOne) Exec(ctx context.Context) error {
	_, err := mguo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mguo *MovieGenreUpdateOne) ExecX(ctx context.Context) {
	if err := mguo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mguo *MovieGenreUpdateOne) sqlSave(ctx context.Context) (_node *MovieGenre, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   moviegenre.Table,
			Columns: moviegenre.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: moviegenre.FieldID,
			},
		},
	}
	id, ok := mguo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MovieGenre.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mguo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, moviegenre.FieldID)
		for _, f := range fields {
			if !moviegenre.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != moviegenre.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mguo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mguo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: moviegenre.FieldName,
		})
	}
	if mguo.mutation.MoviesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mguo.mutation.RemovedMoviesIDs(); len(nodes) > 0 && !mguo.mutation.MoviesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mguo.mutation.MoviesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &MovieGenre{config: mguo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mguo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{moviegenre.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}