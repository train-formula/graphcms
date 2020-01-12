package prescriptionset

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/train-formula/graphcms/logging"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/vektah/gqlparser/gqlerror"
	"go.uber.org/zap"
)

type PrescriptionSetResolver struct {
	db     *pg.DB
	logger *zap.Logger
}

func NewPrescriptionSetResolver(db *pg.DB, logger *zap.Logger) *PrescriptionSetResolver {
	return &PrescriptionSetResolver{
		db:     db,
		logger: logger.Named("PrescriptionSetResolver"),
	}
}

func (r *PrescriptionSetResolver) PrimaryParameter(ctx context.Context, obj *workout.PrescriptionSet) (*workout.UnitData, error) {

	if obj == nil {
		return nil, nil
	}

	if obj.PrimaryParameterNumeral == nil && obj.PrimaryParameterText == nil {
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

	if obj.SecondaryParameterUnitID == nil && obj.SecondaryParameterNumeral == nil && obj.SecondaryParameterText == nil {
		return nil, nil
	} else if obj.SecondaryParameterUnitID == nil && (obj.SecondaryParameterNumeral != nil || obj.SecondaryParameterText != nil) {
		r.logger.Error("Prescription set malformed, has secondary parameter numeral and/or secondary parameter text but no secondary parameter unit ID", logging.UUID("prescriptionSet", obj.ID))
		return nil, gqlerror.Errorf("Prescription set malformed, has secondary parameter numeral and/or secondary parameter text but no secondary parameter unit ID")
	}

	return &workout.UnitData{
		Numeral: obj.SecondaryParameterNumeral,
		Text:    obj.SecondaryParameterText,
		UnitID:  *obj.SecondaryParameterUnitID,
	}, nil
}
