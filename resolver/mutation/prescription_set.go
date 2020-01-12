package mutation

import (
	"context"

	"github.com/train-formula/graphcms/calls/workoutcall"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
)

func (r *MutationResolver) CreatePrescriptionSet(ctx context.Context, request generated.CreatePrescriptionSet) (*workout.PrescriptionSet, error) {

	call := workoutcall.NewCreatePrescriptionSet(request, r.logger, r.db)

	if validation.ValidationChain(ctx, call.Validate(ctx)...) {
		return call.Call(ctx)
	}

	return nil, nil
}
