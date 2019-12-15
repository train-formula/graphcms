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

type SetWorkoutWorkoutCategories struct {
	Request generated.SetWorkoutWorkoutCategories
	DB      *pg.DB
	Logger  *zap.Logger
}

func (c SetWorkoutWorkoutCategories) logger() *zap.Logger {

	return c.Logger.Named("SetWorkoutWorkoutCategories")

}

func (c SetWorkoutWorkoutCategories) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (c SetWorkoutWorkoutCategories) Call(ctx context.Context) (*workout.Workout, error) {

	var finalWorkout *workout.Workout
	err := c.DB.RunInTransaction(func(t *pg.Tx) error {

		var err error
		wrkout, err := workoutdb.GetWorkoutForUpdate(ctx, c.DB, c.Request.WorkoutID)
		if err != nil {
			if err == pg.ErrNoRows {
				return gqlerror.Errorf("Workout does not exist")
			}

			c.logger().Error("Error retrieving workout", zap.Error(err))
			return err
		}

		finalWorkout = &wrkout

		if len(c.Request.WorkoutCategoryIDs) <= 0 {
			return workoutdb.ClearWorkoutWorkoutCategories(ctx, t, wrkout.ID)
		}

		return workoutdb.SetWorkoutWorkoutCategories(ctx, t, wrkout.ID, c.Request.WorkoutCategoryIDs)

	})

	if err != nil {
		return nil, err
	}

	return finalWorkout, nil
}
