package workoutcall

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/train-formula/graphcms/database/workoutdb"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
	"go.uber.org/zap"
)

func NewSetWorkoutWorkoutCategories(request generated.SetWorkoutWorkoutCategories, logger *zap.Logger, db *pg.DB) *SetWorkoutWorkoutCategories {
	return &SetWorkoutWorkoutCategories{
		request: request,
		db:      db,
		logger:  logger.Named("SetWorkoutWorkoutCategories"),
	}
}

type SetWorkoutWorkoutCategories struct {
	request generated.SetWorkoutWorkoutCategories
	db      *pg.DB
	logger  *zap.Logger
}

func (c SetWorkoutWorkoutCategories) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (c SetWorkoutWorkoutCategories) Call(ctx context.Context) (*workout.Workout, error) {

	var finalWorkout *workout.Workout
	err := c.db.RunInTransaction(func(t *pg.Tx) error {

		var err error
		wrkout, err := workoutdb.GetWorkoutForUpdate(ctx, c.db, c.request.WorkoutID)
		if err != nil {
			if err == pg.ErrNoRows {
				return gqlerror.Errorf("Workout does not exist")
			}

			c.logger.Error("Error retrieving workout", zap.Error(err))
			return err
		}

		finalWorkout = &wrkout

		if len(c.request.WorkoutCategoryIDs) <= 0 {

			err = workoutdb.ClearWorkoutWorkoutCategories(ctx, t, wrkout.ID)
			if err != nil {
				c.logger.Error("Failed to clear workout workout categories", zap.Error(err))
				return err
			}

			return nil
		}

		err = workoutdb.SetWorkoutWorkoutCategories(ctx, t, wrkout.ID, c.request.WorkoutCategoryIDs)
		if err != nil {
			c.logger.Error("Failed to set workout workout categories", zap.Error(err))
			return err
		}

		return nil

	})

	if err != nil {
		return nil, err
	}

	return finalWorkout, nil
}
