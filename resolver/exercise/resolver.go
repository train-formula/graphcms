package exercise

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/train-formula/graphcms/calls/workoutcall"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
	"go.uber.org/zap"
)

type ExerciseResolver struct {
	db     *pg.DB
	logger *zap.Logger
}

func NewExerciseResolver(db *pg.DB, logger *zap.Logger) *ExerciseResolver {
	return &ExerciseResolver{
		db:     db,
		logger: logger,
	}
}

func (r *ExerciseResolver) Prescription(ctx context.Context, obj *workout.Exercise) (*workout.Prescription, error) {

	if obj == nil {
		return nil, gqlerror.Errorf("Cannot locate prescription ID unit from nil exercise")
	}

	g := workoutcall.GetPrescription{
		ID:     obj.PrescriptionID,
		DB:     r.db,
		Logger: r.logger,
	}

	if validation.ValidationChain(ctx, g.Validate(ctx)...) {
		return g.Call(ctx)
	}

	return nil, nil
}
