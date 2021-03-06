// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/joelschutz/gomecoma/src/ent/movie"
	"github.com/joelschutz/gomecoma/src/ent/moviegenre"
	"github.com/joelschutz/gomecoma/src/ent/predicate"
)

// MovieGenreQuery is the builder for querying MovieGenre entities.
type MovieGenreQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.MovieGenre
	// eager-loading edges.
	withMovies *MovieQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MovieGenreQuery builder.
func (mgq *MovieGenreQuery) Where(ps ...predicate.MovieGenre) *MovieGenreQuery {
	mgq.predicates = append(mgq.predicates, ps...)
	return mgq
}

// Limit adds a limit step to the query.
func (mgq *MovieGenreQuery) Limit(limit int) *MovieGenreQuery {
	mgq.limit = &limit
	return mgq
}

// Offset adds an offset step to the query.
func (mgq *MovieGenreQuery) Offset(offset int) *MovieGenreQuery {
	mgq.offset = &offset
	return mgq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mgq *MovieGenreQuery) Unique(unique bool) *MovieGenreQuery {
	mgq.unique = &unique
	return mgq
}

// Order adds an order step to the query.
func (mgq *MovieGenreQuery) Order(o ...OrderFunc) *MovieGenreQuery {
	mgq.order = append(mgq.order, o...)
	return mgq
}

