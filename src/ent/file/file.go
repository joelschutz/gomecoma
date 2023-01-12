// Code generated by entc, DO NOT EDIT.

package file

import (
	"fmt"
)

const (
	// Label holds the string label denoting the file type in the database.
	Label = "file"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldPath holds the string denoting the path field in the database.
	FieldPath = "path"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldExternalID holds the string denoting the external_id field in the database.
	FieldExternalID = "external_id"
	// FieldExternalInfoProvider holds the string denoting the external_info_provider field in the database.
	FieldExternalInfoProvider = "external_info_provider"
	// FieldResults holds the string denoting the results field in the database.
	FieldResults = "results"
	// FieldSynced holds the string denoting the synced field in the database.
	FieldSynced = "synced"
	// EdgeMovie holds the string denoting the movie edge name in mutations.
	EdgeMovie = "movie"
	// Table holds the table name of the file in the database.
	Table = "files"
	// MovieTable is the table that holds the movie relation/edge.
	MovieTable = "files"
	// MovieInverseTable is the table name for the Movie entity.
	// It exists in this package in order to avoid circular dependency with the "movie" package.
	MovieInverseTable = "movies"
	// MovieColumn is the table column denoting the movie relation/edge.
	MovieColumn = "movie_file"
)

// Columns holds all SQL columns for file fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldPath,
	FieldType,
	FieldExternalID,
	FieldExternalInfoProvider,
	FieldResults,
	FieldSynced,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "files"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"movie_file",
}

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
	// DefaultSynced holds the default value on creation for the "synced" field.
	DefaultSynced bool
)

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeAudio Type = "audio"
	TypeVideo Type = "video"
	TypeImage Type = "image"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeAudio, TypeVideo, TypeImage:
		return nil
	default:
		return fmt.Errorf("file: invalid enum value for type field: %q", _type)
	}
}