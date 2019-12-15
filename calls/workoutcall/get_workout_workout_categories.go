package workoutcall

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/dataloader/workoutcategoriesbyworkout"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"go.uber.org/zap"
)

type GetWorkoutWorkoutCategories struct {
	ID     uuid.UUID
	DB     *pg.DB
	Logger *zap.Logger
}

func (g GetWorkoutWorkoutCategories) logger() *zap.Logger {

	return g.Logger.Named("GetWorkoutWorkoutCategories")

}

func (g GetWorkoutWorkoutCategories) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (g GetWorkoutWorkoutCategories) Call(ctx context.Context) ([]*workout.WorkoutCategory, error) {

	loader := workoutcategoriesbyworkout.GetContextLoader(ctx)

	loaded, err := loader.Load(g.ID)
	if err != nil {
		g.logger().Error("Failed to load workout categories with dataloader", zap.Error(err))
		return nil, err
	}

	return loaded, nil
}
