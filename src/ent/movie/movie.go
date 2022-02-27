// Code generated by entc, DO NOT EDIT.

package movie

const (
	// Label holds the string label denoting the movie type in the database.
	Label = "movie"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldOriginalTitle holds the string denoting the original_title field in the database.
	FieldOriginalTitle = "original_title"
	// FieldLanguages holds the string denoting the languages field in the database.
	FieldLanguages = "languages"
	// FieldReleaseDate holds the string denoting the release_date field in the database.
	FieldReleaseDate = "release_date"
	// FieldPlot holds the string denoting the plot field in the database.
	FieldPlot = "plot"
	// FieldDuration holds the string denoting the duration field in the database.
	FieldDuration = "duration"
	// FieldWatched holds the string denoting the watched field in the database.
	FieldWatched = "watched"
	// EdgeRatings holds the string denoting the ratings edge name in mutations.
	EdgeRatings = "ratings"
	// EdgePoster holds the string denoting the poster edge name in mutations.
	EdgePoster = "poster"
	// EdgeFanart holds the string denoting the fanart edge name in mutations.
	EdgeFanart = "fanart"
	// EdgeCast holds the string denoting the cast edge name in mutations.
	EdgeCast = "cast"
	// EdgeDirectors holds the string denoting the directors edge name in mutations.
	EdgeDirectors = "directors"
	// EdgeWriters holds the string denoting the writers edge name in mutations.
	EdgeWriters = "writers"
	// EdgeGenres holds the string denoting the genres edge name in mutations.
	EdgeGenres = "genres"
	// EdgeCountries holds the string denoting the countries edge name in mutations.
	EdgeCountries = "countries"
	// Table holds the table name of the movie in the database.
	Table = "movies"
	// RatingsTable is the table that holds the ratings relation/edge.
	RatingsTable = "ratings"
	// RatingsInverseTable is the table name for the Rating entity.
	// It exists in this package in order to avoid circular dependency with the "rating" package.
	RatingsInverseTable = "ratings"
	// RatingsColumn is the table column denoting the ratings relation/edge.
	RatingsColumn = "movie_ratings"
	// PosterTable is the table that holds the poster relation/edge.
	PosterTable = "movies"
	// PosterInverseTable is the table name for the Picture entity.
	// It exists in this package in order to avoid circular dependency with the "picture" package.
	PosterInverseTable = "pictures"
	// PosterColumn is the table column denoting the poster relation/edge.
	PosterColumn = "movie_poster"
	// FanartTable is the table that holds the fanart relation/edge.
	FanartTable = "pictures"
	// FanartInverseTable is the table name for the Picture entity.
	// It exists in this package in order to avoid circular dependency with the "picture" package.
	FanartInverseTable = "pictures"
	// FanartColumn is the table column denoting the fanart relation/edge.
	FanartColumn = "movie_fanart"
	// CastTable is the table that holds the cast relation/edge. The primary key declared below.
	CastTable = "movie_cast"
	// CastInverseTable is the table name for the Artist entity.
	// It exists in this package in order to avoid circular dependency with the "artist" package.
	CastInverseTable = "artists"
	// DirectorsTable is the table that holds the directors relation/edge. The primary key declared below.
	DirectorsTable = "movie_directors"
	// DirectorsInverseTable is the table name for the Artist entity.
	// It exists in this package in order to avoid circular dependency with the "artist" package.
	DirectorsInverseTable = "artists"
	// WritersTable is the table that holds the writers relation/edge. The primary key declared below.
	WritersTable = "movie_writers"
	// WritersInverseTable is the table name for the Artist entity.
	// It exists in this package in order to avoid circular dependency with the "artist" package.
	WritersInverseTable = "artists"
	// GenresTable is the table that holds the genres relation/edge. The primary key declared below.
	GenresTable = "movie_genre_movies"
	// GenresInverseTable is the table name for the MovieGenre entity.
	// It exists in this package in order to avoid circular dependency with the "moviegenre" package.
	GenresInverseTable = "movie_genres"
	// CountriesTable is the table that holds the countries relation/edge. The primary key declared below.
	CountriesTable = "country_movies"
	// CountriesInverseTable is the table name for the Country entity.
	// It exists in this package in order to avoid circular dependency with the "country" package.
	CountriesInverseTable = "countries"
)

// Columns holds all SQL columns for movie fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldOriginalTitle,
	FieldLanguages,
	FieldReleaseDate,
	FieldPlot,
	FieldDuration,
	FieldWatched,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "movies"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"movie_poster",
}

var (
	// CastPrimaryKey and CastColumn2 are the table columns denoting the
	// primary key for the cast relation (M2M).
	CastPrimaryKey = []string{"movie_id", "artist_id"}
	// DirectorsPrimaryKey and DirectorsColumn2 are the table columns denoting the
	// primary key for the directors relation (M2M).
	DirectorsPrimaryKey = []string{"movie_id", "artist_id"}
	// WritersPrimaryKey and WritersColumn2 are the table columns denoting the
	// primary key for the writers relation (M2M).
	WritersPrimaryKey = []string{"movie_id", "artist_id"}
	// GenresPrimaryKey and GenresColumn2 are the table columns denoting the
	// primary key for the genres relation (M2M).
	GenresPrimaryKey = []string{"movie_genre_id", "movie_id"}
	// CountriesPrimaryKey and CountriesColumn2 are the table columns denoting the
	// primary key for the countries relation (M2M).
	CountriesPrimaryKey = []string{"country_id", "movie_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultWatched holds the default value on creation for the "watched" field.
	DefaultWatched bool
)
