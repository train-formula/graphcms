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

func NewDeleteWorkoutBlock(request uuid.UUID, logger *zap.Logger, db *pg.DB) *DeleteWorkoutBlock {
	return &DeleteWorkoutBlock{
		Request: request,
		DB:      db,
		Logger:  logger,
	}
}

type DeleteWorkoutBlock struct {
	Request uuid.UUID
	DB      *pg.DB
	Logger  *zap.Logger
}

func (c DeleteWorkoutBlock) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (c DeleteWorkoutBlock) Call(ctx context.Context) (*uuid.UUID, error) {

	err := c.DB.RunInTransaction(func(t *pg.Tx) error {

		_, err := workoutdb.GetWorkoutBlockForUpdate(ctx, t, c.Request)
		if err != nil {
			if err == pg.ErrNoRows {
				return gqlerror.Errorf("Workout block does not exist")
			}

			c.Logger.Error("Error retrieving workout block", zap.Error(err))
			return err
		}

		err = workoutdb.DeleteWorkoutBlock(ctx, t, c.Request)
		if err != nil {
			c.Logger.Error("Error deleting workout block", zap.Error(err))
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &c.Request, nil

}
