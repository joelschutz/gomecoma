// Code generated by entc, DO NOT EDIT.

package picture

const (
	// Label holds the string label denoting the picture type in the database.
	Label = "picture"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldFilename holds the string denoting the filename field in the database.
	FieldFilename = "filename"
	// FieldPath holds the string denoting the path field in the database.
	FieldPath = "path"
	// Table holds the table name of the picture in the database.
	Table = "pictures"
)

// Columns holds all SQL columns for picture fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldFilename,
	FieldPath,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "pictures"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"artist_pictures",
	"movie_fanart",
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