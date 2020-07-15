package workoutcall

import (
	"context"
	"strings"

	"github.com/jackc/pgx/v4"
	"github.com/train-formula/graphcms/database/types"
	"github.com/train-formula/graphcms/database/workoutdb"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/logging"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
	"github.com/willtrking/pgxload"
	"go.uber.org/zap"
)

func NewEditPrescription(request generated.EditPrescription, logger *zap.Logger, db pgxload.PgxLoader) *EditPrescription {
	return &EditPrescription{
		request: request,
		db:      db,
		kogger:  logger.Named("EditPrescription"),
	}
}

type EditPrescription struct {
	request generated.EditPrescription
	db      pgxload.PgxLoader
	kogger  *zap.Logger
}

func (c EditPrescription) Validate(ctx context.Context) []validation.ValidatorFunc {

	funcs := []validation.ValidatorFunc{
		validation.CheckStringNilOrIsNotEmpty(c.request.Name, "Name must not be empty", true),
		validation.CheckStringNilOrIsNotEmpty(c.request.PrescriptionCategory, "Prescription category must not be empty", true),
	}

	if c.request.DurationSeconds != nil {
		funcs = append(funcs, validation.CheckIntIsNilOrGT(c.request.DurationSeconds.Value, 0, "If duration seconds is set it must be > 0"))
	}

	return funcs
}

func (c EditPrescription) Call(ctx context.Context) (*workout.Prescription, error) {

	var finalPrescription *workout.Prescription

	err := pgxload.RunInTransaction(ctx, c.db, func(ctx context.Context, t pgxload.PgxTxLoader) error {

		prescription, err := workoutdb.GetPrescriptionForUpdate(ctx, t, c.request.ID)
		if err != nil {
			if err == pgx.ErrNoRows {
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
			prescription.DurationSeconds = types.ReadNullInt(c.request.DurationSeconds.Value)
		}

		finalPrescription, err = workoutdb.UpdatePrescription(ctx, t, prescription)
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
