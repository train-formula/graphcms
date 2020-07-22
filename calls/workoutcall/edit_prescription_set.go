package workoutcall

import (
	"context"

	"github.com/gofrs/uuid"
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

func NewEditPrescriptionSet(request generated.EditPrescriptionSet, logger *zap.Logger, db pgxload.PgxLoader) *EditPrescriptionSet {
	return &EditPrescriptionSet{
		request: request,
		db:      db,
		kogger:  logger.Named("EditPrescriptionSet"),
	}
}

type EditPrescriptionSet struct {
	request generated.EditPrescriptionSet
	db      pgxload.PgxLoader
	kogger  *zap.Logger
}

func (c EditPrescriptionSet) Validate(ctx context.Context) []validation.ValidatorFunc {

	funcs := []validation.ValidatorFunc{
		validation.CheckIntIsNilOrGT(c.request.SetNumber, 0, "Set number must be > 0"),
		validation.CheckUnitDataValidOrNil(ctx, c.db, c.request.PrimaryParameter, "primaryParameter"),
	}

	if c.request.SecondaryParameter != nil {
		funcs = append(funcs, validation.CheckUnitDataValidOrNil(ctx, c.db, c.request.SecondaryParameter.Value, "secondaryParameter"))
	}

	return funcs

}

func (c EditPrescriptionSet) Call(ctx context.Context) (*workout.PrescriptionSet, error) {

	var finalPrescriptionSet *workout.PrescriptionSet

	err := pgxload.RunInTransaction(ctx, c.db, func(ctx context.Context, t pgxload.PgxTxLoader) error {

		prescriptionSet, err := workoutdb.GetPrescriptionSetForUpdate(ctx, t, c.request.ID)
		if err != nil {
			if err == pgx.ErrNoRows {
				return gqlerror.Errorf("Prescription set does not exist")
			}

			c.kogger.Error("Error retrieving prescription set", zap.Error(err),
				logging.UUID("prescriptionSetID", c.request.ID))
			return err
		}

		if c.request.SetNumber != nil {
			prescriptionSet.SetNumber = *c.request.SetNumber
		}

		if c.request.PrimaryParameter != nil {
			prescriptionSet.PrimaryParameterNumeral = types.ReadNullInt(c.request.PrimaryParameter.Numeral)
			prescriptionSet.PrimaryParameterText = types.ReadNullString(c.request.PrimaryParameter.Text)
			prescriptionSet.PrimaryParameterUnitID = c.request.PrimaryParameter.UnitID
		}

		if c.request.SecondaryParameter != nil {
			var secondaryUnitID *uuid.UUID
			var secondaryNumeral *int
			var secondaryText *string

			if c.request.SecondaryParameter.Value != nil {
				secondaryUnitID = &c.request.SecondaryParameter.Value.UnitID
				secondaryNumeral = c.request.SecondaryParameter.Value.Numeral
				secondaryText = c.request.SecondaryParameter.Value.Text
			}

			prescriptionSet.SecondaryParameterNumeral = types.ReadNullInt(secondaryNumeral)
			prescriptionSet.SecondaryParameterText = types.ReadNullString(secondaryText)
			prescriptionSet.SecondaryParameterUnitID = secondaryUnitID
		}

		finalPrescriptionSet, err = workoutdb.UpdatePrescriptionSet(ctx, t, prescriptionSet)
		if err != nil {
			c.kogger.Error("Error updating prescription set", zap.Error(err),
				logging.UUID("prescriptionSetID", c.request.ID))
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return finalPrescriptionSet, nil

}
