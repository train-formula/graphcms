package workoutcall

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/train-formula/graphcms/database/workoutdb"
	"github.com/train-formula/graphcms/logging"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
	"github.com/willtrking/pgxload"
	"go.uber.org/zap"
)

func NewDeleteWorkoutBlock(request uuid.UUID, logger *zap.Logger, db pgxload.PgxLoader) *DeleteWorkoutBlock {
	return &DeleteWorkoutBlock{
		request: request,
		db:      db,
		logger:  logger.Named("DeleteWorkoutBlock"),
	}
}

type DeleteWorkoutBlock struct {
	request uuid.UUID
	db      pgxload.PgxLoader
	logger  *zap.Logger
}

func (c DeleteWorkoutBlock) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (c DeleteWorkoutBlock) Call(ctx context.Context) (*uuid.UUID, error) {

	err := pgxload.RunInTransaction(ctx, c.db, func(ctx context.Context, t pgxload.PgxTxLoader) error {

		_, err := workoutdb.GetWorkoutBlockForUpdate(ctx, t, c.request)
		if err != nil {
			if err == pgx.ErrNoRows {
				return gqlerror.Errorf("Workout block does not exist")
			}

			c.logger.Error("Error retrieving workout block", zap.Error(err),
				logging.UUID("workoutBlockID", c.request))
			return err
		}

		err = workoutdb.DeleteWorkoutBlock(ctx, t, c.request)
		if err != nil {
			c.logger.Error("Error deleting workout block", zap.Error(err),
				logging.UUID("workoutBlockID", c.request))
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &c.request, nil

}
