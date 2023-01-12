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

type DeleteArtistParams struct {
	// ID of the Artist.
	ID int
}

type DeleteCountryParams struct {
	// ID of the Country.
	ID int
}

type DeleteFileParams struct {
	// ID of the File.
	ID int
}

type DeleteMovieParams struct {
	// ID of the Movie.
	ID int
}

type DeleteMovieGenreParams struct {
	// ID of the MovieGenre.
	ID int
}

type DeletePictureParams struct {
	// ID of the Picture.
	ID int
}

type DeleteRatingParams struct {
	// ID of the Rating.
	ID int
}

type ListArtistParams struct {
	// What page to render.
	Page OptInt
	// Item count to render per page.
	ItemsPerPage OptInt
}

type ListArtistActedParams struct {
	// ID of the Artist.
	ID int
	// What page to render.
	Page OptInt
	// Item count to render per page.
	ItemsPerPage OptInt
}

type ListArtistCountriesParams struct {
	// ID of the Artist.
	ID int
	// What page to render.
	Page OptInt
	// Item count to render per page.
	ItemsPerPage OptInt
}

type ListArtistDirectedParams struct {
	// ID of the Artist.
	ID int
	// What page to render.
	Page OptInt
	// Item count to render per page.
	ItemsPerPage OptInt
}

type ListArtistPicturesParams struct {
	// ID of the Artist.
	ID int
	// What page to render.
	Page OptInt
	// Item count to render per page.
	ItemsPerPage OptInt
}

type ListArtistWroteParams struct {
	// ID of the Artist.
	ID int
	// What page to render.
	Page OptInt
	// Item count to render per page.
	ItemsPerPage OptInt
}

type ListCountryParams struct {
	// What page to render.
	Page OptInt
	// Item count to render per page.
	ItemsPerPage OptInt
}

type ListCountryArtistsParams struct {
	// ID of the Country.
	ID int
	// What page to render.
	Page OptInt
	// Item count to render per page.
	ItemsPerPage OptInt
}

type ListCountryMoviesParams struct {
	// ID of the Country.
	ID int
	// What page to render.
	Page OptInt
	// Item count to render per page.
	ItemsPerPage OptInt
}

type ListFileParams struct {
	// What page to render.
	Page OptInt
	// Item count to render per page.
	ItemsPerPage OptInt
}

type ListMovieParams struct {
	// What page to render.
	Page OptInt
	// Item count to render per page.
	ItemsPerPage OptInt
}

type ListMovieCastParams struct {
	// ID of the Movie.
	ID int
	// What page to render.
	Page OptInt
	// Item count to render per page.
	ItemsPerPage OptInt
}

type ListMovieCountriesParams struct {
	// ID of the Movie.
	ID int
	// What page to render.
	Page OptInt
	// Item count to render per page.
	ItemsPerPage OptInt
}

type ListMovieDirectorsParams struct {
	// ID of the Movie.
	ID int
	// What page to render.
	Page OptInt
	// Item count to render per page.
	ItemsPerPage OptInt
}

type ListMovieFanartParams struct {
	// ID of the Movie.
	ID int
	// What page to render.
	Page OptInt
	// Item count to render per page.
	ItemsPerPage OptInt
}

type ListMovieGenreParams struct {
	// What page to render.
	Page OptInt
	// Item count to render per page.
	ItemsPerPage OptInt
}

type ListMovieGenreMoviesParams struct {
	// ID of the MovieGenre.
	ID int
	// What page to render.
	Page OptInt
	// Item count to render per page.
	ItemsPerPage OptInt
}

type ListMovieGenresParams struct {
	// ID of the Movie.
	ID int
	// What page to render.
	Page OptInt
	// Item count to render per page.
	ItemsPerPage OptInt
}

type ListMovieRatingsParams struct {
	// ID of the Movie.
	ID int
	// What page to render.
	Page OptInt
	// Item count to render per page.
	ItemsPerPage OptInt
}

type ListMovieWritersParams struct {
	// ID of the Movie.
	ID int
	// What page to render.
	Page OptInt
	// Item count to render per page.
	ItemsPerPage OptInt
}

type ListPictureParams struct {
	// What page to render.
	Page OptInt
	// Item count to render per page.
	ItemsPerPage OptInt
}

type ListRatingParams struct {
	// What page to render.
	Page OptInt
	// Item count to render per page.
	ItemsPerPage OptInt
}

type ReadArtistParams struct {
	// ID of the Artist.
	ID int
}

type ReadArtistProfilePictureParams struct {
	// ID of the Artist.
	ID int
}

type ReadCountryParams struct {
	// ID of the Country.
	ID int
}

type ReadFileParams struct {
	// ID of the File.
	ID int
}

type ReadFileMovieParams struct {
	// ID of the File.
	ID int
}

type ReadMovieParams struct {
	// ID of the Movie.
	ID int
}

type ReadMovieFileParams struct {
	// ID of the Movie.
	ID int
}

type ReadMovieGenreParams struct {
	// ID of the MovieGenre.
	ID int
}

type ReadMoviePosterParams struct {
	// ID of the Movie.
	ID int
}

type ReadPictureParams struct {
	// ID of the Picture.
	ID int
}

type ReadRatingParams struct {
	// ID of the Rating.
	ID int
}

type UpdateArtistParams struct {
	// ID of the Artist.
	ID int
}

type UpdateCountryParams struct {
	// ID of the Country.
	ID int
}

type UpdateFileParams struct {
	// ID of the File.
	ID int
}

type UpdateMovieParams struct {
	// ID of the Movie.
	ID int
}

type UpdateMovieGenreParams struct {
	// ID of the MovieGenre.
	ID int
}

type UpdatePictureParams struct {
	// ID of the Picture.
	ID int
}

type UpdateRatingParams struct {
	// ID of the Rating.
	ID int
}