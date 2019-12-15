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

type DeleteWorkout struct {
	Request uuid.UUID
	DB      *pg.DB
	Logger  *zap.Logger
}

func (c DeleteWorkout) logger() *zap.Logger {

	return c.Logger.Named("DeleteWorkout")

}

func (c DeleteWorkout) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (c DeleteWorkout) Call(ctx context.Context) (*uuid.UUID, error) {

	err := c.DB.RunInTransaction(func(t *pg.Tx) error {

		_, err := workoutdb.GetWorkoutForUpdate(ctx, t, c.Request)
		if err != nil {
			if err == pg.ErrNoRows {
				return gqlerror.Errorf("Workout does not exist")
			}

			c.logger().Error("Error retrieving workout", zap.Error(err))
			return err
		}

		err = workoutdb.DeleteWorkout(ctx, t, c.Request)
		if err != nil {
			c.logger().Error("Error deleting workout", zap.Error(err))
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &c.Request, nil

}
