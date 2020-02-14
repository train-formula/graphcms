package unitdata

import (
	"context"

	"github.com/train-formula/graphcms/calls/workoutcall"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"github.com/willtrking/pgxload"
	"go.uber.org/zap"
)

type UnitDataResolver struct {
	db     pgxload.PgxLoader
	logger *zap.Logger
}

func NewUnitDataResolver(db pgxload.PgxLoader, logger *zap.Logger) *UnitDataResolver {
	return &UnitDataResolver{
		db:     db,
		logger: logger.Named("UnitDataResolver"),
	}
}

func (r *UnitDataResolver) Unit(ctx context.Context, obj *workout.UnitData) (*workout.Unit, error) {

	if obj == nil {
		return nil, nil
	}

	call := workoutcall.NewGetUnit(obj.UnitID, r.logger, r.db)

	if validation.ValidationChain(ctx, call.Validate(ctx)...) {
		return call.Call(ctx)
	}

	return nil, nil
}
