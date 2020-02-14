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

func NewDeletePrescriptionSet(request uuid.UUID, logger *zap.Logger, db pgxload.PgxLoader) *DeletePrescriptionSet {
	return &DeletePrescriptionSet{
		request: request,
		db:      db,
		logger:  logger.Named("DeletePrescriptionSet"),
	}
}

type DeletePrescriptionSet struct {
	request uuid.UUID
	db      pgxload.PgxLoader
	logger  *zap.Logger
}

func (c DeletePrescriptionSet) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (c DeletePrescriptionSet) Call(ctx context.Context) (*uuid.UUID, error) {

	err := pgxload.RunInTransaction(ctx, c.db, func(ctx context.Context, t pgxload.PgxTxLoader) error {

		_, err := workoutdb.GetPrescriptionSetForUpdate(ctx, t, c.request)
		if err != nil {
			if err == pgx.ErrNoRows {
				return gqlerror.Errorf("Prescription set does not exist")
			}

			c.logger.Error("Error retrieving prescription set", zap.Error(err),
				logging.UUID("prescriptionSetID", c.request))
			return err
		}

		err = workoutdb.DeletePrescriptionSet(ctx, t, c.request)
		if err != nil {
			c.logger.Error("Error deleting prescription set", zap.Error(err),
				logging.UUID("prescriptionSetID", c.request))
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &c.request, nil

}
