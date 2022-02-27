// Code generated by entc, DO NOT EDIT.

package moviegenre

const (
	// Label holds the string label denoting the moviegenre type in the database.
	Label = "movie_genre"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeMovies holds the string denoting the movies edge name in mutations.
	EdgeMovies = "movies"
	// Table holds the table name of the moviegenre in the database.
	Table = "movie_genres"
	// MoviesTable is the table that holds the movies relation/edge. The primary key declared below.
	MoviesTable = "movie_genre_movies"
	// MoviesInverseTable is the table name for the Movie entity.
	// It exists in this package in order to avoid circular dependency with the "movie" package.
	MoviesInverseTable = "movies"
)

// Columns holds all SQL columns for moviegenre fields.
var Columns = []string{
	FieldID,
	FieldName,
}

var (
	// MoviesPrimaryKey and MoviesColumn2 are the table columns denoting the
	// primary key for the movies relation (M2M).
	MoviesPrimaryKey = []string{"movie_genre_id", "movie_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}