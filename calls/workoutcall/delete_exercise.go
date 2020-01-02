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

func NewDeleteExercise(request uuid.UUID, logger *zap.Logger, db *pg.DB) *DeleteExercise {
	return &DeleteExercise{
		Request: request,
		DB:      db,
		Logger:  logger.Named("DeleteExercise"),
	}
}

type DeleteExercise struct {
	Request uuid.UUID
	DB      *pg.DB
	Logger  *zap.Logger
}

func (c DeleteExercise) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (c DeleteExercise) Call(ctx context.Context) (*uuid.UUID, error) {

	err := c.DB.RunInTransaction(func(t *pg.Tx) error {

		_, err := workoutdb.GetExerciseForUpdate(ctx, t, c.Request)
		if err != nil {
			if err == pg.ErrNoRows {
				return gqlerror.Errorf("Exercise does not exist")
			}

			c.Logger.Error("Error retrieving exercise", zap.Error(err))
			return err
		}

		connected, err := workoutdb.ExerciseConnectedToBlocks(ctx, t, c.Request)
		if err != nil {
			c.Logger.Error("Error checking if exercise is connected to blocks", zap.Error(err))
			return err
		}

		if connected {
			return gqlerror.Errorf("Exercise is still connected to blocks")
		}

		err = workoutdb.DeleteExercise(ctx, t, c.Request)
		if err != nil {
			c.Logger.Error("Error deleting exercise", zap.Error(err))
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &c.Request, nil

}
