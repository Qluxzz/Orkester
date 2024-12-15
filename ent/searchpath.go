// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"orkester/ent/searchpath"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// SearchPath is the model entity for the SearchPath schema.
type SearchPath struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Path holds the value of the "path" field.
	Path         string `json:"path,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SearchPath) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case searchpath.FieldID:
			values[i] = new(sql.NullInt64)
		case searchpath.FieldPath:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SearchPath fields.
func (sp *SearchPath) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case searchpath.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sp.ID = int(value.Int64)
		case searchpath.FieldPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field path", values[i])
			} else if value.Valid {
				sp.Path = value.String
			}
		default:
			sp.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the SearchPath.
// This includes values selected through modifiers, order, etc.
func (sp *SearchPath) Value(name string) (ent.Value, error) {
	return sp.selectValues.Get(name)
}

// Update returns a builder for updating this SearchPath.
// Note that you need to call SearchPath.Unwrap() before calling this method if this SearchPath
// was returned from a transaction, and the transaction was committed or rolled back.
func (sp *SearchPath) Update() *SearchPathUpdateOne {
	return NewSearchPathClient(sp.config).UpdateOne(sp)
}

// Unwrap unwraps the SearchPath entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sp *SearchPath) Unwrap() *SearchPath {
	_tx, ok := sp.config.driver.(*txDriver)
	if !ok {
		panic("ent: SearchPath is not a transactional entity")
	}
	sp.config.driver = _tx.drv
	return sp
}

// String implements the fmt.Stringer.
func (sp *SearchPath) String() string {
	var builder strings.Builder
	builder.WriteString("SearchPath(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sp.ID))
	builder.WriteString("path=")
	builder.WriteString(sp.Path)
	builder.WriteByte(')')
	return builder.String()
}

// SearchPaths is a parsable slice of SearchPath.
type SearchPaths []*SearchPath
