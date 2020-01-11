package prescriptionset

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/train-formula/graphcms/calls/workoutcall"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
	"go.uber.org/zap"
)

type PrescriptionSetResolver struct {
	db     *pg.DB
	logger *zap.Logger
}

func NewPrescriptionSetResolver(db *pg.DB, logger *zap.Logger) *PrescriptionSetResolver {
	return &PrescriptionSetResolver{
		db:     db,
		logger: logger.Named("PrescriptionSetResolver"),
	}
}

func (r *PrescriptionSetResolver) RepUnit(ctx context.Context, obj *workout.PrescriptionSet) (*workout.Unit, error) {

	if obj == nil {
		return nil, gqlerror.Errorf("Cannot locate rep unit from nil prescription set")
	}

	if obj.RepUnitID == nil {
		return nil, nil
	}

	call := workoutcall.NewGetUnit(*obj.RepUnitID, r.logger, r.db)

	if validation.ValidationChain(ctx, call.Validate(ctx)...) {
		return call.Call(ctx)
	}

	return nil, nil

}

func (r *PrescriptionSetResolver) RepModifierUnit(ctx context.Context, obj *workout.PrescriptionSet) (*workout.Unit, error) {

	if obj == nil {
		return nil, gqlerror.Errorf("Cannot locate rep modifier unit from nil prescription set")
	}

	if obj.RepModifierUnitID == nil {
		return nil, nil
	}

	call := workoutcall.NewGetUnit(*obj.RepModifierUnitID, r.logger, r.db)

	if validation.ValidationChain(ctx, call.Validate(ctx)...) {
		return call.Call(ctx)
	}

	return nil, nil
}
