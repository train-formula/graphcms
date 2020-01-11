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
		request: request,
		db:      db,
		kogger:  logger.Named("EditExercise"),
	}
}

type EditExercise struct {
	request generated.EditExercise
	db      *pg.DB
	kogger  *zap.Logger
}

func (c EditExercise) Validate(ctx context.Context) []validation.ValidatorFunc {

	return []validation.ValidatorFunc{
		validation.CheckStringNilOrIsNotEmpty(c.request.Name, "Name must not be empty"),
		func() *gqlerror.Error {
			if c.request.VideoURL != nil {
				return validation.CheckStringNilOrIsURL(c.request.VideoURL.Value, "Invalid video URL")()
			}

			return nil
		},
	}
}

func (c EditExercise) Call(ctx context.Context) (*workout.Exercise, error) {

	var finalExercise *workout.Exercise

	err := c.db.RunInTransaction(func(t *pg.Tx) error {

		exercise, err := workoutdb.GetExerciseForUpdate(ctx, c.db, c.request.ID)
		if err != nil {
			if err == pg.ErrNoRows {
				return gqlerror.Errorf("Exercise does not exist")
			}

			c.kogger.Error("Error retrieving exercise", zap.Error(err))
			return err
		}

		if c.request.Name != nil {
			exercise.Name = strings.TrimSpace(*c.request.Name)
		}

		if c.request.Description != nil {
			exercise.Description = strings.TrimSpace(*c.request.Description)
		}

		if c.request.VideoURL != nil {
			exercise.VideoURL = c.request.VideoURL.Value
		}

		finalExercise, err = workoutdb.UpdateExercise(ctx, c.db, exercise)
		if err != nil {
			c.kogger.Error("Error updating exercise", zap.Error(err))
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return finalExercise, nil

}