// QueryMovies chains the current query on the "movies" edge.
func (mgq *MovieGenreQuery) QueryMovies() *MovieQuery {
	query := &MovieQuery{config: mgq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mgq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mgq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(moviegenre.Table, moviegenre.FieldID, selector),
			sqlgraph.To(movie.Table, movie.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, moviegenre.MoviesTable, moviegenre.MoviesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(mgq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first MovieGenre entity from the query.
// Returns a *NotFoundError when no MovieGenre was found.
func (mgq *MovieGenreQuery) First(ctx context.Context) (*MovieGenre, error) {
	nodes, err := mgq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{moviegenre.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mgq *MovieGenreQuery) FirstX(ctx context.Context) *MovieGenre {
	node, err := mgq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MovieGenre ID from the query.
// Returns a *NotFoundError when no MovieGenre ID was found.
func (mgq *MovieGenreQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mgq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{moviegenre.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mgq *MovieGenreQuery) FirstIDX(ctx context.Context) int {
	id, err := mgq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MovieGenre entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one MovieGenre entity is not found.
// Returns a *NotFoundError when no MovieGenre entities are found.
func (mgq *MovieGenreQuery) Only(ctx context.Context) (*MovieGenre, error) {
	nodes, err := mgq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{moviegenre.Label}
	default:
		return nil, &NotSingularError{moviegenre.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mgq *MovieGenreQuery) OnlyX(ctx context.Context) *MovieGenre {
	node, err := mgq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MovieGenre ID in the query.
// Returns a *NotSingularError when exactly one MovieGenre ID is not found.
// Returns a *NotFoundError when no entities are found.
func (mgq *MovieGenreQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mgq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{moviegenre.Label}
	default:
		err = &NotSingularError{moviegenre.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mgq *MovieGenreQuery) OnlyIDX(ctx context.Context) int {
	id, err := mgq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MovieGenres.
func (mgq *MovieGenreQuery) All(ctx context.Context) ([]*MovieGenre, error) {
	if err := mgq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return mgq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (mgq *MovieGenreQuery) AllX(ctx context.Context) []*MovieGenre {
	nodes, err := mgq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MovieGenre IDs.
func (mgq *MovieGenreQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := mgq.Select(moviegenre.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mgq *MovieGenreQuery) IDsX(ctx context.Context) []int {
	ids, err := mgq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mgq *MovieGenreQuery) Count(ctx context.Context) (int, error) {
	if err := mgq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return mgq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (mgq *MovieGenreQuery) CountX(ctx context.Context) int {
	count, err := mgq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mgq *MovieGenreQuery) Exist(ctx context.Context) (bool, error) {
	if err := mgq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return mgq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (mgq *MovieGenreQuery) ExistX(ctx context.Context) bool {
	exist, err := mgq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MovieGenreQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mgq *MovieGenreQuery) Clone() *MovieGenreQuery {
	if mgq == nil {
		return nil
	}
	return &MovieGenreQuery{
		config:     mgq.config,
		limit:      mgq.limit,
		offset:     mgq.offset,
		order:      append([]OrderFunc{}, mgq.order...),
		predicates: append([]predicate.MovieGenre{}, mgq.predicates...),
		withMovies: mgq.withMovies.Clone(),
		// clone intermediate query.
		sql:  mgq.sql.Clone(),
		path: mgq.path,
	}
}

// WithMovies tells the query-builder to eager-load the nodes that are connected to
// the "movies" edge. The optional arguments are used to configure the query builder of the edge.
func (mgq *MovieGenreQuery) WithMovies(opts ...func(*MovieQuery)) *MovieGenreQuery {
	query := &MovieQuery{config: mgq.config}
	for _, opt := range opts {
		opt(query)
	}
	mgq.withMovies = query
	return mgq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.MovieGenre.Query().
//		GroupBy(moviegenre.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (mgq *MovieGenreQuery) GroupBy(field string, fields ...string) *MovieGenreGroupBy {
	group := &MovieGenreGroupBy{config: mgq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := mgq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return mgq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.MovieGenre.Query().
//		Select(moviegenre.FieldName).
//		Scan(ctx, &v)
//
func (mgq *MovieGenreQuery) Select(fields ...string) *MovieGenreSelect {
	mgq.fields = append(mgq.fields, fields...)
	return &MovieGenreSelect{MovieGenreQuery: mgq}
}

func (mgq *MovieGenreQuery) prepareQuery(ctx context.Context) error {
	for _, f := range mgq.fields {
		if !moviegenre.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mgq.path != nil {
		prev, err := mgq.path(ctx)
		if err != nil {
			return err
		}
		mgq.sql = prev
	}
	return nil
}

func (mgq *MovieGenreQuery) sqlAll(ctx context.Context) ([]*MovieGenre, error) {
	var (
		nodes       = []*MovieGenre{}
		_spec       = mgq.querySpec()
		loadedTypes = [1]bool{
			mgq.withMovies != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &MovieGenre{config: mgq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, mgq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := mgq.withMovies; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int]*MovieGenre, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.Movies = []*Movie{}
		}
		var (
			edgeids []int
			edges   = make(map[int][]*MovieGenre)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: false,
				Table:   moviegenre.MoviesTable,
				Columns: moviegenre.MoviesPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(moviegenre.MoviesPrimaryKey[0], fks...))
			},
			ScanValues: func() [2]interface{} {
				return [2]interface{}{new(sql.NullInt64), new(sql.NullInt64)}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := int(eout.Int64)
				inValue := int(ein.Int64)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				if _, ok := edges[inValue]; !ok {
					edgeids = append(edgeids, inValue)
				}
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, mgq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "movies": %w`, err)
		}
		query.Where(movie.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "movies" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Movies = append(nodes[i].Edges.Movies, n)
			}
		}
	}

	return nodes, nil
}

func (mgq *MovieGenreQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mgq.querySpec()
	_spec.Node.Columns = mgq.fields
	if len(mgq.fields) > 0 {
		_spec.Unique = mgq.unique != nil && *mgq.unique
	}
	return sqlgraph.CountNodes(ctx, mgq.driver, _spec)
}

func (mgq *MovieGenreQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := mgq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (mgq *MovieGenreQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   moviegenre.Table,
			Columns: moviegenre.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: moviegenre.FieldID,
			},
		},
		From:   mgq.sql,
		Unique: true,
	}
	if unique := mgq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := mgq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, moviegenre.FieldID)
		for i := range fields {
			if fields[i] != moviegenre.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mgq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mgq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mgq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mgq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mgq *MovieGenreQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mgq.driver.Dialect())
	t1 := builder.Table(moviegenre.Table)
	columns := mgq.fields
	if len(columns) == 0 {
		columns = moviegenre.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mgq.sql != nil {
		selector = mgq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mgq.unique != nil && *mgq.unique {
		selector.Distinct()
	}
	for _, p := range mgq.predicates {
		p(selector)
	}
	for _, p := range mgq.order {
		p(selector)
	}
	if offset := mgq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mgq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MovieGenreGroupBy is the group-by builder for MovieGenre entities.
type MovieGenreGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mggb *MovieGenreGroupBy) Aggregate(fns ...AggregateFunc) *MovieGenreGroupBy {
	mggb.fns = append(mggb.fns, fns...)
	return mggb
}

// Scan applies the group-by query and scans the result into the given value.
func (mggb *MovieGenreGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := mggb.path(ctx)
	if err != nil {
		return err
	}
	mggb.sql = query
	return mggb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (mggb *MovieGenreGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := mggb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (mggb *MovieGenreGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(mggb.fields) > 1 {
		return nil, errors.New("ent: MovieGenreGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := mggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (mggb *MovieGenreGroupBy) StringsX(ctx context.Context) []string {
	v, err := mggb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mggb *MovieGenreGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = mggb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{moviegenre.Label}
	default:
		err = fmt.Errorf("ent: MovieGenreGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (mggb *MovieGenreGroupBy) StringX(ctx context.Context) string {
	v, err := mggb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (mggb *MovieGenreGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(mggb.fields) > 1 {
		return nil, errors.New("ent: MovieGenreGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := mggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (mggb *MovieGenreGroupBy) IntsX(ctx context.Context) []int {
	v, err := mggb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mggb *MovieGenreGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = mggb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{moviegenre.Label}
	default:
		err = fmt.Errorf("ent: MovieGenreGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (mggb *MovieGenreGroupBy) IntX(ctx context.Context) int {
	v, err := mggb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (mggb *MovieGenreGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(mggb.fields) > 1 {
		return nil, errors.New("ent: MovieGenreGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := mggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (mggb *MovieGenreGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := mggb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mggb *MovieGenreGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = mggb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{moviegenre.Label}
	default:
		err = fmt.Errorf("ent: MovieGenreGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (mggb *MovieGenreGroupBy) Float64X(ctx context.Context) float64 {
	v, err := mggb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (mggb *MovieGenreGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(mggb.fields) > 1 {
		return nil, errors.New("ent: MovieGenreGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := mggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (mggb *MovieGenreGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := mggb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mggb *MovieGenreGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = mggb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{moviegenre.Label}
	default:
		err = fmt.Errorf("ent: MovieGenreGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (mggb *MovieGenreGroupBy) BoolX(ctx context.Context) bool {
	v, err := mggb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mggb *MovieGenreGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range mggb.fields {
		if !moviegenre.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := mggb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mggb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mggb *MovieGenreGroupBy) sqlQuery() *sql.Selector {
	selector := mggb.sql.Select()
	aggregation := make([]string, 0, len(mggb.fns))
	for _, fn := range mggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(mggb.fields)+len(mggb.fns))
		for _, f := range mggb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(mggb.fields...)...)
}

// MovieGenreSelect is the builder for selecting fields of MovieGenre entities.
type MovieGenreSelect struct {
	*MovieGenreQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (mgs *MovieGenreSelect) Scan(ctx context.Context, v interface{}) error {
	if err := mgs.prepareQuery(ctx); err != nil {
		return err
	}
	mgs.sql = mgs.MovieGenreQuery.sqlQuery(ctx)
	return mgs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (mgs *MovieGenreSelect) ScanX(ctx context.Context, v interface{}) {
	if err := mgs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (mgs *MovieGenreSelect) Strings(ctx context.Context) ([]string, error) {
	if len(mgs.fields) > 1 {
		return nil, errors.New("ent: MovieGenreSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := mgs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (mgs *MovieGenreSelect) StringsX(ctx context.Context) []string {
	v, err := mgs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (mgs *MovieGenreSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = mgs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{moviegenre.Label}
	default:
		err = fmt.Errorf("ent: MovieGenreSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (mgs *MovieGenreSelect) StringX(ctx context.Context) string {
	v, err := mgs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (mgs *MovieGenreSelect) Ints(ctx context.Context) ([]int, error) {
	if len(mgs.fields) > 1 {
		return nil, errors.New("ent: MovieGenreSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := mgs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (mgs *MovieGenreSelect) IntsX(ctx context.Context) []int {
	v, err := mgs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (mgs *MovieGenreSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = mgs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{moviegenre.Label}
	default:
		err = fmt.Errorf("ent: MovieGenreSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (mgs *MovieGenreSelect) IntX(ctx context.Context) int {
	v, err := mgs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (mgs *MovieGenreSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(mgs.fields) > 1 {
		return nil, errors.New("ent: MovieGenreSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := mgs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (mgs *MovieGenreSelect) Float64sX(ctx context.Context) []float64 {
	v, err := mgs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (mgs *MovieGenreSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = mgs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{moviegenre.Label}
	default:
		err = fmt.Errorf("ent: MovieGenreSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (mgs *MovieGenreSelect) Float64X(ctx context.Context) float64 {
	v, err := mgs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (mgs *MovieGenreSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(mgs.fields) > 1 {
		return nil, errors.New("ent: MovieGenreSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := mgs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (mgs *MovieGenreSelect) BoolsX(ctx context.Context) []bool {
	v, err := mgs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (mgs *MovieGenreSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = mgs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{moviegenre.Label}
	default:
		err = fmt.Errorf("ent: MovieGenreSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (mgs *MovieGenreSelect) BoolX(ctx context.Context) bool {
	v, err := mgs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mgs *MovieGenreSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := mgs.sql.Query()
	if err := mgs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
