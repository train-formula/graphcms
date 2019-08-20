package database

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/train-formula/graphcms/database/cursor"
)

type Conn interface {
	QueryContext(
		c context.Context,
		model interface{},
		query interface{},
		params ...interface{},
	) (pg.Result, error)

	QueryOneContext(
		c context.Context,
		model interface{},
		query interface{},
		params ...interface{},
	) (pg.Result, error)
}

type TableModel interface {
	TableName() string
}

type CursorQueryModel interface {
	CursorQuery(prefix string, c cursor.Cursor) (string, []interface{}, error)
}
