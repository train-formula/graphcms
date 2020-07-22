package workoutcall

import (
	"context"
	"strings"

	"github.com/jackc/pgx/v4"
	"github.com/train-formula/graphcms/database/types"
	"github.com/train-formula/graphcms/database/workoutdb"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/logging"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
	"github.com/willtrking/pgxload"
	"go.uber.org/zap"
)

func NewEditExercise(request generated.EditExercise, logger *zap.Logger, db pgxload.PgxLoader) *EditExercise {
	return &EditExercise{
		request: request,
		db:      db,
		kogger:  logger.Named("EditExercise"),
	}
}

type EditExercise struct {
	request generated.EditExercise
	db      pgxload.PgxLoader
	kogger  *zap.Logger
}

func (c EditExercise) Validate(ctx context.Context) []validation.ValidatorFunc {

	return []validation.ValidatorFunc{
		validation.CheckStringNilOrIsNotEmpty(c.request.Name, "Name must not be empty", true),
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

	err := pgxload.RunInTransaction(ctx, c.db, func(ctx context.Context, t pgxload.PgxTxLoader) error {

		exercise, err := workoutdb.GetExerciseForUpdate(ctx, t, c.request.ID)
		if err != nil {
			if err == pgx.ErrNoRows {
				return gqlerror.Errorf("Exercise does not exist")
			}

			c.kogger.Error("Error retrieving exercise", zap.Error(err),
				logging.UUID("exerciseID", c.request.ID))
			return err
		}

		if c.request.Name != nil {
			exercise.Name = strings.TrimSpace(*c.request.Name)
		}

		if c.request.Description != nil {
			exercise.Description = strings.TrimSpace(*c.request.Description)
		}

		if c.request.VideoURL != nil {
			exercise.VideoURL = types.ReadNullString(c.request.VideoURL.Value)
		}

		finalExercise, err = workoutdb.UpdateExercise(ctx, t, exercise)
		if err != nil {
			c.kogger.Error("Error updating exercise", zap.Error(err),
				logging.UUID("exerciseID", c.request.ID))
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return finalExercise, nil

}
