package plancall

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database/plandb"
	"github.com/train-formula/graphcms/logging"
	"github.com/train-formula/graphcms/models/plan"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
	"go.uber.org/zap"
)

func NewArchivePlan(request uuid.UUID, logger *zap.Logger, db *pg.DB) *ArchivePlan {
	return &ArchivePlan{
		request: request,
		db:      db,
		logger:  logger.Named("ArchivePlan"),
	}
}

type ArchivePlan struct {
	request uuid.UUID
	db      *pg.DB
	logger  *zap.Logger
}

func (c ArchivePlan) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (c ArchivePlan) Call(ctx context.Context) (*plan.Plan, error) {

	var finalPlan *plan.Plan

	err := c.db.RunInTransaction(func(t *pg.Tx) error {

		plan, err := plandb.GetPlanForUpdate(ctx, t, c.request)
		if err != nil {
			if err == pg.ErrNoRows {
				return gqlerror.Errorf("Plan does not exist")
			}

			c.logger.Error("Error retrieving plan", zap.Error(err),
				logging.UUID("planID", c.request))
			return err
		}

		subscriberCount, err := plandb.CountPlanActiveSubscribers(ctx, t, c.request)
		if err != nil {
			c.logger.Error("Error counting active subscribers", zap.Error(err),
				logging.UUID("planID", c.request))
			return err
		}

		if subscriberCount > 0 {
			return gqlerror.Errorf("Cannot archive plan with active subscribers")
		}

		plan.Archived = true

		finalPlan, err = plandb.UpdatePlan(ctx, t, plan)
		if err != nil {
			c.logger.Error("Error archiving plan", zap.Error(err),
				logging.UUID("planID", c.request))
			return err
		}

		return nil

	})

	if err != nil {
		return nil, err
	}

	return finalPlan, nil

}
