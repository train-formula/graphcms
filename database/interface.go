package database

import (
	"github.com/train-formula/graphcms/database/cursor"
)

type TableModel interface {
	TableName() string
}

type CursorQueryModel interface {
	CursorQuery(prefix string, c cursor.Cursor) (string, []interface{}, error)
}
