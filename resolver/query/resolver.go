package query

import (
	"github.com/go-pg/pg/v9"
	"go.uber.org/zap"
)

type QueryResolver struct {
	db     *pg.DB
	logger *zap.Logger
}

func NewQueryResolver(db *pg.DB, logger *zap.Logger) *QueryResolver {
	return &QueryResolver{
		db:     db,
		logger: logger.Named("QueryResolver"),
	}
}
