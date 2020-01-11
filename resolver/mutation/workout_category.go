package mutation

import (
	"context"

	"github.com/train-formula/graphcms/calls/workoutcall"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
)

func (m *MutationResolver) CreateWorkoutCategory(ctx context.Context, request generated.CreateWorkoutCategory) (*workout.WorkoutCategory, error) {
	call := workoutcall.CreateWorkoutCategory{
		Request: request,
		DB:      m.db,
	}

	if validation.ValidationChain(ctx, call.Validate(ctx)...) {
		return call.Call(ctx)
	}

	return nil, nil
}

func (m *MutationResolver) EditWorkoutCategory(ctx context.Context, request generated.EditWorkoutCategory) (*workout.WorkoutCategory, error) {

	call := workoutcall.NewEditWorkoutCategory(request, m.logger, m.db)

	if validation.ValidationChain(ctx, call.Validate(ctx)...) {
		return call.Call(ctx)
	}

	return nil, nil
}
