package workoutcall

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/dataloader/workoutcategoriesbyworkout"
	"github.com/train-formula/graphcms/logging"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"go.uber.org/zap"
)

func NewGetWorkoutWorkoutCategories(workoutID uuid.UUID, logger *zap.Logger, db *pg.DB) *GetWorkoutWorkoutCategories {
	return &GetWorkoutWorkoutCategories{
		workoutID: workoutID,
		db:        db,
		logger:    logger.Named("GetWorkoutWorkoutCategories"),
	}
}

type GetWorkoutWorkoutCategories struct {
	workoutID uuid.UUID
	db        *pg.DB
	logger    *zap.Logger
}

func (g GetWorkoutWorkoutCategories) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (g GetWorkoutWorkoutCategories) Call(ctx context.Context) ([]*workout.WorkoutCategory, error) {

	loader := workoutcategoriesbyworkout.GetContextLoader(ctx)

	loaded, err := loader.Load(g.workoutID)
	if err != nil {
		g.logger.Error("Failed to load workout categories with dataloader", zap.Error(err),
			logging.UUID("workoutID", g.workoutID))
		return nil, err
	}

	return loaded, nil
}
