// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/joelschutz/gomecoma/src/ent/artist"
	"github.com/joelschutz/gomecoma/src/ent/country"
	"github.com/joelschutz/gomecoma/src/ent/movie"
)

// CountryCreate is the builder for creating a Country entity.
type CountryCreate struct {
	config
	mutation *CountryMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (cc *CountryCreate) SetName(s string) *CountryCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetCode sets the "code" field.
func (cc *CountryCreate) SetCode(s string) *CountryCreate {
	cc.mutation.SetCode(s)
	return cc
}

// AddMovieIDs adds the "movies" edge to the Movie entity by IDs.
func (cc *CountryCreate) AddMovieIDs(ids ...int) *CountryCreate {
	cc.mutation.AddMovieIDs(ids...)
	return cc
}

// AddMovies adds the "movies" edges to the Movie entity.
func (cc *CountryCreate) AddMovies(m ...*Movie) *CountryCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return cc.AddMovieIDs(ids...)
}

// AddArtistIDs adds the "artists" edge to the Artist entity by IDs.
func (cc *CountryCreate) AddArtistIDs(ids ...int) *CountryCreate {
	cc.mutation.AddArtistIDs(ids...)
	return cc
}

// AddArtists adds the "artists" edges to the Artist entity.
func (cc *CountryCreate) AddArtists(a ...*Artist) *CountryCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cc.AddArtistIDs(ids...)
}

// Mutation returns the CountryMutation object of the builder.
func (cc *CountryCreate) Mutation() *CountryMutation {
	return cc.mutation
}

// Save creates the Country in the database.
func (cc *CountryCreate) Save(ctx context.Context) (*Country, error) {
	var (
		err  error
		node *Country
	)
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CountryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CountryCreate) SaveX(ctx context.Context) *Country {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CountryCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CountryCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CountryCreate) check() error {
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Country.name"`)}
	}
	if _, ok := cc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "Country.code"`)}
	}
	return nil
}

func (cc *CountryCreate) sqlSave(ctx context.Context) (*Country, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (cc *CountryCreate) createSpec() (*Country, *sqlgraph.CreateSpec) {
	var (
		_node = &Country{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: country.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: country.FieldID,
			},
		}
	)
	if value, ok := cc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: country.FieldName,
		})
		_node.Name = value
	}
	if value, ok := cc.mutation.Code(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: country.FieldCode,
		})
		_node.Code = value
	}
	if nodes := cc.mutation.MoviesIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.ArtistsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CountryCreateBulk is the builder for creating many Country entities in bulk.
type CountryCreateBulk struct {
	config
	builders []*CountryCreate
}

// Save creates the Country entities in the database.
func (ccb *CountryCreateBulk) Save(ctx context.Context) ([]*Country, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Country, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CountryMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CountryCreateBulk) SaveX(ctx context.Context) []*Country {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CountryCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CountryCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}