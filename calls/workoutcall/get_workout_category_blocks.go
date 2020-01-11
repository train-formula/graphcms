package workoutcall

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/dataloader/workoutblocksbycategory"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"go.uber.org/zap"
)

func NewGetWorkoutCategoryBlocks(id uuid.UUID, logger *zap.Logger, db *pg.DB) *GetWorkoutCategoryBlocks {
	return &GetWorkoutCategoryBlocks{
		id:     id,
		db:     db,
		logger: logger.Named("GetWorkoutCategoryBlocks"),
	}
}

type GetWorkoutCategoryBlocks struct {
	id     uuid.UUID
	db     *pg.DB
	logger *zap.Logger
}

func (g GetWorkoutCategoryBlocks) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (g GetWorkoutCategoryBlocks) Call(ctx context.Context) ([]*workout.WorkoutBlock, error) {

	loader := workoutblocksbycategory.GetContextLoader(ctx)

	loaded, err := loader.Load(g.id)
	if err != nil {
		g.logger.Error("Failed to load workout category blocks with dataloader", zap.Error(err))
		return nil, err
	}

	return loaded, nil
}
