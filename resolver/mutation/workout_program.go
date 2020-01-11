package mutation

import (
	"context"

	"github.com/train-formula/graphcms/calls/workoutcall"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
)

func (m *MutationResolver) CreateWorkoutProgram(ctx context.Context, request generated.CreateWorkoutProgram) (*workout.WorkoutProgram, error) {

	call := workoutcall.NewCreateWorkoutProgram(request, m.logger, m.db)

	if validation.ValidationChain(ctx, call.Validate(ctx)...) {
		return call.Call(ctx)
	}

	return nil, nil
}
