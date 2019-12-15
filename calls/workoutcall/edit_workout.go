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

type EditWorkout struct {
	Request generated.EditWorkout
	DB      *pg.DB
	Logger  *zap.Logger
}

func (c EditWorkout) logger() *zap.Logger {

	return c.Logger.Named("EditWorkout")

}

func (c EditWorkout) Validate(ctx context.Context) []validation.ValidatorFunc {

	return []validation.ValidatorFunc{
		validation.CheckStringNilOrIsNotEmpty(c.Request.Name, "Name must not be empty"),
	}
}

func (c EditWorkout) Call(ctx context.Context) (*workout.Workout, error) {

	var finalWorkout *workout.Workout

	err := c.DB.RunInTransaction(func(t *pg.Tx) error {

		wrkout, err := workoutdb.GetWorkoutForUpdate(ctx, c.DB, c.Request.ID)
		if err != nil {
			if err == pg.ErrNoRows {
				return gqlerror.Errorf("Workout does not exist")
			}

			c.logger().Error("Error retrieving workout", zap.Error(err))
			return err
		}

		if c.Request.Name != nil {
			wrkout.Name = *c.Request.Name
		}

		if c.Request.Description != nil {
			wrkout.Description = *c.Request.Description
		}

		finalWorkout, err = workoutdb.UpdateWorkout(ctx, c.DB, wrkout)
		if err != nil {
			c.logger().Error("Error updating workout", zap.Error(err))
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return finalWorkout, nil

}
