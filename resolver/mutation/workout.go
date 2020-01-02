package mutation

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/calls/workoutcall"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
)

func (m *MutationResolver) CreateWorkout(ctx context.Context, request generated.CreateWorkout) (*workout.Workout, error) {

	call := workoutcall.CreateWorkout{
		Request: request,
		DB:      m.db,
	}

	if validation.ValidationChain(ctx, call.Validate(ctx)...) {
		return call.Call(ctx)
	}

	return nil, nil

}

func (m *MutationResolver) EditWorkout(ctx context.Context, request generated.EditWorkout) (*workout.Workout, error) {

	call := workoutcall.EditWorkout{
		Logger:  m.logger,
		Request: request,
		DB:      m.db,
	}

	if validation.ValidationChain(ctx, call.Validate(ctx)...) {
		return call.Call(ctx)
	}

	return nil, nil
}

func (m *MutationResolver) DeleteWorkout(ctx context.Context, request uuid.UUID) (*uuid.UUID, error) {

	call := workoutcall.DeleteWorkout{
		Logger:  m.logger,
		Request: request,
		DB:      m.db,
	}

	if validation.ValidationChain(ctx, call.Validate(ctx)...) {
		return call.Call(ctx)
	}

	return nil, nil
}

func (m *MutationResolver) SetWorkoutWorkoutCategories(ctx context.Context, request generated.SetWorkoutWorkoutCategories) (*workout.Workout, error) {

	call := workoutcall.NewSetWorkoutWorkoutCategories(request, m.logger, m.db)

	if validation.ValidationChain(ctx, call.Validate(ctx)...) {
		return call.Call(ctx)
	}

	return nil, nil
}
