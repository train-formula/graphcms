package planschedule

import (
	"context"

	"github.com/train-formula/graphcms/calls/plancall"
	"github.com/train-formula/graphcms/logging"
	"github.com/train-formula/graphcms/models/interval"
	"github.com/train-formula/graphcms/models/plan"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
	"github.com/willtrking/pgxload"
	"go.uber.org/zap"
)

type PlanScheduleResolver struct {
	db     pgxload.PgxLoader
	logger *zap.Logger
}

func NewPlanScheduleResolver(db pgxload.PgxLoader, logger *zap.Logger) *PlanScheduleResolver {
	return &PlanScheduleResolver{
		db:     db,
		logger: logger.Named("PlanScheduleResolver"),
	}
}

func (r *PlanScheduleResolver) PaymentInterval(ctx context.Context, obj *plan.PlanSchedule) (*interval.DiurnalInterval, error) {

	if obj == nil {
		return nil, nil
	}

	if obj.PaymentInterval == interval.UnknownDiurnalInterval {
		r.logger.Error("Plan schedule malformed, payment interval of unknown type", logging.UUID("planSchedule", obj.ID))
		return nil, gqlerror.Errorf("Plan schedule malformed, payment interval of unknown type")
	}

	return &interval.DiurnalInterval{
		Interval: obj.PaymentInterval,
		Count:    int64(obj.PaymentIntervalCount),
	}, nil
}

func (r *PlanScheduleResolver) DurationInterval(ctx context.Context, obj *plan.PlanSchedule) (*interval.DiurnalInterval, error) {

	if obj == nil {
		return nil, nil
	}

	if obj.DurationInterval == nil && !obj.DurationIntervalCount.Valid {
		return nil, nil
	} else if obj.DurationInterval == nil && obj.DurationIntervalCount.Valid {

		r.logger.Error("Plan schedule malformed, duration interval has count but no type", logging.UUID("planSchedule", obj.ID))
		return nil, gqlerror.Errorf("Plan schedule malformed, duration interval has count but no type")
	} else if obj.DurationInterval != nil && !obj.DurationIntervalCount.Valid {

		r.logger.Error("Plan schedule malformed, duration interval has type but no count", logging.UUID("planSchedule", obj.ID))
		return nil, gqlerror.Errorf("Plan schedule malformed, duration interval has type but no count")
	} else if obj.DurationInterval != nil && *obj.DurationInterval == interval.UnknownDiurnalInterval {
		r.logger.Error("Plan schedule malformed, duration interval of unknown type", logging.UUID("planSchedule", obj.ID))
		return nil, gqlerror.Errorf("Plan schedule malformed, duration interval of unknown type")
	}

	return &interval.DiurnalInterval{
		Interval: *obj.DurationInterval,
		Count:    obj.DurationIntervalCount.Int64,
	}, nil
}

func (r *PlanScheduleResolver) Plan(ctx context.Context, obj *plan.PlanSchedule) (*plan.Plan, error) {

	if obj == nil {
		return nil, nil
	}

	g := plancall.NewGetPlan(obj.PlanID, r.logger, r.db)

	if validation.ValidationChain(ctx, g.Validate(ctx)...) {
		return g.Call(ctx)
	}

	return nil, nil

}
