package workoutcall

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/dataloader/workoutid"
	"github.com/train-formula/graphcms/logging"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
	"go.uber.org/zap"
)

func NewGetWorkout(id uuid.UUID, logger *zap.Logger, db *pg.DB) *GetWorkout {
	return &GetWorkout{
		id:     id,
		db:     db,
		logger: logger.Named("GetWorkout"),
	}
}

type GetWorkout struct {
	id     uuid.UUID
	db     *pg.DB
	logger *zap.Logger
}

func (g GetWorkout) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (g GetWorkout) Call(ctx context.Context) (*workout.Workout, error) {

	loader := workoutid.GetContextLoader(ctx)

	loaded, err := loader.Load(g.id)
	if err != nil {
		g.logger.Error("Failed to load workout with dataloader", zap.Error(err),
			logging.UUID("workoutID", g.id))
		return nil, err
	}

	if loaded == nil {
		return nil, gqlerror.Errorf("Unknown workout id")
	}

	return loaded, nil
}
