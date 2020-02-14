package plancall

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/dataloader/planscheduleid"
	"github.com/train-formula/graphcms/logging"
	"github.com/train-formula/graphcms/models/plan"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
	"github.com/willtrking/pgxload"
	"go.uber.org/zap"
)

// Retrieve individual schedule
func NewGetPlanSchedule(id uuid.UUID, logger *zap.Logger, db pgxload.PgxLoader) *GetPlanSchedule {
	return &GetPlanSchedule{
		id:     id,
		db:     db,
		logger: logger.Named("GetPlanSchedule"),
	}
}

type GetPlanSchedule struct {
	id     uuid.UUID
	db     pgxload.PgxLoader
	logger *zap.Logger
}

func (g GetPlanSchedule) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (g GetPlanSchedule) Call(ctx context.Context) (*plan.PlanSchedule, error) {
	loader := planscheduleid.GetContextLoader(ctx)

	loaded, err := loader.Load(g.id)
	if err != nil {
		g.logger.Error("Failed to retrieve plan schedule with dataloader", zap.Error(err), logging.UUID("planScheduleID", g.id))
		return nil, err
	}

	if loaded == nil {
		return nil, gqlerror.Errorf("Unknown plan schedule id")
	}

	return loaded, nil
}
