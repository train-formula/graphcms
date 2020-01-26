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

func NewArchivePlanSchedule(request uuid.UUID, logger *zap.Logger, db *pg.DB) *ArchivePlanSchedule {
	return &ArchivePlanSchedule{
		request: request,
		db:      db,
		logger:  logger.Named("ArchivePlanSchedule"),
	}
}

type ArchivePlanSchedule struct {
	request uuid.UUID
	db      *pg.DB
	logger  *zap.Logger
}

func (c ArchivePlanSchedule) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (c ArchivePlanSchedule) Call(ctx context.Context) (*plan.PlanSchedule, error) {

	var finalPlanSchedule *plan.PlanSchedule

	err := c.db.RunInTransaction(func(t *pg.Tx) error {

		plan, err := plandb.GetPlanScheduleForUpdate(ctx, t, c.request)
		if err != nil {
			if err == pg.ErrNoRows {
				return gqlerror.Errorf("Plan schedule does not exist")
			}

			c.logger.Error("Error retrieving plan schedule", zap.Error(err),
				logging.UUID("planScheduleID", c.request))
			return err
		}

		subscriberCount, err := plandb.CountPlanScheduleActiveSubscribers(ctx, t, c.request)
		if err != nil {
			c.logger.Error("Error counting active subscribers", zap.Error(err),
				logging.UUID("planScheduleID", c.request))
			return err
		}

		if subscriberCount > 0 {
			return gqlerror.Errorf("Cannot archive plan schedule with active subscribers")
		}

		plan.Archived = true

		finalPlanSchedule, err = plandb.UpdatePlanSchedule(ctx, t, plan)
		if err != nil {
			c.logger.Error("Error archiving plan schedule", zap.Error(err),
				logging.UUID("planScheduleID", c.request))
			return err
		}

		return nil

	})

	if err != nil {
		return nil, err
	}

	return finalPlanSchedule, nil

}
