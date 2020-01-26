package mutation

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/calls/plancall"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models/plan"
	"github.com/train-formula/graphcms/validation"
)

func (r *MutationResolver) CreatePlan(ctx context.Context, request generated.CreatePlan) (*plan.Plan, error) {

	call := plancall.NewCreatePlan(request, r.logger, r.db)

	if validation.ValidationChain(ctx, call.Validate(ctx)...) {
		return call.Call(ctx)
	}

	return nil, nil
}

func (r *MutationResolver) EditPlan(ctx context.Context, request generated.EditPlan) (*plan.Plan, error) {

	call := plancall.NewEditPlan(request, r.logger, r.db)

	if validation.ValidationChain(ctx, call.Validate(ctx)...) {
		return call.Call(ctx)
	}

	return nil, nil
}

func (r *MutationResolver) ArchivePlan(ctx context.Context, request uuid.UUID) (*plan.Plan, error) {

	call := plancall.NewArchivePlan(request, r.logger, r.db)

	if validation.ValidationChain(ctx, call.Validate(ctx)...) {
		return call.Call(ctx)
	}

	return nil, nil
}
