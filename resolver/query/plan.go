package query

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/calls/plancall"
	"github.com/train-formula/graphcms/database/cursor"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models/plan"
	"github.com/train-formula/graphcms/validation"
	"go.uber.org/zap"
)

func (r *QueryResolver) Plan(ctx context.Context, id uuid.UUID) (*plan.Plan, error) {

	g := plancall.NewGetPlan(id, r.logger, r.db)

	if validation.ValidationChain(ctx, g.Validate(ctx)...) {
		return g.Call(ctx)
	}

	return nil, nil
}

func (r *QueryResolver) PlanSchedule(ctx context.Context, id uuid.UUID) (*plan.PlanSchedule, error) {
	g := plancall.NewGetPlanSchedule(id, r.logger, r.db)

	if validation.ValidationChain(ctx, g.Validate(ctx)...) {
		return g.Call(ctx)
	}

	return nil, nil
}

func (r *QueryResolver) PlanSearch(ctx context.Context, request generated.PlanSearchRequest, first int, after *string) (*generated.PlanSearchResults, error) {
	curse, err := cursor.DeserializeCursor(after)
	if err != nil {
		r.logger.Error("Failed to deserialize cursor", zap.Error(err))
		return nil, err
	}

	s := plancall.NewSearchPlans(request, first, curse, r.logger, r.db)

	if validation.ValidationChain(ctx, s.Validate(ctx)...) {

		return s.Call(ctx)
	}

	return nil, nil
}
