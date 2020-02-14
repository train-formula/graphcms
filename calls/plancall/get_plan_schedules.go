package plancall

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/dataloader/planschedulesbyplan"
	"github.com/train-formula/graphcms/logging"
	"github.com/train-formula/graphcms/models/plan"
	"github.com/train-formula/graphcms/validation"
	"github.com/willtrking/pgxload"
	"go.uber.org/zap"
)

// Retrieve schedules by plan ID
func NewGetPlanSchedules(planID uuid.UUID, logger *zap.Logger, db pgxload.PgxLoader) *GetPlanSchedules {
	return &GetPlanSchedules{
		planID: planID,
		db:     db,
		logger: logger.Named("GetPlanSchedules"),
	}
}

type GetPlanSchedules struct {
	planID uuid.UUID
	db     pgxload.PgxLoader
	logger *zap.Logger
}

func (g GetPlanSchedules) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (g GetPlanSchedules) Call(ctx context.Context) ([]*plan.PlanSchedule, error) {
	loader := planschedulesbyplan.GetContextLoader(ctx)

	loaded, err := loader.Load(g.planID)
	if err != nil {
		g.logger.Error("Failed to retrieve plan schedules for plan with dataloader", zap.Error(err), logging.UUID("planID", g.planID))
		return nil, err
	}

	return loaded, nil
}
