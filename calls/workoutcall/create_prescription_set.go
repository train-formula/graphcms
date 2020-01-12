package workoutcall

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database/workoutdb"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
	"go.uber.org/zap"
)

func NewCreatePrescriptionSet(request generated.CreatePrescriptionSet, logger *zap.Logger, db *pg.DB) *CreatePrescriptionSet {
	return &CreatePrescriptionSet{
		request: request,
		db:      db,
		logger:  logger.Named("CreatePrescriptionSet"),
	}
}

type CreatePrescriptionSet struct {
	request generated.CreatePrescriptionSet
	db      *pg.DB
	logger  *zap.Logger
}

func (c CreatePrescriptionSet) Validate(ctx context.Context) []validation.ValidatorFunc {

	var funcs []validation.ValidatorFunc

	if c.request.Data == nil {
		funcs = append(funcs, func() *gqlerror.Error {
			return gqlerror.Errorf("Data is required")
		})
	} else {
		funcs = append(funcs, validation.CheckCreatePrescriptionSetData(ctx, c.db, *c.request.Data, nil)...)
	}

	funcs = append(funcs, validation.PrescriptionExists(ctx, c.db, c.request.PrescriptionID))

	return funcs
}

func (c CreatePrescriptionSet) Call(ctx context.Context) (*workout.PrescriptionSet, error) {

	newUuid, err := uuid.NewV4()
	if err != nil {
		c.logger.Error("Failed to generate UUID", zap.Error(err))
		return nil, err
	}

	var finalPrescriptionSet *workout.PrescriptionSet

	err = c.db.RunInTransaction(func(t *pg.Tx) error {

		var secondaryUnitID *uuid.UUID
		var secondaryNumeral *int
		var secondaryText *string

		if c.request.Data.SecondaryParameter != nil {
			secondaryUnitID = &c.request.Data.SecondaryParameter.UnitID
			secondaryNumeral = c.request.Data.SecondaryParameter.Numeral
			secondaryText = c.request.Data.SecondaryParameter.Text
		}

		newPrescriptionSet := workout.PrescriptionSet{
			ID:                        newUuid,
			PrescriptionID:            c.request.PrescriptionID,
			SetNumber:                 c.request.Data.SetNumber,
			PrimaryParameterNumeral:   c.request.Data.PrimaryParameter.Numeral,
			PrimaryParameterText:      c.request.Data.PrimaryParameter.Text,
			PrimaryParameterUnitID:    c.request.Data.PrimaryParameter.UnitID,
			SecondaryParameterNumeral: secondaryNumeral,
			SecondaryParameterText:    secondaryText,
			SecondaryParameterUnitID:  secondaryUnitID,
		}

		finalPrescriptionSet, err = workoutdb.InsertPrescriptionSet(ctx, t, newPrescriptionSet)
		if err != nil {
			c.logger.Error("Failed to insert prescription set", zap.Error(err))
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return finalPrescriptionSet, nil
}
