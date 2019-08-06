package query

import "github.com/go-pg/pg/v9"

func NewQueryResolver(db *pg.DB) *QueryResolver {
	return &QueryResolver{
		db: db,
	}
}

type QueryResolver struct {
	db *pg.DB
}
