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

var _ Handler = UnimplementedHandler{}

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

// CreateArtist implements createArtist operation.
//
// Creates a new Artist and persists it to storage.
//
// POST /artists
func (UnimplementedHandler) CreateArtist(ctx context.Context, req CreateArtistReq) (r CreateArtistRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateCountry implements createCountry operation.
//
// Creates a new Country and persists it to storage.
//
// POST /countries
func (UnimplementedHandler) CreateCountry(ctx context.Context, req CreateCountryReq) (r CreateCountryRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateFile implements createFile operation.
//
// Creates a new File and persists it to storage.
//
// POST /files
func (UnimplementedHandler) CreateFile(ctx context.Context, req CreateFileReq) (r CreateFileRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateMovie implements createMovie operation.
//
// Creates a new Movie and persists it to storage.
//
// POST /movies
func (UnimplementedHandler) CreateMovie(ctx context.Context, req CreateMovieReq) (r CreateMovieRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateMovieGenre implements createMovieGenre operation.
//
// Creates a new MovieGenre and persists it to storage.
//
// POST /movie-genres
func (UnimplementedHandler) CreateMovieGenre(ctx context.Context, req CreateMovieGenreReq) (r CreateMovieGenreRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreatePicture implements createPicture operation.
//
// Creates a new Picture and persists it to storage.
//
// POST /pictures
func (UnimplementedHandler) CreatePicture(ctx context.Context, req CreatePictureReq) (r CreatePictureRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateRating implements createRating operation.
//
// Creates a new Rating and persists it to storage.
//
// POST /ratings
func (UnimplementedHandler) CreateRating(ctx context.Context, req CreateRatingReq) (r CreateRatingRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteArtist implements deleteArtist operation.
//
// Deletes the Artist with the requested ID.
//
// DELETE /artists/{id}
func (UnimplementedHandler) DeleteArtist(ctx context.Context, params DeleteArtistParams) (r DeleteArtistRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteCountry implements deleteCountry operation.
//
// Deletes the Country with the requested ID.
//
// DELETE /countries/{id}
func (UnimplementedHandler) DeleteCountry(ctx context.Context, params DeleteCountryParams) (r DeleteCountryRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteFile implements deleteFile operation.
//
// Deletes the File with the requested ID.
//
// DELETE /files/{id}
func (UnimplementedHandler) DeleteFile(ctx context.Context, params DeleteFileParams) (r DeleteFileRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteMovie implements deleteMovie operation.
//
// Deletes the Movie with the requested ID.
//
// DELETE /movies/{id}
func (UnimplementedHandler) DeleteMovie(ctx context.Context, params DeleteMovieParams) (r DeleteMovieRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteMovieGenre implements deleteMovieGenre operation.
//
// Deletes the MovieGenre with the requested ID.
//
// DELETE /movie-genres/{id}
func (UnimplementedHandler) DeleteMovieGenre(ctx context.Context, params DeleteMovieGenreParams) (r DeleteMovieGenreRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeletePicture implements deletePicture operation.
//
// Deletes the Picture with the requested ID.
//
// DELETE /pictures/{id}
func (UnimplementedHandler) DeletePicture(ctx context.Context, params DeletePictureParams) (r DeletePictureRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteRating implements deleteRating operation.
//
// Deletes the Rating with the requested ID.
//
// DELETE /ratings/{id}
func (UnimplementedHandler) DeleteRating(ctx context.Context, params DeleteRatingParams) (r DeleteRatingRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListArtist implements listArtist operation.
//
// List Artists.
//
// GET /artists
func (UnimplementedHandler) ListArtist(ctx context.Context, params ListArtistParams) (r ListArtistRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListArtistActed implements listArtistActed operation.
//
// List attached Acteds.
//
// GET /artists/{id}/acted
func (UnimplementedHandler) ListArtistActed(ctx context.Context, params ListArtistActedParams) (r ListArtistActedRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListArtistCountries implements listArtistCountries operation.
//
// List attached Countries.
//
// GET /artists/{id}/countries
func (UnimplementedHandler) ListArtistCountries(ctx context.Context, params ListArtistCountriesParams) (r ListArtistCountriesRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListArtistDirected implements listArtistDirected operation.
//
// List attached Directeds.
//
// GET /artists/{id}/directed
func (UnimplementedHandler) ListArtistDirected(ctx context.Context, params ListArtistDirectedParams) (r ListArtistDirectedRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListArtistPictures implements listArtistPictures operation.
//
// List attached Pictures.
//
// GET /artists/{id}/pictures
func (UnimplementedHandler) ListArtistPictures(ctx context.Context, params ListArtistPicturesParams) (r ListArtistPicturesRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListArtistWrote implements listArtistWrote operation.
//
// List attached Wrotes.
//
// GET /artists/{id}/wrote
func (UnimplementedHandler) ListArtistWrote(ctx context.Context, params ListArtistWroteParams) (r ListArtistWroteRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListCountry implements listCountry operation.
//
// List Countries.
//
// GET /countries
func (UnimplementedHandler) ListCountry(ctx context.Context, params ListCountryParams) (r ListCountryRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListCountryArtists implements listCountryArtists operation.
//
// List attached Artists.
//
// GET /countries/{id}/artists
func (UnimplementedHandler) ListCountryArtists(ctx context.Context, params ListCountryArtistsParams) (r ListCountryArtistsRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListCountryMovies implements listCountryMovies operation.
//
// List attached Movies.
//
// GET /countries/{id}/movies
func (UnimplementedHandler) ListCountryMovies(ctx context.Context, params ListCountryMoviesParams) (r ListCountryMoviesRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListFile implements listFile operation.
//
// List Files.
//
// GET /files
func (UnimplementedHandler) ListFile(ctx context.Context, params ListFileParams) (r ListFileRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListMovie implements listMovie operation.
//
// List Movies.
//
// GET /movies
func (UnimplementedHandler) ListMovie(ctx context.Context, params ListMovieParams) (r ListMovieRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListMovieCast implements listMovieCast operation.
//
// List attached Casts.
//
// GET /movies/{id}/cast
func (UnimplementedHandler) ListMovieCast(ctx context.Context, params ListMovieCastParams) (r ListMovieCastRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListMovieCountries implements listMovieCountries operation.
//
// List attached Countries.
//
// GET /movies/{id}/countries
func (UnimplementedHandler) ListMovieCountries(ctx context.Context, params ListMovieCountriesParams) (r ListMovieCountriesRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListMovieDirectors implements listMovieDirectors operation.
//
// List attached Directors.
//
// GET /movies/{id}/directors
func (UnimplementedHandler) ListMovieDirectors(ctx context.Context, params ListMovieDirectorsParams) (r ListMovieDirectorsRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListMovieFanart implements listMovieFanart operation.
//
// List attached Fanarts.
//
// GET /movies/{id}/fanart
func (UnimplementedHandler) ListMovieFanart(ctx context.Context, params ListMovieFanartParams) (r ListMovieFanartRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListMovieGenre implements listMovieGenre operation.
//
// List MovieGenres.
//
// GET /movie-genres
func (UnimplementedHandler) ListMovieGenre(ctx context.Context, params ListMovieGenreParams) (r ListMovieGenreRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListMovieGenreMovies implements listMovieGenreMovies operation.
//
// List attached Movies.
//
// GET /movie-genres/{id}/movies
func (UnimplementedHandler) ListMovieGenreMovies(ctx context.Context, params ListMovieGenreMoviesParams) (r ListMovieGenreMoviesRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListMovieGenres implements listMovieGenres operation.
//
// List attached Genres.
//
// GET /movies/{id}/genres
func (UnimplementedHandler) ListMovieGenres(ctx context.Context, params ListMovieGenresParams) (r ListMovieGenresRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListMovieRatings implements listMovieRatings operation.
//
// List attached Ratings.
//
// GET /movies/{id}/ratings
func (UnimplementedHandler) ListMovieRatings(ctx context.Context, params ListMovieRatingsParams) (r ListMovieRatingsRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListMovieWriters implements listMovieWriters operation.
//
// List attached Writers.
//
// GET /movies/{id}/writers
func (UnimplementedHandler) ListMovieWriters(ctx context.Context, params ListMovieWritersParams) (r ListMovieWritersRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListPicture implements listPicture operation.
//
// List Pictures.
//
// GET /pictures
func (UnimplementedHandler) ListPicture(ctx context.Context, params ListPictureParams) (r ListPictureRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListRating implements listRating operation.
//
// List Ratings.
//
// GET /ratings
func (UnimplementedHandler) ListRating(ctx context.Context, params ListRatingParams) (r ListRatingRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadArtist implements readArtist operation.
//
// Finds the Artist with the requested ID and returns it.
//
// GET /artists/{id}
func (UnimplementedHandler) ReadArtist(ctx context.Context, params ReadArtistParams) (r ReadArtistRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadArtistProfilePicture implements readArtistProfilePicture operation.
//
// Find the attached Picture of the Artist with the given ID.
//
// GET /artists/{id}/profile-picture
func (UnimplementedHandler) ReadArtistProfilePicture(ctx context.Context, params ReadArtistProfilePictureParams) (r ReadArtistProfilePictureRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadCountry implements readCountry operation.
//
// Finds the Country with the requested ID and returns it.
//
// GET /countries/{id}
func (UnimplementedHandler) ReadCountry(ctx context.Context, params ReadCountryParams) (r ReadCountryRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadFile implements readFile operation.
//
// Finds the File with the requested ID and returns it.
//
// GET /files/{id}
func (UnimplementedHandler) ReadFile(ctx context.Context, params ReadFileParams) (r ReadFileRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadFileMovie implements readFileMovie operation.
//
// Find the attached Movie of the File with the given ID.
//
// GET /files/{id}/movie
func (UnimplementedHandler) ReadFileMovie(ctx context.Context, params ReadFileMovieParams) (r ReadFileMovieRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadMovie implements readMovie operation.
//
// Finds the Movie with the requested ID and returns it.
//
// GET /movies/{id}
func (UnimplementedHandler) ReadMovie(ctx context.Context, params ReadMovieParams) (r ReadMovieRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadMovieFile implements readMovieFile operation.
//
// Find the attached File of the Movie with the given ID.
//
// GET /movies/{id}/file
func (UnimplementedHandler) ReadMovieFile(ctx context.Context, params ReadMovieFileParams) (r ReadMovieFileRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadMovieGenre implements readMovieGenre operation.
//
// Finds the MovieGenre with the requested ID and returns it.
//
// GET /movie-genres/{id}
func (UnimplementedHandler) ReadMovieGenre(ctx context.Context, params ReadMovieGenreParams) (r ReadMovieGenreRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadMoviePoster implements readMoviePoster operation.
//
// Find the attached Picture of the Movie with the given ID.
//
// GET /movies/{id}/poster
func (UnimplementedHandler) ReadMoviePoster(ctx context.Context, params ReadMoviePosterParams) (r ReadMoviePosterRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadPicture implements readPicture operation.
//
// Finds the Picture with the requested ID and returns it.
//
// GET /pictures/{id}
func (UnimplementedHandler) ReadPicture(ctx context.Context, params ReadPictureParams) (r ReadPictureRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadRating implements readRating operation.
//
// Finds the Rating with the requested ID and returns it.
//
// GET /ratings/{id}
func (UnimplementedHandler) ReadRating(ctx context.Context, params ReadRatingParams) (r ReadRatingRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateArtist implements updateArtist operation.
//
// Updates a Artist and persists changes to storage.
//
// PATCH /artists/{id}
func (UnimplementedHandler) UpdateArtist(ctx context.Context, req UpdateArtistReq, params UpdateArtistParams) (r UpdateArtistRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateCountry implements updateCountry operation.
//
// Updates a Country and persists changes to storage.
//
// PATCH /countries/{id}
func (UnimplementedHandler) UpdateCountry(ctx context.Context, req UpdateCountryReq, params UpdateCountryParams) (r UpdateCountryRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateFile implements updateFile operation.
//
// Updates a File and persists changes to storage.
//
// PATCH /files/{id}
func (UnimplementedHandler) UpdateFile(ctx context.Context, req UpdateFileReq, params UpdateFileParams) (r UpdateFileRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateMovie implements updateMovie operation.
//
// Updates a Movie and persists changes to storage.
//
// PATCH /movies/{id}
func (UnimplementedHandler) UpdateMovie(ctx context.Context, req UpdateMovieReq, params UpdateMovieParams) (r UpdateMovieRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateMovieGenre implements updateMovieGenre operation.
//
// Updates a MovieGenre and persists changes to storage.
//
// PATCH /movie-genres/{id}
func (UnimplementedHandler) UpdateMovieGenre(ctx context.Context, req UpdateMovieGenreReq, params UpdateMovieGenreParams) (r UpdateMovieGenreRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdatePicture implements updatePicture operation.
//
// Updates a Picture and persists changes to storage.
//
// PATCH /pictures/{id}
func (UnimplementedHandler) UpdatePicture(ctx context.Context, req UpdatePictureReq, params UpdatePictureParams) (r UpdatePictureRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateRating implements updateRating operation.
//
// Updates a Rating and persists changes to storage.
//
// PATCH /ratings/{id}
func (UnimplementedHandler) UpdateRating(ctx context.Context, req UpdateRatingReq, params UpdateRatingParams) (r UpdateRatingRes, _ error) {
	return r, ht.ErrNotImplemented
}
