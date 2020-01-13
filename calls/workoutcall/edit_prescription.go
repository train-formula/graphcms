package workoutcall

import (
	"context"
	"strings"

	"github.com/go-pg/pg/v9"
	"github.com/train-formula/graphcms/database/workoutdb"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/logging"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
	"go.uber.org/zap"
)

func NewEditPrescription(request generated.EditPrescription, logger *zap.Logger, db *pg.DB) *EditPrescription {
	return &EditPrescription{
		request: request,
		db:      db,
		kogger:  logger.Named("EditPrescription"),
	}
}

type EditPrescription struct {
	request generated.EditPrescription
	db      *pg.DB
	kogger  *zap.Logger
}

func (c EditPrescription) Validate(ctx context.Context) []validation.ValidatorFunc {

	return []validation.ValidatorFunc{
		validation.CheckStringNilOrIsNotEmpty(c.request.Name, "Name must not be empty", true),
		validation.CheckStringNilOrIsNotEmpty(c.request.PrescriptionCategory, "Prescription category must not be empty", true),
		func() *gqlerror.Error {
			if c.request.DurationSeconds != nil {
				return nil
			}

			return validation.CheckIntIsNilOrGT(c.request.DurationSeconds.Value, 0, "If duration seconds is set it must be > 0")()
		},
	}

}

func (c EditPrescription) Call(ctx context.Context) (*workout.Prescription, error) {

	var finalPrescription *workout.Prescription

	err := c.db.RunInTransaction(func(t *pg.Tx) error {

		prescription, err := workoutdb.GetPrescriptionForUpdate(ctx, c.db, c.request.ID)
		if err != nil {
			if err == pg.ErrNoRows {
				return gqlerror.Errorf("Prescription does not exist")
			}

			c.kogger.Error("Error retrieving prescription", zap.Error(err),
				logging.UUID("prescriptionID", c.request.ID))
			return err
		}

		if c.request.Name != nil {
			prescription.Name = strings.TrimSpace(*c.request.Name)
		}

		if c.request.PrescriptionCategory != nil {
			prescription.PrescriptionCategory = strings.TrimSpace(*c.request.PrescriptionCategory)
		}

		if c.request.DurationSeconds != nil {
			prescription.DurationSeconds = c.request.DurationSeconds.Value
		}

		finalPrescription, err = workoutdb.UpdatePrescription(ctx, c.db, prescription)
		if err != nil {
			c.kogger.Error("Error updating prescription", zap.Error(err),
				logging.UUID("prescriptionID", c.request.ID))
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return finalPrescription, nil

}
