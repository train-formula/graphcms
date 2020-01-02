package workoutblock

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/train-formula/graphcms/calls/organizationcall"
	"github.com/train-formula/graphcms/calls/workoutcall"
	"github.com/train-formula/graphcms/models/trainer"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
	"go.uber.org/zap"
)

type WorkoutBlockResolver struct {
	db     *pg.DB
	logger *zap.Logger
}

func NewWorkoutBlockResolver(db *pg.DB, logger *zap.Logger) *WorkoutBlockResolver {
	return &WorkoutBlockResolver{
		db:     db,
		logger: logger.Named("WorkoutBlockResolver"),
	}
}

func (r *WorkoutBlockResolver) TrainerOrganization(ctx context.Context, obj *workout.WorkoutBlock) (*trainer.Organization, error) {
	g := organizationcall.GetOrganization{
		ID: obj.TrainerOrganizationID,
		DB: r.db,
	}

	if validation.ValidationChain(ctx, g.Validate(ctx)...) {
		return g.Call(ctx)
	}

	return nil, nil
}

func (r *WorkoutBlockResolver) WorkoutCategory(ctx context.Context, obj *workout.WorkoutBlock) (*workout.WorkoutCategory, error) {

	if obj == nil {
		return nil, gqlerror.Errorf("Cannot locate workout category ID from nil workout block")
	}

	g := workoutcall.GetWorkoutCategory{
		ID:     obj.WorkoutCategoryID,
		DB:     r.db,
		Logger: r.logger,
	}

	if validation.ValidationChain(ctx, g.Validate(ctx)...) {
		return g.Call(ctx)
	}

	return nil, nil
}

func (r *WorkoutBlockResolver) Exercises(ctx context.Context, obj *workout.WorkoutBlock) ([]*workout.BlockExercise, error) {
	if obj == nil {
		return nil, gqlerror.Errorf("Cannot locate exercises from nil workout block")
	}

	call := workoutcall.NewGetBlockExercises(obj.ID, r.logger, r.db)

	if validation.ValidationChain(ctx, call.Validate(ctx)...) {
		return call.Call(ctx)
	}

	return nil, nil

}
