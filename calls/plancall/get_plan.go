package plancall

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/dataloader/planid"
	"github.com/train-formula/graphcms/logging"
	"github.com/train-formula/graphcms/models/plan"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
	"go.uber.org/zap"
)

func NewGetPlan(id uuid.UUID, logger *zap.Logger, db *pg.DB) *GetPlan {
	return &GetPlan{
		id:     id,
		db:     db,
		logger: logger.Named("GetPlan"),
	}
}

type GetPlan struct {
	id     uuid.UUID
	db     *pg.DB
	logger *zap.Logger
}

func (g GetPlan) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (g GetPlan) Call(ctx context.Context) (*plan.Plan, error) {
	loader := planid.GetContextLoader(ctx)

	loaded, err := loader.Load(g.id)
	if err != nil {
		g.logger.Error("Failed to retrieve plan with dataloader", zap.Error(err), logging.UUID("planID", g.id))
		return nil, err
	}

	if loaded == nil {
		return nil, gqlerror.Errorf("Unknown plan id")
	}

	return loaded, nil
}
