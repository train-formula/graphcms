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

func NewDeletePrescription(request uuid.UUID, logger *zap.Logger, db pgxload.PgxLoader) *DeletePrescription {
	return &DeletePrescription{
		request: request,
		db:      db,
		logger:  logger.Named("DeletePrescription"),
	}
}

type DeletePrescription struct {
	request uuid.UUID
	db      pgxload.PgxLoader
	logger  *zap.Logger
}

func (c DeletePrescription) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (c DeletePrescription) Call(ctx context.Context) (*uuid.UUID, error) {

	err := pgxload.RunInTransaction(ctx, c.db, func(ctx context.Context, t pgxload.PgxTxLoader) error {

		_, err := workoutdb.GetPrescriptionForUpdate(ctx, t, c.request)
		if err != nil {
			if err == pgx.ErrNoRows {
				return gqlerror.Errorf("Prescription does not exist")
			}

			c.logger.Error("Error retrieving prescription", zap.Error(err),
				logging.UUID("prescriptionID", c.request))
			return err
		}

		connected, err := workoutdb.PrescriptionConnectedToBlocks(ctx, t, c.request)
		if err != nil {
			c.logger.Error("Error checking if prescription is connected to blocks", zap.Error(err),
				logging.UUID("prescriptionID", c.request))
			return err
		}

		if connected {
			return gqlerror.Errorf("Prescription is still connected to blocks")
		}

		err = workoutdb.DeletePrescription(ctx, t, c.request)
		if err != nil {
			c.logger.Error("Error deleting prescription", zap.Error(err),
				logging.UUID("prescriptionID", c.request))
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &c.request, nil

}
