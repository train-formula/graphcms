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

type EditWorkoutCategory struct {
	Request generated.EditWorkoutCategory
	DB      *pg.DB
	Logger  *zap.Logger
}

func (c EditWorkoutCategory) logger() *zap.Logger {

	return c.Logger.Named("EditWorkoutCategory")

}

func (c EditWorkoutCategory) Validate(ctx context.Context) []validation.ValidatorFunc {

	return []validation.ValidatorFunc{
		validation.CheckStringNilOrIsNotEmpty(c.Request.Name, "Name must not be empty"),
	}
}

func (c EditWorkoutCategory) Call(ctx context.Context) (*workout.WorkoutCategory, error) {

	var finalCategory *workout.WorkoutCategory

	err := c.DB.RunInTransaction(func(t *pg.Tx) error {

		category, err := workoutdb.GetWorkoutCategoryForUpdate(ctx, t, c.Request.ID)
		if err != nil {
			if err == pg.ErrNoRows {
				return gqlerror.Errorf("Workout category does not exist")
			}

			c.logger().Error("Error retrieving workout category", zap.Error(err))
			return err
		}

		if c.Request.Name != nil {
			category.Name = *c.Request.Name
		}

		if c.Request.Description != nil {
			category.Description = *c.Request.Description
		}

		finalCategory, err = workoutdb.UpdateWorkoutCategory(ctx, t, category)
		if err != nil {
			c.logger().Error("Error updating workout category", zap.Error(err))
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return finalCategory, nil

}
