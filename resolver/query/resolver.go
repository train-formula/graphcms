package query

import (
	"github.com/willtrking/pgxload"
	"go.uber.org/zap"
)

type QueryResolver struct {
	db     pgxload.PgxLoader
	logger *zap.Logger
}

func NewQueryResolver(db pgxload.PgxLoader, logger *zap.Logger) *QueryResolver {
	return &QueryResolver{
		db:     db,
		logger: logger.Named("QueryResolver"),
	}
}
