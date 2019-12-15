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

type GetWorkoutCategoryBlocks struct {
	ID     uuid.UUID
	DB     *pg.DB
	Logger *zap.Logger
}

func (g GetWorkoutCategoryBlocks) logger() *zap.Logger {

	return g.Logger.Named("GetWorkoutCategoryBlocks")

}

func (g GetWorkoutCategoryBlocks) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (g GetWorkoutCategoryBlocks) Call(ctx context.Context) ([]*workout.WorkoutBlock, error) {

	loader := workoutblocksbycategory.GetContextLoader(ctx)

	loaded, err := loader.Load(g.ID)
	if err != nil {
		g.logger().Error("Failed to load workout category blocks with dataloader", zap.Error(err))
		return nil, err
	}

	return loaded, nil
}
