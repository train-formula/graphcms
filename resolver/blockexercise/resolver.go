package blockexercise

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/train-formula/graphcms/calls/workoutcall"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
	"go.uber.org/zap"
)

type BlockExerciseResolver struct {
	db     *pg.DB
	logger *zap.Logger
}

func NewBlockExerciseResolver(db *pg.DB, logger *zap.Logger) *BlockExerciseResolver {
	return &BlockExerciseResolver{
		db:     db,
		logger: logger,
	}
}

func (r *BlockExerciseResolver) Exercise(ctx context.Context, obj *workout.BlockExercise) (*workout.Exercise, error) {

	if obj == nil {
		return nil, gqlerror.Errorf("Cannot locate exercise from nil block exercise")
	}

	call := workoutcall.NewGetExercise(obj.ExerciseID, r.logger, r.db)

	if validation.ValidationChain(ctx, call.Validate(ctx)...) {
		return call.Call(ctx)
	}

	return nil, nil
}

func (r *BlockExerciseResolver) Prescription(ctx context.Context, obj *workout.BlockExercise) (*workout.Prescription, error) {

	if obj == nil {
		return nil, gqlerror.Errorf("Cannot locate prescription from nil block exercise")
	}

	call := workoutcall.NewGetPrescription(obj.PrescriptionID, r.logger, r.db)

	if validation.ValidationChain(ctx, call.Validate(ctx)...) {
		return call.Call(ctx)
	}

	return nil, nil
}
