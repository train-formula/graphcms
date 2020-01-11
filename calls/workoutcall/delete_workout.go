package workoutcall

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database/workoutdb"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
	"go.uber.org/zap"
)

func NewDeleteWorkout(request uuid.UUID, logger *zap.Logger, db *pg.DB) *DeleteWorkout {
	return &DeleteWorkout{
		request: request,
		db:      db,
		logger:  logger.Named("DeleteWorkout"),
	}
}

type DeleteWorkout struct {
	request uuid.UUID
	db      *pg.DB
	logger  *zap.Logger
}

func (c DeleteWorkout) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (c DeleteWorkout) Call(ctx context.Context) (*uuid.UUID, error) {

	err := c.db.RunInTransaction(func(t *pg.Tx) error {

		_, err := workoutdb.GetWorkoutForUpdate(ctx, t, c.request)
		if err != nil {
			if err == pg.ErrNoRows {
				return gqlerror.Errorf("Workout does not exist")
			}

			c.logger.Error("Error retrieving workout", zap.Error(err))
			return err
		}

		err = workoutdb.DeleteWorkout(ctx, t, c.request)
		if err != nil {
			c.logger.Error("Error deleting workout", zap.Error(err))
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &c.request, nil

}
