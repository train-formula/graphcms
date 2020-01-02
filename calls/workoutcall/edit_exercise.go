package workoutcall

import (
	"context"
	"strings"

	"github.com/go-pg/pg/v9"
	"github.com/train-formula/graphcms/database/workoutdb"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
	"go.uber.org/zap"
)

func NewEditExercise(request generated.EditExercise, logger *zap.Logger, db *pg.DB) *EditExercise {
	return &EditExercise{
		Request: request,
		DB:      db,
		Logger:  logger.Named("EditExercise"),
	}
}

type EditExercise struct {
	Request generated.EditExercise
	DB      *pg.DB
	Logger  *zap.Logger
}

func (c EditExercise) Validate(ctx context.Context) []validation.ValidatorFunc {

	return []validation.ValidatorFunc{
		validation.CheckStringNilOrIsNotEmpty(c.Request.Name, "Name must not be empty"),
		func() *gqlerror.Error {
			if c.Request.VideoURL != nil {
				return validation.CheckStringNilOrIsURL(c.Request.VideoURL.Value, "Invalid video URL")()
			}

			return nil
		},
	}
}

func (c EditExercise) Call(ctx context.Context) (*workout.Exercise, error) {

	var finalExercise *workout.Exercise

	err := c.DB.RunInTransaction(func(t *pg.Tx) error {

		exercise, err := workoutdb.GetExerciseForUpdate(ctx, c.DB, c.Request.ID)
		if err != nil {
			if err == pg.ErrNoRows {
				return gqlerror.Errorf("Exercise does not exist")
			}

			c.Logger.Error("Error retrieving exercise", zap.Error(err))
			return err
		}

		if c.Request.Name != nil {
			exercise.Name = strings.TrimSpace(*c.Request.Name)
		}

		if c.Request.Description != nil {
			exercise.Description = strings.TrimSpace(*c.Request.Description)
		}

		if c.Request.VideoURL != nil {
			exercise.VideoURL = c.Request.VideoURL.Value
		}

		finalExercise, err = workoutdb.UpdateExercise(ctx, c.DB, exercise)
		if err != nil {
			c.Logger.Error("Error updating exercise", zap.Error(err))
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return finalExercise, nil

}
