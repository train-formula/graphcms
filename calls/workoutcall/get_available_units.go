package workoutcall

import (
	"context"

	"github.com/train-formula/graphcms/database/workoutdb"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"github.com/willtrking/pgxload"
	"go.uber.org/zap"
)

func NewGetAvailableUnits(logger *zap.Logger, db pgxload.PgxLoader) *GetAvailableUnits {
	return &GetAvailableUnits{
		db:     db,
		logger: logger.Named("GetAvailableUnits"),
	}
}

type GetAvailableUnits struct {
	db     pgxload.PgxLoader
	logger *zap.Logger
}

func (g GetAvailableUnits) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (g GetAvailableUnits) Call(ctx context.Context) ([]*workout.Unit, error) {

	results, err := workoutdb.GetAllUnits(ctx, g.db)

	if err != nil {
		g.logger.Error("Failed to retrieve all available units", zap.Error(err))
		return nil, err
	}

	return results, nil
}
