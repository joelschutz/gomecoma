// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/joelschutz/gomecoma/src/ent/file"
	"github.com/joelschutz/gomecoma/src/ent/movie"
	"github.com/joelschutz/gomecoma/src/ent/rating"
	"github.com/joelschutz/gomecoma/src/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	fileFields := schema.File{}.Fields()
	_ = fileFields
	// fileDescSynced is the schema descriptor for synced field.
	fileDescSynced := fileFields[6].Descriptor()
	// file.DefaultSynced holds the default value on creation for the synced field.
	file.DefaultSynced = fileDescSynced.Default.(bool)
	movieFields := schema.Movie{}.Fields()
	_ = movieFields
	// movieDescWatched is the schema descriptor for watched field.
	movieDescWatched := movieFields[5].Descriptor()
	// movie.DefaultWatched holds the default value on creation for the watched field.
	movie.DefaultWatched = movieDescWatched.Default.(bool)
	ratingFields := schema.Rating{}.Fields()
	_ = ratingFields
	// ratingDescNormalizedRating is the schema descriptor for normalized_rating field.
	ratingDescNormalizedRating := ratingFields[2].Descriptor()
	// rating.NormalizedRatingValidator is a validator for the "normalized_rating" field. It is called by the builders before save.
	rating.NormalizedRatingValidator = func() func(int) error {
		validators := ratingDescNormalizedRating.Validators
		fns := [...]func(int) error{
			validators[0].(func(int) error),
			validators[1].(func(int) error),
		}
		return func(normalized_rating int) error {
			for _, fn := range fns {
				if err := fn(normalized_rating); err != nil {
					return err
				}
			}
			return nil
		}
	}()
}
