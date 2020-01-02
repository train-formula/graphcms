package database

import (
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

func ReflectValue(s interface{}) reflect.Value {
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

func StructColumns(v reflect.Value, prefix string) []string {

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	t := v.Type()

	var results []string

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if unicode.IsUpper(rune(field.Name[0])) {
			results = append(results, PGPrefixedColumn(strcase.ToSnake(field.Name), prefix))
		}

	}

	return results

}
