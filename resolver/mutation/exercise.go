package mutation

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/calls/workoutcall"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
)

func (r *MutationResolver) CreateExercise(ctx context.Context, request generated.CreateExercise) (*workout.Exercise, error) {

	call := workoutcall.NewCreateExercise(request, r.logger, r.db)

	if validation.ValidationChain(ctx, call.Validate(ctx)...) {
		return call.Call(ctx)
	}

	return nil, nil
}

func (r *MutationResolver) EditExercise(ctx context.Context, request generated.EditExercise) (*workout.Exercise, error) {
	call := workoutcall.NewEditExercise(request, r.logger, r.db)

	if validation.ValidationChain(ctx, call.Validate(ctx)...) {
		return call.Call(ctx)
	}

	return nil, nil
}

func (r *MutationResolver) DeleteExercise(ctx context.Context, request uuid.UUID) (*uuid.UUID, error) {

	call := workoutcall.NewDeleteExercise(request, r.logger, r.db)

	if validation.ValidationChain(ctx, call.Validate(ctx)...) {
		return call.Call(ctx)
	}

	return nil, nil
}
