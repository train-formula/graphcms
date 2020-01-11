package workoutcall

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/dataloader/workoutblockid"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
	"go.uber.org/zap"
)

func NewGetWorkoutBlock(id uuid.UUID, logger *zap.Logger, db *pg.DB) *GetWorkoutBlock {
	return &GetWorkoutBlock{
		id:     id,
		db:     db,
		logger: logger.Named("GetWorkoutBlock"),
	}
}

type GetWorkoutBlock struct {
	id     uuid.UUID
	db     *pg.DB
	logger *zap.Logger
}

func (g GetWorkoutBlock) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (g GetWorkoutBlock) Call(ctx context.Context) (*workout.WorkoutBlock, error) {

	loader := workoutblockid.GetContextLoader(ctx)

	loaded, err := loader.Load(g.id)
	if err != nil {
		g.logger.Error("Failed to load workout block with dataloader", zap.Error(err))
		return nil, err
	}

	if loaded == nil {
		return nil, gqlerror.Errorf("Unknown workout block ID")
	}

	return loaded, nil
}
