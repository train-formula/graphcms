package workoutcall

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/dataloader/workoutblocksbycategory"
	"github.com/train-formula/graphcms/logging"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"github.com/willtrking/pgxload"
	"go.uber.org/zap"
)

func NewGetWorkoutCategoryBlocks(workoutCategoryID uuid.UUID, logger *zap.Logger, db pgxload.PgxLoader) *GetWorkoutCategoryBlocks {
	return &GetWorkoutCategoryBlocks{
		workoutCategoryID: workoutCategoryID,
		db:                db,
		logger:            logger.Named("GetWorkoutCategoryBlocks"),
	}
}

type GetWorkoutCategoryBlocks struct {
	workoutCategoryID uuid.UUID
	db                pgxload.PgxLoader
	logger            *zap.Logger
}

func (g GetWorkoutCategoryBlocks) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (g GetWorkoutCategoryBlocks) Call(ctx context.Context) ([]*workout.WorkoutBlock, error) {

	loader := workoutblocksbycategory.GetContextLoader(ctx)

	loaded, err := loader.Load(g.workoutCategoryID)
	if err != nil {
		g.logger.Error("Failed to load workout category blocks with dataloader", zap.Error(err),
			logging.UUID("workoutCategoryID", g.workoutCategoryID))
		return nil, err
	}

	return loaded, nil
}
