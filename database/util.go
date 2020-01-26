package database

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"unicode"

	"github.com/iancoleman/strcase"
	"github.com/train-formula/graphcms/database/cursor"
)

func TableName(m TableModel) string {
	return m.TableName()
}

func CursorQuery(prefix string, c cursor.Cursor, m CursorQueryModel) (string, []interface{}, error) {
	return m.CursorQuery(prefix, c)
}

func PGColumn(column string) string {

	if column[0] != '"' {
		column = fmt.Sprintf("\"%s", column)
	}

	if column[len(column)-1] != '"' {
		column = fmt.Sprintf("%s\"", column)
	}

	return column
}

func PGPrefixedColumn(column, prefix string) string {

	column = PGColumn(column)

	if prefix == "" {
		return column
	}

	prefix = PGColumn(prefix)

	return fmt.Sprintf("%s.%s", prefix, column)

}

func BasicCursorPaginationQuery(query string, prefix string, after cursor.Cursor, m CursorQueryModel, limit int, params ...interface{}) (string, []interface{}, error) {

	if after != nil {

		cursorQuery, cursorParams, err := CursorQuery(prefix, after, m)
		if err != nil {
			return "", nil, err
		}
		query += " AND " + cursorQuery

		params = append(params, cursorParams...)
	}

	query += " LIMIT " + strconv.Itoa(limit)

	return query, params, nil
}

// Reflects a struct value and resolves pointers. Panics if input is not a struct
func ReflectStructValue(s interface{}) reflect.Value {
	v := reflect.ValueOf(s)

	// if pointer get the underlying element≤
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		panic("not struct")
	}

	return v
}

// Get a list of columns from the fields on the specified struct, with the specified prefix added to column names
func StructColumns(s interface{}, prefix string) []string {

	v := ReflectStructValue(s)

	t := v.Type()

	var results []string

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if unicode.IsUpper(rune(field.Name[0])) && !shouldIgnoreColumn(field.Name) {
			results = append(results, PGPrefixedColumn(strcase.ToSnake(field.Name), prefix))
		}

	}

	return results

}

// Return value for StructColumnValues
type StructColumnValue struct {
	Column string
	Value  interface{}
}

// Get a list of columns + values from the fields in the specified struct, with the specified prefix added to column names
func StructColumnValues(s interface{}, prefix string) []StructColumnValue {

	var results []StructColumnValue

	v := ReflectStructValue(s)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fieldVal := v.Field(i)
		field := t.Field(i)

		if unicode.IsUpper(rune(field.Name[0])) && !shouldIgnoreColumn(field.Name) {
			results = append(results, StructColumnValue{
				Column: PGPrefixedColumn(strcase.ToSnake(field.Name), prefix),
				Value:  fieldVal.Interface(),
			})
		}
	}

	return results
}

// Generate a basic insert statement (INSERT INTO ... (...) VALUES (...) from the specified struct
// Also returns necessary list of params
// Does NOT include a ; at the end of the query to allow for easy extension (e.g. adding a ON CONFLICT)
func StructInsertStatement(s TableModel, prefix string) (string, []interface{}, error) {

	columnValues := StructColumnValues(s, prefix)

	if len(columnValues) <= 0 {
		return "", nil, errors.New("cannot generate insert statement for struct without columns")
	}

	cols := ""
	vals := ""
	var params []interface{}

	for idx, c := range columnValues {

		cols += c.Column
		vals += "?"
		params = append(params, c.Value)

		if idx < len(columnValues)-1 {
			cols += ", "
			vals += ", "
		}
	}

	return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", TableName(s), cols, vals), params, nil
}

// Columns to ignore when extracting columns from structs programmatically
var columnsToIgnore = map[string]struct{}{
	"CreatedAt": {},
	"UpdatedAt": {},
}

func shouldIgnoreColumn(col string) bool {
	_, ignore := columnsToIgnore[col]
	return ignore
}
