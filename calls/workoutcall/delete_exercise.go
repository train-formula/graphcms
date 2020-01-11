package workoutcall

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database/workoutdb"
	"github.com/train-formula/graphcms/logging"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
	"go.uber.org/zap"
)

func NewDeleteExercise(request uuid.UUID, logger *zap.Logger, db *pg.DB) *DeleteExercise {
	return &DeleteExercise{
		request: request,
		db:      db,
		logger:  logger.Named("DeleteExercise"),
	}
}

type DeleteExercise struct {
	request uuid.UUID
	db      *pg.DB
	logger  *zap.Logger
}

func (c DeleteExercise) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (c DeleteExercise) Call(ctx context.Context) (*uuid.UUID, error) {

	err := c.db.RunInTransaction(func(t *pg.Tx) error {

		_, err := workoutdb.GetExerciseForUpdate(ctx, t, c.request)
		if err != nil {
			if err == pg.ErrNoRows {
				return gqlerror.Errorf("Exercise does not exist")
			}

			c.logger.Error("Error retrieving exercise", zap.Error(err),
				logging.UUID("exerciseID", c.request))
			return err
		}

		connected, err := workoutdb.ExerciseConnectedToBlocks(ctx, t, c.request)
		if err != nil {
			c.logger.Error("Error checking if exercise is connected to blocks", zap.Error(err),
				logging.UUID("exerciseID", c.request))
			return err
		}

		if connected {
			return gqlerror.Errorf("Exercise is still connected to blocks")
		}

		err = workoutdb.DeleteExercise(ctx, t, c.request)
		if err != nil {
			c.logger.Error("Error deleting exercise", zap.Error(err),
				logging.UUID("exerciseID", c.request))
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &c.request, nil

}
