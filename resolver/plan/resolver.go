package plan

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/train-formula/graphcms/calls/plancall"
	"github.com/train-formula/graphcms/calls/tagcall"
	"github.com/train-formula/graphcms/database/tagdb"
	"github.com/train-formula/graphcms/models/plan"
	"github.com/train-formula/graphcms/models/tag"
	"github.com/train-formula/graphcms/validation"
	"go.uber.org/zap"
)

type PlanResolver struct {
	db     *pg.DB
	logger *zap.Logger
}

func NewPlanResolver(db *pg.DB, logger *zap.Logger) *PlanResolver {
	return &PlanResolver{
		db:     db,
		logger: logger.Named("PlanResolver"),
	}
}

func (r *PlanResolver) Schedules(ctx context.Context, obj *plan.Plan) ([]*plan.PlanSchedule, error) {

	if obj == nil {
		return nil, nil
	}

	g := plancall.NewGetPlanSchedules(obj.ID, r.logger, r.db)

	if validation.ValidationChain(ctx, g.Validate(ctx)...) {
		return g.Call(ctx)
	}

	return nil, nil

}

func (r *PlanResolver) Tags(ctx context.Context, obj *plan.Plan) ([]*tag.Tag, error) {

	if obj == nil {
		return nil, nil
	}

	request := tagdb.TagsByObject{
		ObjectUUID: obj.ID,
		ObjectType: tag.PlanTagType,
	}

	g := tagcall.NewGetObjectTags(request, r.logger, r.db)

	if validation.ValidationChain(ctx, g.Validate(ctx)...) {
		return g.Call(ctx)
	}

	return nil, nil
}
