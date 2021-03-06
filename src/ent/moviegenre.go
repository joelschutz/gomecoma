// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/joelschutz/gomecoma/src/ent/moviegenre"
)

// MovieGenre is the model entity for the MovieGenre schema.
type MovieGenre struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MovieGenreQuery when eager-loading is set.
	Edges MovieGenreEdges `json:"edges"`
}

// MovieGenreEdges holds the relations/edges for other nodes in the graph.
type MovieGenreEdges struct {
	// Movies holds the value of the movies edge.
	Movies []*Movie `json:"movies,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// MoviesOrErr returns the Movies value or an error if the edge
// was not loaded in eager-loading.
func (e MovieGenreEdges) MoviesOrErr() ([]*Movie, error) {
	if e.loadedTypes[0] {
		return e.Movies, nil
	}
	return nil, &NotLoadedError{edge: "movies"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MovieGenre) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case moviegenre.FieldID:
			values[i] = new(sql.NullInt64)
		case moviegenre.FieldName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type MovieGenre", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MovieGenre fields.
func (mg *MovieGenre) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case moviegenre.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			mg.ID = int(value.Int64)
		case moviegenre.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				mg.Name = value.String
			}
		}
	}
	return nil
}

// QueryMovies queries the "movies" edge of the MovieGenre entity.
func (mg *MovieGenre) QueryMovies() *MovieQuery {
	return (&MovieGenreClient{config: mg.config}).QueryMovies(mg)
}

// Update returns a builder for updating this MovieGenre.
// Note that you need to call MovieGenre.Unwrap() before calling this method if this MovieGenre
// was returned from a transaction, and the transaction was committed or rolled back.
func (mg *MovieGenre) Update() *MovieGenreUpdateOne {
	return (&MovieGenreClient{config: mg.config}).UpdateOne(mg)
}

// Unwrap unwraps the MovieGenre entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mg *MovieGenre) Unwrap() *MovieGenre {
	tx, ok := mg.config.driver.(*txDriver)
	if !ok {
		panic("ent: MovieGenre is not a transactional entity")
	}
	mg.config.driver = tx.drv
	return mg
}

// String implements the fmt.Stringer.
func (mg *MovieGenre) String() string {
	var builder strings.Builder
	builder.WriteString("MovieGenre(")
	builder.WriteString(fmt.Sprintf("id=%v", mg.ID))
	builder.WriteString(", name=")
	builder.WriteString(mg.Name)
	builder.WriteByte(')')
	return builder.String()
}

// MovieGenres is a parsable slice of MovieGenre.
type MovieGenres []*MovieGenre

func (mg MovieGenres) config(cfg config) {
	for _i := range mg {
		mg[_i].config = cfg
	}
}
