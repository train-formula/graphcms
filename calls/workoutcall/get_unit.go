package workoutcall

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/dataloader/unitid"
	"github.com/train-formula/graphcms/logging"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
	"go.uber.org/zap"
)

func NewGetUnit(id uuid.UUID, logger *zap.Logger, db *pg.DB) *GetUnit {

	return &GetUnit{
		id:     id,
		db:     db,
		logger: logger.Named("GetUnit"),
	}
}

type GetUnit struct {
	id     uuid.UUID
	db     *pg.DB
	logger *zap.Logger
}

func (g GetUnit) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (g GetUnit) Call(ctx context.Context) (*workout.Unit, error) {

	loader := unitid.GetContextLoader(ctx)

	loaded, err := loader.Load(g.id)
	if err != nil {
		g.logger.Error("Failed to load unit with dataloader", zap.Error(err),
			logging.UUID("unitID", g.id))
		return nil, err
	}

	if loaded == nil {
		return nil, gqlerror.Errorf("Unknown unit id")
	}

	return loaded, nil
}
