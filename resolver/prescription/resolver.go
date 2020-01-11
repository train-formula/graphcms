package prescription

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/train-formula/graphcms/calls/workoutcall"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
	"go.uber.org/zap"
)

type PrescriptionResolver struct {
	db     *pg.DB
	logger *zap.Logger
}

func NewPrescriptionResolver(db *pg.DB, logger *zap.Logger) *PrescriptionResolver {
	return &PrescriptionResolver{
		db:     db,
		logger: logger.Named("PrescriptionResolver"),
	}
}

func (r *PrescriptionResolver) HasReps(ctx context.Context, obj *workout.Prescription) (bool, error) {
	return true, nil
}

func (r *PrescriptionResolver) HasRepModifier(ctx context.Context, obj *workout.Prescription) (bool, error) {
	return true, nil
}

func (r *PrescriptionResolver) Sets(ctx context.Context, obj *workout.Prescription) ([]*workout.PrescriptionSet, error) {

	if obj == nil {
		return nil, gqlerror.Errorf("Cannot locate prescription sets from nil prescription")
	}

	call := workoutcall.NewGetPrescriptionPrescriptionSets(obj.ID, r.logger, r.db)

	if validation.ValidationChain(ctx, call.Validate(ctx)...) {
		return call.Call(ctx)
	}

	return nil, nil

}
