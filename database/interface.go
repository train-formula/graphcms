package database

import (
	"context"

	"github.com/go-pg/pg/v9"
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