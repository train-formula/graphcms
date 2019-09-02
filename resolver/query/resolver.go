package query

import "github.com/go-pg/pg/v9"

type QueryResolver struct {
	db *pg.DB
}

func NewQueryResolver(db *pg.DB) *QueryResolver {
	return &QueryResolver{
		db: db,
	}
}
