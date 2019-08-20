package database

import (
	"fmt"
	"strconv"

	"github.com/train-formula/graphcms/database/cursor"
)

func TableName(m TableModel) string {
	return m.TableName()
}

func CursorQuery(prefix string, c cursor.Cursor, m CursorQueryModel) (string, []interface{}, error) {
	return m.CursorQuery(prefix, c)
}

func BasicCursorfyQuery(query string, prefix string, after cursor.Cursor, m CursorQueryModel, limit int, params ...interface{}) (string, []interface{}, error) {

	if after != nil {
		fmt.Println("as")
		fmt.Println(after == nil)
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
