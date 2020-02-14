package mutation

import (
	"github.com/willtrking/pgxload"
	"go.uber.org/zap"
)

func NewMutationResolver(db pgxload.PgxLoader, logger *zap.Logger) *MutationResolver {
	return &MutationResolver{
		db:     db,
		logger: logger.Named("MutationResolver"),
	}
}

type MutationResolver struct {
	db     pgxload.PgxLoader
	logger *zap.Logger
}
