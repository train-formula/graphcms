package workoutcall

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/train-formula/graphcms/database/workoutdb"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/logging"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
	"go.uber.org/zap"
)

func NewEditWorkoutCategory(request generated.EditWorkoutCategory, logger *zap.Logger, db *pg.DB) *EditWorkoutCategory {
	return &EditWorkoutCategory{
		request: request,
		db:      db,
		logger:  logger.Named("EditWorkoutCategory"),
	}
}

type EditWorkoutCategory struct {
	request generated.EditWorkoutCategory
	db      *pg.DB
	logger  *zap.Logger
}

func (c EditWorkoutCategory) Validate(ctx context.Context) []validation.ValidatorFunc {

	return []validation.ValidatorFunc{
		validation.CheckStringNilOrIsNotEmpty(c.request.Name, "Name must not be empty"),
	}
}

func (c EditWorkoutCategory) Call(ctx context.Context) (*workout.WorkoutCategory, error) {

	var finalCategory *workout.WorkoutCategory

	err := c.db.RunInTransaction(func(t *pg.Tx) error {

		category, err := workoutdb.GetWorkoutCategoryForUpdate(ctx, t, c.request.ID)
		if err != nil {
			if err == pg.ErrNoRows {
				return gqlerror.Errorf("Workout category does not exist")
			}

			c.logger.Error("Error retrieving workout category", zap.Error(err),
				logging.UUID("workoutCategoryID", c.request.ID))
			return err
		}

		if c.request.Name != nil {
			category.Name = *c.request.Name
		}

		if c.request.Description != nil {
			category.Description = *c.request.Description
		}

		finalCategory, err = workoutdb.UpdateWorkoutCategory(ctx, t, category)
		if err != nil {
			c.logger.Error("Error updating workout category", zap.Error(err),
				logging.UUID("workoutCategoryID", c.request.ID))
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return finalCategory, nil

}
