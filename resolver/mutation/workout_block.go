package mutation

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/calls/workoutcall"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
)

func (m *MutationResolver) CreateWorkoutBlock(ctx context.Context, request generated.CreateWorkoutBlock) (*workout.WorkoutBlock, error) {

	call := workoutcall.NewCreateWorkoutBlock(request, m.logger, m.db)

	if validation.ValidationChain(ctx, call.Validate(ctx)...) {
		return call.Call(ctx)
	}

	return nil, nil
}

func (m *MutationResolver) EditWorkoutBlock(ctx context.Context, request generated.EditWorkoutBlock) (*workout.WorkoutBlock, error) {
	call := workoutcall.NewEditWorkoutBlock(request, m.logger, m.db)

	if validation.ValidationChain(ctx, call.Validate(ctx)...) {
		return call.Call(ctx)
	}

	return nil, nil
}

func (m *MutationResolver) DeleteWorkoutBlock(ctx context.Context, request uuid.UUID) (*uuid.UUID, error) {

	call := workoutcall.NewDeleteWorkoutBlock(request, m.logger, m.db)

	if validation.ValidationChain(ctx, call.Validate(ctx)...) {
		return call.Call(ctx)
	}

	return nil, nil
}

func (m *MutationResolver) SetWorkoutBlockExercises(ctx context.Context, request generated.SetWorkoutBlockExercises) (*workout.WorkoutBlock, error) {

	call := workoutcall.NewSetWorkoutBlockExercises(request, m.logger, m.db)

	if validation.ValidationChain(ctx, call.Validate(ctx)...) {
		return call.Call(ctx)
	}

	return nil, nil
}
