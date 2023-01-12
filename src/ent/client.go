// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/joelschutz/gomecoma/src/ent/migrate"

	"github.com/joelschutz/gomecoma/src/ent/artist"
	"github.com/joelschutz/gomecoma/src/ent/country"
	"github.com/joelschutz/gomecoma/src/ent/file"
	"github.com/joelschutz/gomecoma/src/ent/movie"
	"github.com/joelschutz/gomecoma/src/ent/moviegenre"
	"github.com/joelschutz/gomecoma/src/ent/picture"
	"github.com/joelschutz/gomecoma/src/ent/rating"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Artist is the client for interacting with the Artist builders.
	Artist *ArtistClient
	// Country is the client for interacting with the Country builders.
	Country *CountryClient
	// File is the client for interacting with the File builders.
	File *FileClient
	// Movie is the client for interacting with the Movie builders.
	Movie *MovieClient
	// MovieGenre is the client for interacting with the MovieGenre builders.
	MovieGenre *MovieGenreClient
	// Picture is the client for interacting with the Picture builders.
	Picture *PictureClient
	// Rating is the client for interacting with the Rating builders.
	Rating *RatingClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Artist = NewArtistClient(c.config)
	c.Country = NewCountryClient(c.config)
	c.File = NewFileClient(c.config)
	c.Movie = NewMovieClient(c.config)
	c.MovieGenre = NewMovieGenreClient(c.config)
	c.Picture = NewPictureClient(c.config)
	c.Rating = NewRatingClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Artist:     NewArtistClient(cfg),
		Country:    NewCountryClient(cfg),
		File:       NewFileClient(cfg),
		Movie:      NewMovieClient(cfg),
		MovieGenre: NewMovieGenreClient(cfg),
		Picture:    NewPictureClient(cfg),
		Rating:     NewRatingClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Artist:     NewArtistClient(cfg),
		Country:    NewCountryClient(cfg),
		File:       NewFileClient(cfg),
		Movie:      NewMovieClient(cfg),
		MovieGenre: NewMovieGenreClient(cfg),
		Picture:    NewPictureClient(cfg),
		Rating:     NewRatingClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Artist.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Artist.Use(hooks...)
	c.Country.Use(hooks...)
	c.File.Use(hooks...)
	c.Movie.Use(hooks...)
	c.MovieGenre.Use(hooks...)
	c.Picture.Use(hooks...)
	c.Rating.Use(hooks...)
}

// ArtistClient is a client for the Artist schema.
type ArtistClient struct {
	config
}

