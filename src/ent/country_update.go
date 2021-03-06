// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/joelschutz/gomecoma/src/ent/artist"
	"github.com/joelschutz/gomecoma/src/ent/country"
	"github.com/joelschutz/gomecoma/src/ent/movie"
	"github.com/joelschutz/gomecoma/src/ent/predicate"
)

// CountryUpdate is the builder for updating Country entities.
type CountryUpdate struct {
	config
	hooks    []Hook
	mutation *CountryMutation
}

// Where appends a list predicates to the CountryUpdate builder.
func (cu *CountryUpdate) Where(ps ...predicate.Country) *CountryUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetName sets the "name" field.
func (cu *CountryUpdate) SetName(s string) *CountryUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetCode sets the "code" field.
func (cu *CountryUpdate) SetCode(s string) *CountryUpdate {
	cu.mutation.SetCode(s)
	return cu
}

// AddMovieIDs adds the "movies" edge to the Movie entity by IDs.
func (cu *CountryUpdate) AddMovieIDs(ids ...int) *CountryUpdate {
	cu.mutation.AddMovieIDs(ids...)
	return cu
}

// AddMovies adds the "movies" edges to the Movie entity.
func (cu *CountryUpdate) AddMovies(m ...*Movie) *CountryUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return cu.AddMovieIDs(ids...)
}

// AddArtistIDs adds the "artists" edge to the Artist entity by IDs.
func (cu *CountryUpdate) AddArtistIDs(ids ...int) *CountryUpdate {
	cu.mutation.AddArtistIDs(ids...)
	return cu
}

// AddArtists adds the "artists" edges to the Artist entity.
func (cu *CountryUpdate) AddArtists(a ...*Artist) *CountryUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cu.AddArtistIDs(ids...)
}

// Mutation returns the CountryMutation object of the builder.
func (cu *CountryUpdate) Mutation() *CountryMutation {
	return cu.mutation
}

// ClearMovies clears all "movies" edges to the Movie entity.
func (cu *CountryUpdate) ClearMovies() *CountryUpdate {
	cu.mutation.ClearMovies()
	return cu
}

// RemoveMovieIDs removes the "movies" edge to Movie entities by IDs.
func (cu *CountryUpdate) RemoveMovieIDs(ids ...int) *CountryUpdate {
	cu.mutation.RemoveMovieIDs(ids...)
	return cu
}

// RemoveMovies removes "movies" edges to Movie entities.
func (cu *CountryUpdate) RemoveMovies(m ...*Movie) *CountryUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return cu.RemoveMovieIDs(ids...)
}

// ClearArtists clears all "artists" edges to the Artist entity.
func (cu *CountryUpdate) ClearArtists() *CountryUpdate {
	cu.mutation.ClearArtists()
	return cu
}

// RemoveArtistIDs removes the "artists" edge to Artist entities by IDs.
func (cu *CountryUpdate) RemoveArtistIDs(ids ...int) *CountryUpdate {
	cu.mutation.RemoveArtistIDs(ids...)
	return cu
}

// RemoveArtists removes "artists" edges to Artist entities.
func (cu *CountryUpdate) RemoveArtists(a ...*Artist) *CountryUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cu.RemoveArtistIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CountryUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CountryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CountryUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CountryUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CountryUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CountryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   country.Table,
			Columns: country.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: country.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: country.FieldName,
		})
	}
	if value, ok := cu.mutation.Code(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: country.FieldCode,
		})
	}
	if cu.mutation.MoviesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   country.MoviesTable,
			Columns: country.MoviesPrimaryKey,
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
	if nodes := cu.mutation.RemovedMoviesIDs(); len(nodes) > 0 && !cu.mutation.MoviesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   country.MoviesTable,
			Columns: country.MoviesPrimaryKey,
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
	if nodes := cu.mutation.MoviesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   country.MoviesTable,
			Columns: country.MoviesPrimaryKey,
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
	if cu.mutation.ArtistsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   country.ArtistsTable,
			Columns: country.ArtistsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: artist.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedArtistsIDs(); len(nodes) > 0 && !cu.mutation.ArtistsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   country.ArtistsTable,
			Columns: country.ArtistsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: artist.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ArtistsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   country.ArtistsTable,
			Columns: country.ArtistsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: artist.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{country.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// CountryUpdateOne is the builder for updating a single Country entity.
type CountryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CountryMutation
}

