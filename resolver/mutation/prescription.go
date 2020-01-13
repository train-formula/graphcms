package mutation

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/calls/workoutcall"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
)

func (r *MutationResolver) CreatePrescription(ctx context.Context, request generated.CreatePrescription) (*workout.Prescription, error) {

	call := workoutcall.NewCreatePrescription(request, r.logger, r.db)

	if validation.ValidationChain(ctx, call.Validate(ctx)...) {
		return call.Call(ctx)
	}

	return nil, nil
}

func (r *MutationResolver) DeletePrescription(ctx context.Context, request uuid.UUID) (*uuid.UUID, error) {
	call := workoutcall.NewDeletePrescription(request, r.logger, r.db)

	if validation.ValidationChain(ctx, call.Validate(ctx)...) {
		return call.Call(ctx)
	}

	return nil, nil
}

func (r *MutationResolver) EditPrescription(ctx context.Context, request generated.EditPrescription) (*workout.Prescription, error) {

	call := workoutcall.NewEditPrescription(request, r.logger, r.db)

	if validation.ValidationChain(ctx, call.Validate(ctx)...) {
		return call.Call(ctx)
	}

	return nil, nil

}