// NewArtistClient returns a client for the Artist from the given config.
func NewArtistClient(c config) *ArtistClient {
	return &ArtistClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `artist.Hooks(f(g(h())))`.
func (c *ArtistClient) Use(hooks ...Hook) {
	c.hooks.Artist = append(c.hooks.Artist, hooks...)
}

// Create returns a create builder for Artist.
func (c *ArtistClient) Create() *ArtistCreate {
	mutation := newArtistMutation(c.config, OpCreate)
	return &ArtistCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Artist entities.
func (c *ArtistClient) CreateBulk(builders ...*ArtistCreate) *ArtistCreateBulk {
	return &ArtistCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Artist.
func (c *ArtistClient) Update() *ArtistUpdate {
	mutation := newArtistMutation(c.config, OpUpdate)
	return &ArtistUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ArtistClient) UpdateOne(a *Artist) *ArtistUpdateOne {
	mutation := newArtistMutation(c.config, OpUpdateOne, withArtist(a))
	return &ArtistUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ArtistClient) UpdateOneID(id int) *ArtistUpdateOne {
	mutation := newArtistMutation(c.config, OpUpdateOne, withArtistID(id))
	return &ArtistUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Artist.
func (c *ArtistClient) Delete() *ArtistDelete {
	mutation := newArtistMutation(c.config, OpDelete)
	return &ArtistDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ArtistClient) DeleteOne(a *Artist) *ArtistDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ArtistClient) DeleteOneID(id int) *ArtistDeleteOne {
	builder := c.Delete().Where(artist.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ArtistDeleteOne{builder}
}

// Query returns a query builder for Artist.
func (c *ArtistClient) Query() *ArtistQuery {
	return &ArtistQuery{
		config: c.config,
	}
}

// Get returns a Artist entity by its id.
func (c *ArtistClient) Get(ctx context.Context, id int) (*Artist, error) {
	return c.Query().Where(artist.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ArtistClient) GetX(ctx context.Context, id int) *Artist {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryProfilePicture queries the profile_picture edge of a Artist.
func (c *ArtistClient) QueryProfilePicture(a *Artist) *PictureQuery {
	query := &PictureQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(artist.Table, artist.FieldID, id),
			sqlgraph.To(picture.Table, picture.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, artist.ProfilePictureTable, artist.ProfilePictureColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryPictures queries the pictures edge of a Artist.
func (c *ArtistClient) QueryPictures(a *Artist) *PictureQuery {
	query := &PictureQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(artist.Table, artist.FieldID, id),
			sqlgraph.To(picture.Table, picture.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, artist.PicturesTable, artist.PicturesColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryDirected queries the directed edge of a Artist.
func (c *ArtistClient) QueryDirected(a *Artist) *MovieQuery {
	query := &MovieQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(artist.Table, artist.FieldID, id),
			sqlgraph.To(movie.Table, movie.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, artist.DirectedTable, artist.DirectedPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryActed queries the acted edge of a Artist.
func (c *ArtistClient) QueryActed(a *Artist) *MovieQuery {
	query := &MovieQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(artist.Table, artist.FieldID, id),
			sqlgraph.To(movie.Table, movie.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, artist.ActedTable, artist.ActedPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryWrote queries the wrote edge of a Artist.
func (c *ArtistClient) QueryWrote(a *Artist) *MovieQuery {
	query := &MovieQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(artist.Table, artist.FieldID, id),
			sqlgraph.To(movie.Table, movie.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, artist.WroteTable, artist.WrotePrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCountries queries the countries edge of a Artist.
func (c *ArtistClient) QueryCountries(a *Artist) *CountryQuery {
	query := &CountryQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(artist.Table, artist.FieldID, id),
			sqlgraph.To(country.Table, country.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, artist.CountriesTable, artist.CountriesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ArtistClient) Hooks() []Hook {
	return c.hooks.Artist
}

// CountryClient is a client for the Country schema.
type CountryClient struct {
	config
}

// NewCountryClient returns a client for the Country from the given config.
func NewCountryClient(c config) *CountryClient {
	return &CountryClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `country.Hooks(f(g(h())))`.
func (c *CountryClient) Use(hooks ...Hook) {
	c.hooks.Country = append(c.hooks.Country, hooks...)
}

// Create returns a create builder for Country.
func (c *CountryClient) Create() *CountryCreate {
	mutation := newCountryMutation(c.config, OpCreate)
	return &CountryCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Country entities.
func (c *CountryClient) CreateBulk(builders ...*CountryCreate) *CountryCreateBulk {
	return &CountryCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Country.
func (c *CountryClient) Update() *CountryUpdate {
	mutation := newCountryMutation(c.config, OpUpdate)
	return &CountryUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CountryClient) UpdateOne(co *Country) *CountryUpdateOne {
	mutation := newCountryMutation(c.config, OpUpdateOne, withCountry(co))
	return &CountryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CountryClient) UpdateOneID(id int) *CountryUpdateOne {
	mutation := newCountryMutation(c.config, OpUpdateOne, withCountryID(id))
	return &CountryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Country.
func (c *CountryClient) Delete() *CountryDelete {
	mutation := newCountryMutation(c.config, OpDelete)
	return &CountryDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CountryClient) DeleteOne(co *Country) *CountryDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CountryClient) DeleteOneID(id int) *CountryDeleteOne {
	builder := c.Delete().Where(country.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CountryDeleteOne{builder}
}

// Query returns a query builder for Country.
func (c *CountryClient) Query() *CountryQuery {
	return &CountryQuery{
		config: c.config,
	}
}

// Get returns a Country entity by its id.
func (c *CountryClient) Get(ctx context.Context, id int) (*Country, error) {
	return c.Query().Where(country.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CountryClient) GetX(ctx context.Context, id int) *Country {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryMovies queries the movies edge of a Country.
func (c *CountryClient) QueryMovies(co *Country) *MovieQuery {
	query := &MovieQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(country.Table, country.FieldID, id),
			sqlgraph.To(movie.Table, movie.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, country.MoviesTable, country.MoviesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryArtists queries the artists edge of a Country.
func (c *CountryClient) QueryArtists(co *Country) *ArtistQuery {
	query := &ArtistQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(country.Table, country.FieldID, id),
			sqlgraph.To(artist.Table, artist.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, country.ArtistsTable, country.ArtistsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CountryClient) Hooks() []Hook {
	return c.hooks.Country
}

// FileClient is a client for the File schema.
type FileClient struct {
	config
}

// NewFileClient returns a client for the File from the given config.
func NewFileClient(c config) *FileClient {
	return &FileClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `file.Hooks(f(g(h())))`.
func (c *FileClient) Use(hooks ...Hook) {
	c.hooks.File = append(c.hooks.File, hooks...)
}

// Create returns a create builder for File.
func (c *FileClient) Create() *FileCreate {
	mutation := newFileMutation(c.config, OpCreate)
	return &FileCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of File entities.
func (c *FileClient) CreateBulk(builders ...*FileCreate) *FileCreateBulk {
	return &FileCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for File.
func (c *FileClient) Update() *FileUpdate {
	mutation := newFileMutation(c.config, OpUpdate)
	return &FileUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *FileClient) UpdateOne(f *File) *FileUpdateOne {
	mutation := newFileMutation(c.config, OpUpdateOne, withFile(f))
	return &FileUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *FileClient) UpdateOneID(id int) *FileUpdateOne {
	mutation := newFileMutation(c.config, OpUpdateOne, withFileID(id))
	return &FileUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for File.
func (c *FileClient) Delete() *FileDelete {
	mutation := newFileMutation(c.config, OpDelete)
	return &FileDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *FileClient) DeleteOne(f *File) *FileDeleteOne {
	return c.DeleteOneID(f.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *FileClient) DeleteOneID(id int) *FileDeleteOne {
	builder := c.Delete().Where(file.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &FileDeleteOne{builder}
}

// Query returns a query builder for File.
func (c *FileClient) Query() *FileQuery {
	return &FileQuery{
		config: c.config,
	}
}

// Get returns a File entity by its id.
func (c *FileClient) Get(ctx context.Context, id int) (*File, error) {
	return c.Query().Where(file.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *FileClient) GetX(ctx context.Context, id int) *File {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryMovie queries the movie edge of a File.
func (c *FileClient) QueryMovie(f *File) *MovieQuery {
	query := &MovieQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := f.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(file.Table, file.FieldID, id),
			sqlgraph.To(movie.Table, movie.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, file.MovieTable, file.MovieColumn),
		)
		fromV = sqlgraph.Neighbors(f.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *FileClient) Hooks() []Hook {
	return c.hooks.File
}

// MovieClient is a client for the Movie schema.
type MovieClient struct {
	config
}

// NewMovieClient returns a client for the Movie from the given config.
func NewMovieClient(c config) *MovieClient {
	return &MovieClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `movie.Hooks(f(g(h())))`.
func (c *MovieClient) Use(hooks ...Hook) {
	c.hooks.Movie = append(c.hooks.Movie, hooks...)
}

// Create returns a create builder for Movie.
func (c *MovieClient) Create() *MovieCreate {
	mutation := newMovieMutation(c.config, OpCreate)
	return &MovieCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Movie entities.
func (c *MovieClient) CreateBulk(builders ...*MovieCreate) *MovieCreateBulk {
	return &MovieCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Movie.
func (c *MovieClient) Update() *MovieUpdate {
	mutation := newMovieMutation(c.config, OpUpdate)
	return &MovieUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MovieClient) UpdateOne(m *Movie) *MovieUpdateOne {
	mutation := newMovieMutation(c.config, OpUpdateOne, withMovie(m))
	return &MovieUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MovieClient) UpdateOneID(id int) *MovieUpdateOne {
	mutation := newMovieMutation(c.config, OpUpdateOne, withMovieID(id))
	return &MovieUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Movie.
func (c *MovieClient) Delete() *MovieDelete {
	mutation := newMovieMutation(c.config, OpDelete)
	return &MovieDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *MovieClient) DeleteOne(m *Movie) *MovieDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *MovieClient) DeleteOneID(id int) *MovieDeleteOne {
	builder := c.Delete().Where(movie.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MovieDeleteOne{builder}
}

// Query returns a query builder for Movie.
func (c *MovieClient) Query() *MovieQuery {
	return &MovieQuery{
		config: c.config,
	}
}

// Get returns a Movie entity by its id.
func (c *MovieClient) Get(ctx context.Context, id int) (*Movie, error) {
	return c.Query().Where(movie.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MovieClient) GetX(ctx context.Context, id int) *Movie {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryFile queries the file edge of a Movie.
func (c *MovieClient) QueryFile(m *Movie) *FileQuery {
	query := &FileQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(movie.Table, movie.FieldID, id),
			sqlgraph.To(file.Table, file.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, movie.FileTable, movie.FileColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryRatings queries the ratings edge of a Movie.
func (c *MovieClient) QueryRatings(m *Movie) *RatingQuery {
	query := &RatingQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(movie.Table, movie.FieldID, id),
			sqlgraph.To(rating.Table, rating.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, movie.RatingsTable, movie.RatingsColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryPoster queries the poster edge of a Movie.
func (c *MovieClient) QueryPoster(m *Movie) *PictureQuery {
	query := &PictureQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(movie.Table, movie.FieldID, id),
			sqlgraph.To(picture.Table, picture.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, movie.PosterTable, movie.PosterColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryFanart queries the fanart edge of a Movie.
func (c *MovieClient) QueryFanart(m *Movie) *PictureQuery {
	query := &PictureQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(movie.Table, movie.FieldID, id),
			sqlgraph.To(picture.Table, picture.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, movie.FanartTable, movie.FanartColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCast queries the cast edge of a Movie.
func (c *MovieClient) QueryCast(m *Movie) *ArtistQuery {
	query := &ArtistQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(movie.Table, movie.FieldID, id),
			sqlgraph.To(artist.Table, artist.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, movie.CastTable, movie.CastPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryDirectors queries the directors edge of a Movie.
func (c *MovieClient) QueryDirectors(m *Movie) *ArtistQuery {
	query := &ArtistQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(movie.Table, movie.FieldID, id),
			sqlgraph.To(artist.Table, artist.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, movie.DirectorsTable, movie.DirectorsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryWriters queries the writers edge of a Movie.
func (c *MovieClient) QueryWriters(m *Movie) *ArtistQuery {
	query := &ArtistQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(movie.Table, movie.FieldID, id),
			sqlgraph.To(artist.Table, artist.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, movie.WritersTable, movie.WritersPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryGenres queries the genres edge of a Movie.
func (c *MovieClient) QueryGenres(m *Movie) *MovieGenreQuery {
	query := &MovieGenreQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(movie.Table, movie.FieldID, id),
			sqlgraph.To(moviegenre.Table, moviegenre.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, movie.GenresTable, movie.GenresPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCountries queries the countries edge of a Movie.
func (c *MovieClient) QueryCountries(m *Movie) *CountryQuery {
	query := &CountryQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(movie.Table, movie.FieldID, id),
			sqlgraph.To(country.Table, country.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, movie.CountriesTable, movie.CountriesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *MovieClient) Hooks() []Hook {
	return c.hooks.Movie
}

// MovieGenreClient is a client for the MovieGenre schema.
type MovieGenreClient struct {
	config
}

// NewMovieGenreClient returns a client for the MovieGenre from the given config.
func NewMovieGenreClient(c config) *MovieGenreClient {
	return &MovieGenreClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `moviegenre.Hooks(f(g(h())))`.
func (c *MovieGenreClient) Use(hooks ...Hook) {
	c.hooks.MovieGenre = append(c.hooks.MovieGenre, hooks...)
}

// Create returns a create builder for MovieGenre.
func (c *MovieGenreClient) Create() *MovieGenreCreate {
	mutation := newMovieGenreMutation(c.config, OpCreate)
	return &MovieGenreCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of MovieGenre entities.
func (c *MovieGenreClient) CreateBulk(builders ...*MovieGenreCreate) *MovieGenreCreateBulk {
	return &MovieGenreCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for MovieGenre.
func (c *MovieGenreClient) Update() *MovieGenreUpdate {
	mutation := newMovieGenreMutation(c.config, OpUpdate)
	return &MovieGenreUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MovieGenreClient) UpdateOne(mg *MovieGenre) *MovieGenreUpdateOne {
	mutation := newMovieGenreMutation(c.config, OpUpdateOne, withMovieGenre(mg))
	return &MovieGenreUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MovieGenreClient) UpdateOneID(id int) *MovieGenreUpdateOne {
	mutation := newMovieGenreMutation(c.config, OpUpdateOne, withMovieGenreID(id))
	return &MovieGenreUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for MovieGenre.
func (c *MovieGenreClient) Delete() *MovieGenreDelete {
	mutation := newMovieGenreMutation(c.config, OpDelete)
	return &MovieGenreDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *MovieGenreClient) DeleteOne(mg *MovieGenre) *MovieGenreDeleteOne {
	return c.DeleteOneID(mg.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *MovieGenreClient) DeleteOneID(id int) *MovieGenreDeleteOne {
	builder := c.Delete().Where(moviegenre.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MovieGenreDeleteOne{builder}
}

// Query returns a query builder for MovieGenre.
func (c *MovieGenreClient) Query() *MovieGenreQuery {
	return &MovieGenreQuery{
		config: c.config,
	}
}

// Get returns a MovieGenre entity by its id.
func (c *MovieGenreClient) Get(ctx context.Context, id int) (*MovieGenre, error) {
	return c.Query().Where(moviegenre.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MovieGenreClient) GetX(ctx context.Context, id int) *MovieGenre {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryMovies queries the movies edge of a MovieGenre.
func (c *MovieGenreClient) QueryMovies(mg *MovieGenre) *MovieQuery {
	query := &MovieQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := mg.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(moviegenre.Table, moviegenre.FieldID, id),
			sqlgraph.To(movie.Table, movie.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, moviegenre.MoviesTable, moviegenre.MoviesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(mg.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *MovieGenreClient) Hooks() []Hook {
	return c.hooks.MovieGenre
}

// PictureClient is a client for the Picture schema.
type PictureClient struct {
	config
}

// NewPictureClient returns a client for the Picture from the given config.
func NewPictureClient(c config) *PictureClient {
	return &PictureClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `picture.Hooks(f(g(h())))`.
func (c *PictureClient) Use(hooks ...Hook) {
	c.hooks.Picture = append(c.hooks.Picture, hooks...)
}

// Create returns a create builder for Picture.
func (c *PictureClient) Create() *PictureCreate {
	mutation := newPictureMutation(c.config, OpCreate)
	return &PictureCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Picture entities.
func (c *PictureClient) CreateBulk(builders ...*PictureCreate) *PictureCreateBulk {
	return &PictureCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Picture.
func (c *PictureClient) Update() *PictureUpdate {
	mutation := newPictureMutation(c.config, OpUpdate)
	return &PictureUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PictureClient) UpdateOne(pi *Picture) *PictureUpdateOne {
	mutation := newPictureMutation(c.config, OpUpdateOne, withPicture(pi))
	return &PictureUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PictureClient) UpdateOneID(id int) *PictureUpdateOne {
	mutation := newPictureMutation(c.config, OpUpdateOne, withPictureID(id))
	return &PictureUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Picture.
func (c *PictureClient) Delete() *PictureDelete {
	mutation := newPictureMutation(c.config, OpDelete)
	return &PictureDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PictureClient) DeleteOne(pi *Picture) *PictureDeleteOne {
	return c.DeleteOneID(pi.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PictureClient) DeleteOneID(id int) *PictureDeleteOne {
	builder := c.Delete().Where(picture.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PictureDeleteOne{builder}
}

// Query returns a query builder for Picture.
func (c *PictureClient) Query() *PictureQuery {
	return &PictureQuery{
		config: c.config,
	}
}

// Get returns a Picture entity by its id.
func (c *PictureClient) Get(ctx context.Context, id int) (*Picture, error) {
	return c.Query().Where(picture.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PictureClient) GetX(ctx context.Context, id int) *Picture {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *PictureClient) Hooks() []Hook {
	return c.hooks.Picture
}

// RatingClient is a client for the Rating schema.
type RatingClient struct {
	config
}

// NewRatingClient returns a client for the Rating from the given config.
func NewRatingClient(c config) *RatingClient {
	return &RatingClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `rating.Hooks(f(g(h())))`.
func (c *RatingClient) Use(hooks ...Hook) {
	c.hooks.Rating = append(c.hooks.Rating, hooks...)
}

// Create returns a create builder for Rating.
func (c *RatingClient) Create() *RatingCreate {
	mutation := newRatingMutation(c.config, OpCreate)
	return &RatingCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Rating entities.
func (c *RatingClient) CreateBulk(builders ...*RatingCreate) *RatingCreateBulk {
	return &RatingCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Rating.
func (c *RatingClient) Update() *RatingUpdate {
	mutation := newRatingMutation(c.config, OpUpdate)
	return &RatingUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RatingClient) UpdateOne(r *Rating) *RatingUpdateOne {
	mutation := newRatingMutation(c.config, OpUpdateOne, withRating(r))
	return &RatingUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RatingClient) UpdateOneID(id int) *RatingUpdateOne {
	mutation := newRatingMutation(c.config, OpUpdateOne, withRatingID(id))
	return &RatingUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Rating.
func (c *RatingClient) Delete() *RatingDelete {
	mutation := newRatingMutation(c.config, OpDelete)
	return &RatingDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *RatingClient) DeleteOne(r *Rating) *RatingDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *RatingClient) DeleteOneID(id int) *RatingDeleteOne {
	builder := c.Delete().Where(rating.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RatingDeleteOne{builder}
}

// Query returns a query builder for Rating.
func (c *RatingClient) Query() *RatingQuery {
	return &RatingQuery{
		config: c.config,
	}
}

// Get returns a Rating entity by its id.
func (c *RatingClient) Get(ctx context.Context, id int) (*Rating, error) {
	return c.Query().Where(rating.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RatingClient) GetX(ctx context.Context, id int) *Rating {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *RatingClient) Hooks() []Hook {
	return c.hooks.Rating
}
