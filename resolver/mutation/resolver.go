package mutation

import (
	"github.com/go-pg/pg/v9"
	"go.uber.org/zap"
)

func NewMutationResolver(db *pg.DB, logger *zap.Logger) *MutationResolver {
	return &MutationResolver{
		db:     db,
		logger: logger.Named("MutationResolver"),
	}
}

type MutationResolver struct {
	db     *pg.DB
	logger *zap.Logger
}
