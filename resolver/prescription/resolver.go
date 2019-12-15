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
	return obj.RepNumeral != nil || obj.RepText != nil, nil
}

func (r *PrescriptionResolver) HasRepModifier(ctx context.Context, obj *workout.Prescription) (bool, error) {
	return obj.RepModifierNumeral != nil || obj.RepModifierText != nil, nil
}

func (r *PrescriptionResolver) HasSets(ctx context.Context, obj *workout.Prescription) (bool, error) {
	return obj.SetNumeral != nil || obj.SetText != nil, nil
}

func (r *PrescriptionResolver) RepUnit(ctx context.Context, obj *workout.Prescription) (*workout.Unit, error) {

	if obj == nil {
		return nil, gqlerror.Errorf("Cannot locate rep unit from nil prescription")
	}

	if obj.RepUnitID == nil {
		return nil, nil
	}

	g := workoutcall.GetUnit{
		ID:     *obj.RepUnitID,
		DB:     r.db,
		Logger: r.logger,
	}

	if validation.ValidationChain(ctx, g.Validate(ctx)...) {
		return g.Call(ctx)
	}

	return nil, nil
}

func (r *PrescriptionResolver) RepModifierUnit(ctx context.Context, obj *workout.Prescription) (*workout.Unit, error) {
	if obj == nil {
		return nil, gqlerror.Errorf("Cannot locate rep unit from nil prescription")
	}

	if obj.RepModifierUnitID == nil {
		return nil, nil
	}

	g := workoutcall.GetUnit{
		ID:     *obj.RepModifierUnitID,
		DB:     r.db,
		Logger: r.logger,
	}

	if validation.ValidationChain(ctx, g.Validate(ctx)...) {
		return g.Call(ctx)
	}

	return nil, nil
}

func (r *PrescriptionResolver) SetUnit(ctx context.Context, obj *workout.Prescription) (*workout.Unit, error) {

	if obj == nil {
		return nil, gqlerror.Errorf("Cannot locate rep unit from nil prescription")
	}

	if obj.SetUnitID == nil {
		return nil, nil
	}

	g := workoutcall.GetUnit{
		ID:     *obj.SetUnitID,
		DB:     r.db,
		Logger: r.logger,
	}

	if validation.ValidationChain(ctx, g.Validate(ctx)...) {
		return g.Call(ctx)
	}

	return nil, nil

}
