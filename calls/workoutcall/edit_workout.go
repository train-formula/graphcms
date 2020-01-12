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

func NewEditWorkout(request generated.EditWorkout, logger *zap.Logger, db *pg.DB) *EditWorkout {
	return &EditWorkout{
		request: request,
		db:      db,
		logger:  logger.Named("EditWorkout"),
	}
}

type EditWorkout struct {
	request generated.EditWorkout
	db      *pg.DB
	logger  *zap.Logger
}

func (c EditWorkout) Validate(ctx context.Context) []validation.ValidatorFunc {

	return []validation.ValidatorFunc{
		validation.CheckStringNilOrIsNotEmpty(c.request.Name, "Name must not be empty", true),
	}
}

func (c EditWorkout) Call(ctx context.Context) (*workout.Workout, error) {

	var finalWorkout *workout.Workout

	err := c.db.RunInTransaction(func(t *pg.Tx) error {

		wrkout, err := workoutdb.GetWorkoutForUpdate(ctx, c.db, c.request.ID)
		if err != nil {
			if err == pg.ErrNoRows {
				return gqlerror.Errorf("Workout does not exist")
			}

			c.logger.Error("Error retrieving workout", zap.Error(err),
				logging.UUID("workoutID", c.request.ID))
			return err
		}

		if c.request.Name != nil {
			wrkout.Name = *c.request.Name
		}

		if c.request.Description != nil {
			wrkout.Description = *c.request.Description
		}

		finalWorkout, err = workoutdb.UpdateWorkout(ctx, c.db, wrkout)
		if err != nil {
			c.logger.Error("Error updating workout", zap.Error(err),
				logging.UUID("workoutID", c.request.ID))
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return finalWorkout, nil

}
