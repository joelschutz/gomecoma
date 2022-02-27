// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/joelschutz/gomecoma/src/ent/artist"
	"github.com/joelschutz/gomecoma/src/ent/country"
	"github.com/joelschutz/gomecoma/src/ent/movie"
	"github.com/joelschutz/gomecoma/src/ent/picture"
)

// ArtistCreate is the builder for creating a Artist entity.
type ArtistCreate struct {
	config
	mutation *ArtistMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ac *ArtistCreate) SetName(s string) *ArtistCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetBirthday sets the "birthday" field.
func (ac *ArtistCreate) SetBirthday(t time.Time) *ArtistCreate {
	ac.mutation.SetBirthday(t)
	return ac
}

// SetNillableBirthday sets the "birthday" field if the given value is not nil.
func (ac *ArtistCreate) SetNillableBirthday(t *time.Time) *ArtistCreate {
	if t != nil {
		ac.SetBirthday(*t)
	}
	return ac
}

// SetProfilePictureID sets the "profile_picture" edge to the Picture entity by ID.
func (ac *ArtistCreate) SetProfilePictureID(id int) *ArtistCreate {
	ac.mutation.SetProfilePictureID(id)
	return ac
}

// SetNillableProfilePictureID sets the "profile_picture" edge to the Picture entity by ID if the given value is not nil.
func (ac *ArtistCreate) SetNillableProfilePictureID(id *int) *ArtistCreate {
	if id != nil {
		ac = ac.SetProfilePictureID(*id)
	}
	return ac
}

// SetProfilePicture sets the "profile_picture" edge to the Picture entity.
func (ac *ArtistCreate) SetProfilePicture(p *Picture) *ArtistCreate {
	return ac.SetProfilePictureID(p.ID)
}

// AddPictureIDs adds the "pictures" edge to the Picture entity by IDs.
func (ac *ArtistCreate) AddPictureIDs(ids ...int) *ArtistCreate {
	ac.mutation.AddPictureIDs(ids...)
	return ac
}

// AddPictures adds the "pictures" edges to the Picture entity.
func (ac *ArtistCreate) AddPictures(p ...*Picture) *ArtistCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ac.AddPictureIDs(ids...)
}

// AddDirectedIDs adds the "directed" edge to the Movie entity by IDs.
func (ac *ArtistCreate) AddDirectedIDs(ids ...int) *ArtistCreate {
	ac.mutation.AddDirectedIDs(ids...)
	return ac
}

// AddDirected adds the "directed" edges to the Movie entity.
func (ac *ArtistCreate) AddDirected(m ...*Movie) *ArtistCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return ac.AddDirectedIDs(ids...)
}

// AddActedIDs adds the "acted" edge to the Movie entity by IDs.
func (ac *ArtistCreate) AddActedIDs(ids ...int) *ArtistCreate {
	ac.mutation.AddActedIDs(ids...)
	return ac
}

// AddActed adds the "acted" edges to the Movie entity.
func (ac *ArtistCreate) AddActed(m ...*Movie) *ArtistCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return ac.AddActedIDs(ids...)
}

// AddWroteIDs adds the "wrote" edge to the Movie entity by IDs.
func (ac *ArtistCreate) AddWroteIDs(ids ...int) *ArtistCreate {
	ac.mutation.AddWroteIDs(ids...)
	return ac
}

// AddWrote adds the "wrote" edges to the Movie entity.
func (ac *ArtistCreate) AddWrote(m ...*Movie) *ArtistCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return ac.AddWroteIDs(ids...)
}

// AddCountryIDs adds the "countries" edge to the Country entity by IDs.
func (ac *ArtistCreate) AddCountryIDs(ids ...int) *ArtistCreate {
	ac.mutation.AddCountryIDs(ids...)
	return ac
}

// AddCountries adds the "countries" edges to the Country entity.
func (ac *ArtistCreate) AddCountries(c ...*Country) *ArtistCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ac.AddCountryIDs(ids...)
}

// Mutation returns the ArtistMutation object of the builder.
func (ac *ArtistCreate) Mutation() *ArtistMutation {
	return ac.mutation
}

// Save creates the Artist in the database.
func (ac *ArtistCreate) Save(ctx context.Context) (*Artist, error) {
	var (
		err  error
		node *Artist
	)
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ArtistMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			if node, err = ac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			if ac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *ArtistCreate) SaveX(ctx context.Context) *Artist {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *ArtistCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *ArtistCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *ArtistCreate) check() error {
	if _, ok := ac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Artist.name"`)}
	}
	return nil
}

func (ac *ArtistCreate) sqlSave(ctx context.Context) (*Artist, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ac *ArtistCreate) createSpec() (*Artist, *sqlgraph.CreateSpec) {
	var (
		_node = &Artist{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: artist.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: artist.FieldID,
			},
		}
	)
	if value, ok := ac.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: artist.FieldName,
		})
		_node.Name = value
	}
	if value, ok := ac.mutation.Birthday(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: artist.FieldBirthday,
		})
		_node.Birthday = value
	}
	if nodes := ac.mutation.ProfilePictureIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   artist.ProfilePictureTable,
			Columns: []string{artist.ProfilePictureColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: picture.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.artist_profile_picture = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.PicturesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   artist.PicturesTable,
			Columns: []string{artist.PicturesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: picture.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.DirectedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   artist.DirectedTable,
			Columns: artist.DirectedPrimaryKey,
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
	if nodes := ac.mutation.ActedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   artist.ActedTable,
			Columns: artist.ActedPrimaryKey,
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
	if nodes := ac.mutation.WroteIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   artist.WroteTable,
			Columns: artist.WrotePrimaryKey,
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
	if nodes := ac.mutation.CountriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   artist.CountriesTable,
			Columns: artist.CountriesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: country.FieldID,
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

// ArtistCreateBulk is the builder for creating many Artist entities in bulk.
type ArtistCreateBulk struct {
	config
	builders []*ArtistCreate
}

// Save creates the Artist entities in the database.
func (acb *ArtistCreateBulk) Save(ctx context.Context) ([]*Artist, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Artist, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ArtistMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *ArtistCreateBulk) SaveX(ctx context.Context) []*Artist {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *ArtistCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *ArtistCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
