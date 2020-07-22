package prescriptionset

import (
	"context"

	"github.com/train-formula/graphcms/calls/workoutcall"
	"github.com/train-formula/graphcms/logging"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
	"github.com/willtrking/pgxload"
	"go.uber.org/zap"
)

type PrescriptionSetResolver struct {
	db     pgxload.PgxLoader
	logger *zap.Logger
}

func NewPrescriptionSetResolver(db pgxload.PgxLoader, logger *zap.Logger) *PrescriptionSetResolver {
	return &PrescriptionSetResolver{
		db:     db,
		logger: logger.Named("PrescriptionSetResolver"),
	}
}

func (r *PrescriptionSetResolver) Prescription(ctx context.Context, obj *workout.PrescriptionSet) (*workout.Prescription, error) {
	if obj == nil {
		return nil, nil
	}

	g := workoutcall.NewGetPrescription(obj.PrescriptionID, r.logger, r.db)

	if validation.ValidationChain(ctx, g.Validate(ctx)...) {
		return g.Call(ctx)
	}

	return nil, nil
}

func (r *PrescriptionSetResolver) PrimaryParameter(ctx context.Context, obj *workout.PrescriptionSet) (*workout.UnitData, error) {

	if obj == nil {
		return nil, nil
	}

	if !obj.PrimaryParameterNumeral.Valid && !obj.PrimaryParameterText.Valid {
		r.logger.Error("Prescription set malformed, missing primary parameter numeral and/or primary parameter text", logging.UUID("prescriptionSet", obj.ID))
		return nil, gqlerror.Errorf("Prescription set malformed, missing primary parameter numeral and/or primary parameter text")
	}

	return &workout.UnitData{
		Numeral: obj.PrimaryParameterNumeral,
		Text:    obj.PrimaryParameterText,
		UnitID:  obj.PrimaryParameterUnitID,
	}, nil
}

func (r *PrescriptionSetResolver) SecondaryParameter(ctx context.Context, obj *workout.PrescriptionSet) (*workout.UnitData, error) {

	if obj == nil {
		return nil, nil
	}

	if obj.SecondaryParameterUnitID == nil && !obj.SecondaryParameterNumeral.Valid && !obj.SecondaryParameterText.Valid {
		return nil, nil
	} else if obj.SecondaryParameterUnitID == nil && (obj.SecondaryParameterNumeral.Valid || obj.SecondaryParameterText.Valid) {
		r.logger.Error("Prescription set malformed, has secondary parameter numeral and/or secondary parameter text but no secondary parameter unit ID", logging.UUID("prescriptionSet", obj.ID))
		return nil, gqlerror.Errorf("Prescription set malformed, has secondary parameter numeral and/or secondary parameter text but no secondary parameter unit ID")
	}

	return &workout.UnitData{
		Numeral: obj.SecondaryParameterNumeral,
		Text:    obj.SecondaryParameterText,
		UnitID:  *obj.SecondaryParameterUnitID,
	}, nil
}
