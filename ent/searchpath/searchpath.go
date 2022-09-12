// Code generated by ent, DO NOT EDIT.

package searchpath

const (
	// Label holds the string label denoting the searchpath type in the database.
	Label = "search_path"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPath holds the string denoting the path field in the database.
	FieldPath = "path"
	// Table holds the table name of the searchpath in the database.
	Table = "search_paths"
)

// Columns holds all SQL columns for searchpath fields.
var Columns = []string{
	FieldID,
	FieldPath,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}