// SetName sets the "name" field.
func (cuo *CountryUpdateOne) SetName(s string) *CountryUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetCode sets the "code" field.
func (cuo *CountryUpdateOne) SetCode(s string) *CountryUpdateOne {
	cuo.mutation.SetCode(s)
	return cuo
}

// AddMovieIDs adds the "movies" edge to the Movie entity by IDs.
func (cuo *CountryUpdateOne) AddMovieIDs(ids ...int) *CountryUpdateOne {
	cuo.mutation.AddMovieIDs(ids...)
	return cuo
}

// AddMovies adds the "movies" edges to the Movie entity.
func (cuo *CountryUpdateOne) AddMovies(m ...*Movie) *CountryUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return cuo.AddMovieIDs(ids...)
}

// AddArtistIDs adds the "artists" edge to the Artist entity by IDs.
func (cuo *CountryUpdateOne) AddArtistIDs(ids ...int) *CountryUpdateOne {
	cuo.mutation.AddArtistIDs(ids...)
	return cuo
}

// AddArtists adds the "artists" edges to the Artist entity.
func (cuo *CountryUpdateOne) AddArtists(a ...*Artist) *CountryUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cuo.AddArtistIDs(ids...)
}

// Mutation returns the CountryMutation object of the builder.
func (cuo *CountryUpdateOne) Mutation() *CountryMutation {
	return cuo.mutation
}

// ClearMovies clears all "movies" edges to the Movie entity.
func (cuo *CountryUpdateOne) ClearMovies() *CountryUpdateOne {
	cuo.mutation.ClearMovies()
	return cuo
}

// RemoveMovieIDs removes the "movies" edge to Movie entities by IDs.
func (cuo *CountryUpdateOne) RemoveMovieIDs(ids ...int) *CountryUpdateOne {
	cuo.mutation.RemoveMovieIDs(ids...)
	return cuo
}

// RemoveMovies removes "movies" edges to Movie entities.
func (cuo *CountryUpdateOne) RemoveMovies(m ...*Movie) *CountryUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return cuo.RemoveMovieIDs(ids...)
}

// ClearArtists clears all "artists" edges to the Artist entity.
func (cuo *CountryUpdateOne) ClearArtists() *CountryUpdateOne {
	cuo.mutation.ClearArtists()
	return cuo
}

// RemoveArtistIDs removes the "artists" edge to Artist entities by IDs.
func (cuo *CountryUpdateOne) RemoveArtistIDs(ids ...int) *CountryUpdateOne {
	cuo.mutation.RemoveArtistIDs(ids...)
	return cuo
}

// RemoveArtists removes "artists" edges to Artist entities.
func (cuo *CountryUpdateOne) RemoveArtists(a ...*Artist) *CountryUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cuo.RemoveArtistIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CountryUpdateOne) Select(field string, fields ...string) *CountryUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Country entity.
func (cuo *CountryUpdateOne) Save(ctx context.Context) (*Country, error) {
	var (
		err  error
		node *Country
	)
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CountryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CountryUpdateOne) SaveX(ctx context.Context) *Country {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CountryUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CountryUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CountryUpdateOne) sqlSave(ctx context.Context) (_node *Country, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   country.Table,
			Columns: country.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: country.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Country.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, country.FieldID)
		for _, f := range fields {
			if !country.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != country.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: country.FieldName,
		})
	}
	if value, ok := cuo.mutation.Code(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: country.FieldCode,
		})
	}
	if cuo.mutation.MoviesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   country.MoviesTable,
			Columns: country.MoviesPrimaryKey,
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
	if nodes := cuo.mutation.RemovedMoviesIDs(); len(nodes) > 0 && !cuo.mutation.MoviesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   country.MoviesTable,
			Columns: country.MoviesPrimaryKey,
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
	if nodes := cuo.mutation.MoviesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   country.MoviesTable,
			Columns: country.MoviesPrimaryKey,
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
	if cuo.mutation.ArtistsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   country.ArtistsTable,
			Columns: country.ArtistsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: artist.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedArtistsIDs(); len(nodes) > 0 && !cuo.mutation.ArtistsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   country.ArtistsTable,
			Columns: country.ArtistsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: artist.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ArtistsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   country.ArtistsTable,
			Columns: country.ArtistsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: artist.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Country{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{country.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
