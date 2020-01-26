package mutation

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/calls/plancall"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models/plan"
	"github.com/train-formula/graphcms/validation"
)

func (r *MutationResolver) CreatePlanSchedule(ctx context.Context, request generated.CreatePlanSchedule) (*plan.PlanSchedule, error) {

	call := plancall.NewCreatePlanSchedule(request, r.logger, r.db)

	if validation.ValidationChain(ctx, call.Validate(ctx)...) {
		return call.Call(ctx)
	}

	return nil, nil
}

func (r *MutationResolver) EditPlanSchedule(ctx context.Context, request generated.EditPlanSchedule) (*plan.PlanSchedule, error) {
	call := plancall.NewEditPlanSchedule(request, r.logger, r.db)

	if validation.ValidationChain(ctx, call.Validate(ctx)...) {
		return call.Call(ctx)
	}

	return nil, nil
}

func (r *MutationResolver) ArchivePlanSchedule(ctx context.Context, request uuid.UUID) (*plan.PlanSchedule, error) {
	call := plancall.NewArchivePlanSchedule(request, r.logger, r.db)

	if validation.ValidationChain(ctx, call.Validate(ctx)...) {
		return call.Call(ctx)
	}

	return nil, nil
}
