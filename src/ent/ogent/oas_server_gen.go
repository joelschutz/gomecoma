// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"math/big"
	"math/bits"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = url.URL{}
	_ = math.Mod
	_ = bits.LeadingZeros64
	_ = big.Rat{}
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = attribute.KeyValue{}
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
	_ = regexp.MustCompile
	_ = jx.Null
	_ = sync.Pool{}
	_ = codes.Unset
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// CreateArtist implements createArtist operation.
	//
	// Creates a new Artist and persists it to storage.
	//
	// POST /artists
	CreateArtist(ctx context.Context, req CreateArtistReq) (CreateArtistRes, error)
	// CreateCountry implements createCountry operation.
	//
	// Creates a new Country and persists it to storage.
	//
	// POST /countries
	CreateCountry(ctx context.Context, req CreateCountryReq) (CreateCountryRes, error)
	// CreateFile implements createFile operation.
	//
	// Creates a new File and persists it to storage.
	//
	// POST /files
	CreateFile(ctx context.Context, req CreateFileReq) (CreateFileRes, error)
	// CreateMovie implements createMovie operation.
	//
	// Creates a new Movie and persists it to storage.
	//
	// POST /movies
	CreateMovie(ctx context.Context, req CreateMovieReq) (CreateMovieRes, error)
	// CreateMovieGenre implements createMovieGenre operation.
	//
	// Creates a new MovieGenre and persists it to storage.
	//
	// POST /movie-genres
	CreateMovieGenre(ctx context.Context, req CreateMovieGenreReq) (CreateMovieGenreRes, error)
	// CreatePicture implements createPicture operation.
	//
	// Creates a new Picture and persists it to storage.
	//
	// POST /pictures
	CreatePicture(ctx context.Context, req CreatePictureReq) (CreatePictureRes, error)
	// CreateRating implements createRating operation.
	//
	// Creates a new Rating and persists it to storage.
	//
	// POST /ratings
	CreateRating(ctx context.Context, req CreateRatingReq) (CreateRatingRes, error)
	// DeleteArtist implements deleteArtist operation.
	//
	// Deletes the Artist with the requested ID.
	//
	// DELETE /artists/{id}
	DeleteArtist(ctx context.Context, params DeleteArtistParams) (DeleteArtistRes, error)
	// DeleteCountry implements deleteCountry operation.
	//
	// Deletes the Country with the requested ID.
	//
	// DELETE /countries/{id}
	DeleteCountry(ctx context.Context, params DeleteCountryParams) (DeleteCountryRes, error)
	// DeleteFile implements deleteFile operation.
	//
	// Deletes the File with the requested ID.
	//
	// DELETE /files/{id}
	DeleteFile(ctx context.Context, params DeleteFileParams) (DeleteFileRes, error)
	// DeleteMovie implements deleteMovie operation.
	//
	// Deletes the Movie with the requested ID.
	//
	// DELETE /movies/{id}
	DeleteMovie(ctx context.Context, params DeleteMovieParams) (DeleteMovieRes, error)
	// DeleteMovieGenre implements deleteMovieGenre operation.
	//
	// Deletes the MovieGenre with the requested ID.
	//
	// DELETE /movie-genres/{id}
	DeleteMovieGenre(ctx context.Context, params DeleteMovieGenreParams) (DeleteMovieGenreRes, error)
	// DeletePicture implements deletePicture operation.
	//
	// Deletes the Picture with the requested ID.
	//
	// DELETE /pictures/{id}
	DeletePicture(ctx context.Context, params DeletePictureParams) (DeletePictureRes, error)
	// DeleteRating implements deleteRating operation.
	//
	// Deletes the Rating with the requested ID.
	//
	// DELETE /ratings/{id}
	DeleteRating(ctx context.Context, params DeleteRatingParams) (DeleteRatingRes, error)
	// ListArtist implements listArtist operation.
	//
	// List Artists.
	//
	// GET /artists
	ListArtist(ctx context.Context, params ListArtistParams) (ListArtistRes, error)
	// ListArtistActed implements listArtistActed operation.
	//
	// List attached Acteds.
	//
	// GET /artists/{id}/acted
	ListArtistActed(ctx context.Context, params ListArtistActedParams) (ListArtistActedRes, error)
	// ListArtistCountries implements listArtistCountries operation.
	//
	// List attached Countries.
	//
	// GET /artists/{id}/countries
	ListArtistCountries(ctx context.Context, params ListArtistCountriesParams) (ListArtistCountriesRes, error)
	// ListArtistDirected implements listArtistDirected operation.
	//
	// List attached Directeds.
	//
	// GET /artists/{id}/directed
	ListArtistDirected(ctx context.Context, params ListArtistDirectedParams) (ListArtistDirectedRes, error)
	// ListArtistPictures implements listArtistPictures operation.
	//
	// List attached Pictures.
	//
	// GET /artists/{id}/pictures
	ListArtistPictures(ctx context.Context, params ListArtistPicturesParams) (ListArtistPicturesRes, error)
	// ListArtistWrote implements listArtistWrote operation.
	//
	// List attached Wrotes.
	//
	// GET /artists/{id}/wrote
	ListArtistWrote(ctx context.Context, params ListArtistWroteParams) (ListArtistWroteRes, error)
	// ListCountry implements listCountry operation.
	//
	// List Countries.
	//
	// GET /countries
	ListCountry(ctx context.Context, params ListCountryParams) (ListCountryRes, error)
	// ListCountryArtists implements listCountryArtists operation.
	//
	// List attached Artists.
	//
	// GET /countries/{id}/artists
	ListCountryArtists(ctx context.Context, params ListCountryArtistsParams) (ListCountryArtistsRes, error)
	// ListCountryMovies implements listCountryMovies operation.
	//
	// List attached Movies.
	//
	// GET /countries/{id}/movies
	ListCountryMovies(ctx context.Context, params ListCountryMoviesParams) (ListCountryMoviesRes, error)
	// ListFile implements listFile operation.
	//
	// List Files.
	//
	// GET /files
	ListFile(ctx context.Context, params ListFileParams) (ListFileRes, error)
	// ListMovie implements listMovie operation.
	//
	// List Movies.
	//
	// GET /movies
	ListMovie(ctx context.Context, params ListMovieParams) (ListMovieRes, error)
	// ListMovieCast implements listMovieCast operation.
	//
	// List attached Casts.
	//
	// GET /movies/{id}/cast
	ListMovieCast(ctx context.Context, params ListMovieCastParams) (ListMovieCastRes, error)
	// ListMovieCountries implements listMovieCountries operation.
	//
	// List attached Countries.
	//
	// GET /movies/{id}/countries
	ListMovieCountries(ctx context.Context, params ListMovieCountriesParams) (ListMovieCountriesRes, error)
	// ListMovieDirectors implements listMovieDirectors operation.
	//
	// List attached Directors.
	//
	// GET /movies/{id}/directors
	ListMovieDirectors(ctx context.Context, params ListMovieDirectorsParams) (ListMovieDirectorsRes, error)
	// ListMovieFanart implements listMovieFanart operation.
	//
	// List attached Fanarts.
	//
	// GET /movies/{id}/fanart
	ListMovieFanart(ctx context.Context, params ListMovieFanartParams) (ListMovieFanartRes, error)
	// ListMovieGenre implements listMovieGenre operation.
	//
	// List MovieGenres.
	//
	// GET /movie-genres
	ListMovieGenre(ctx context.Context, params ListMovieGenreParams) (ListMovieGenreRes, error)
	// ListMovieGenreMovies implements listMovieGenreMovies operation.
	//
	// List attached Movies.
	//
	// GET /movie-genres/{id}/movies
	ListMovieGenreMovies(ctx context.Context, params ListMovieGenreMoviesParams) (ListMovieGenreMoviesRes, error)
	// ListMovieGenres implements listMovieGenres operation.
	//
	// List attached Genres.
	//
	// GET /movies/{id}/genres
	ListMovieGenres(ctx context.Context, params ListMovieGenresParams) (ListMovieGenresRes, error)
	// ListMovieRatings implements listMovieRatings operation.
	//
	// List attached Ratings.
	//
	// GET /movies/{id}/ratings
	ListMovieRatings(ctx context.Context, params ListMovieRatingsParams) (ListMovieRatingsRes, error)
	// ListMovieWriters implements listMovieWriters operation.
	//
	// List attached Writers.
	//
	// GET /movies/{id}/writers
	ListMovieWriters(ctx context.Context, params ListMovieWritersParams) (ListMovieWritersRes, error)
	// ListPicture implements listPicture operation.
	//
	// List Pictures.
	//
	// GET /pictures
	ListPicture(ctx context.Context, params ListPictureParams) (ListPictureRes, error)
	// ListRating implements listRating operation.
	//
	// List Ratings.
	//
	// GET /ratings
	ListRating(ctx context.Context, params ListRatingParams) (ListRatingRes, error)
	// ReadArtist implements readArtist operation.
	//
	// Finds the Artist with the requested ID and returns it.
	//
	// GET /artists/{id}
	ReadArtist(ctx context.Context, params ReadArtistParams) (ReadArtistRes, error)
	// ReadArtistProfilePicture implements readArtistProfilePicture operation.
	//
	// Find the attached Picture of the Artist with the given ID.
	//
	// GET /artists/{id}/profile-picture
	ReadArtistProfilePicture(ctx context.Context, params ReadArtistProfilePictureParams) (ReadArtistProfilePictureRes, error)
	// ReadCountry implements readCountry operation.
	//
	// Finds the Country with the requested ID and returns it.
	//
	// GET /countries/{id}
	ReadCountry(ctx context.Context, params ReadCountryParams) (ReadCountryRes, error)
	// ReadFile implements readFile operation.
	//
	// Finds the File with the requested ID and returns it.
	//
	// GET /files/{id}
	ReadFile(ctx context.Context, params ReadFileParams) (ReadFileRes, error)
	// ReadFileMovie implements readFileMovie operation.
	//
	// Find the attached Movie of the File with the given ID.
	//
	// GET /files/{id}/movie
	ReadFileMovie(ctx context.Context, params ReadFileMovieParams) (ReadFileMovieRes, error)
	// ReadMovie implements readMovie operation.
	//
	// Finds the Movie with the requested ID and returns it.
	//
	// GET /movies/{id}
	ReadMovie(ctx context.Context, params ReadMovieParams) (ReadMovieRes, error)
	// ReadMovieFile implements readMovieFile operation.
	//
	// Find the attached File of the Movie with the given ID.
	//
	// GET /movies/{id}/file
	ReadMovieFile(ctx context.Context, params ReadMovieFileParams) (ReadMovieFileRes, error)
	// ReadMovieGenre implements readMovieGenre operation.
	//
	// Finds the MovieGenre with the requested ID and returns it.
	//
	// GET /movie-genres/{id}
	ReadMovieGenre(ctx context.Context, params ReadMovieGenreParams) (ReadMovieGenreRes, error)
	// ReadMoviePoster implements readMoviePoster operation.
	//
	// Find the attached Picture of the Movie with the given ID.
	//
	// GET /movies/{id}/poster
	ReadMoviePoster(ctx context.Context, params ReadMoviePosterParams) (ReadMoviePosterRes, error)
	// ReadPicture implements readPicture operation.
	//
	// Finds the Picture with the requested ID and returns it.
	//
	// GET /pictures/{id}
	ReadPicture(ctx context.Context, params ReadPictureParams) (ReadPictureRes, error)
	// ReadRating implements readRating operation.
	//
	// Finds the Rating with the requested ID and returns it.
	//
	// GET /ratings/{id}
	ReadRating(ctx context.Context, params ReadRatingParams) (ReadRatingRes, error)
	// UpdateArtist implements updateArtist operation.
	//
	// Updates a Artist and persists changes to storage.
	//
	// PATCH /artists/{id}
	UpdateArtist(ctx context.Context, req UpdateArtistReq, params UpdateArtistParams) (UpdateArtistRes, error)
	// UpdateCountry implements updateCountry operation.
	//
	// Updates a Country and persists changes to storage.
	//
	// PATCH /countries/{id}
	UpdateCountry(ctx context.Context, req UpdateCountryReq, params UpdateCountryParams) (UpdateCountryRes, error)
	// UpdateFile implements updateFile operation.
	//
	// Updates a File and persists changes to storage.
	//
	// PATCH /files/{id}
	UpdateFile(ctx context.Context, req UpdateFileReq, params UpdateFileParams) (UpdateFileRes, error)
	// UpdateMovie implements updateMovie operation.
	//
	// Updates a Movie and persists changes to storage.
	//
	// PATCH /movies/{id}
	UpdateMovie(ctx context.Context, req UpdateMovieReq, params UpdateMovieParams) (UpdateMovieRes, error)
	// UpdateMovieGenre implements updateMovieGenre operation.
	//
	// Updates a MovieGenre and persists changes to storage.
	//
	// PATCH /movie-genres/{id}
	UpdateMovieGenre(ctx context.Context, req UpdateMovieGenreReq, params UpdateMovieGenreParams) (UpdateMovieGenreRes, error)
	// UpdatePicture implements updatePicture operation.
	//
	// Updates a Picture and persists changes to storage.
	//
	// PATCH /pictures/{id}
	UpdatePicture(ctx context.Context, req UpdatePictureReq, params UpdatePictureParams) (UpdatePictureRes, error)
	// UpdateRating implements updateRating operation.
	//
	// Updates a Rating and persists changes to storage.
	//
	// PATCH /ratings/{id}
	UpdateRating(ctx context.Context, req UpdateRatingReq, params UpdateRatingParams) (UpdateRatingRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	cfg config

	requests metric.Int64Counter
	errors   metric.Int64Counter
	duration metric.Int64Histogram
}

func NewServer(h Handler, opts ...Option) (*Server, error) {
	s := &Server{
		h:   h,
		cfg: newConfig(opts...),
	}
	var err error
	if s.requests, err = s.cfg.Meter.NewInt64Counter(otelogen.ServerRequestCount); err != nil {
		return nil, err
	}
	if s.errors, err = s.cfg.Meter.NewInt64Counter(otelogen.ServerErrorsCount); err != nil {
		return nil, err
	}
	if s.duration, err = s.cfg.Meter.NewInt64Histogram(otelogen.ServerDuration); err != nil {
		return nil, err
	}
	return s, nil
}
