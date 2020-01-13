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

func NewDeletePrescriptionSet(request uuid.UUID, logger *zap.Logger, db *pg.DB) *DeletePrescriptionSet {
	return &DeletePrescriptionSet{
		request: request,
		db:      db,
		logger:  logger.Named("DeletePrescriptionSet"),
	}
}

type DeletePrescriptionSet struct {
	request uuid.UUID
	db      *pg.DB
	logger  *zap.Logger
}

func (c DeletePrescriptionSet) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (c DeletePrescriptionSet) Call(ctx context.Context) (*uuid.UUID, error) {

	err := c.db.RunInTransaction(func(t *pg.Tx) error {

		_, err := workoutdb.GetPrescriptionSetForUpdate(ctx, t, c.request)
		if err != nil {
			if err == pg.ErrNoRows {
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
