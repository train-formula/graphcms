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

func (r *PrescriptionSetResolver) Rep(ctx context.Context, obj *workout.PrescriptionSet) (*workout.UnitData, error) {

	if obj == nil {
		return nil, nil
	}

	if obj.RepNumeral == nil && obj.RepText == nil {
		r.logger.Error("Prescription set malformed, missing rep numerap and/or rep text", logging.UUID("prescriptionSet", obj.ID))
		return nil, gqlerror.Errorf("Prescription set malformed, missing rep numerap and/or rep text")
	}

	return &workout.UnitData{
		Numeral: obj.RepNumeral,
		Text:    obj.RepText,
		UnitID:  obj.RepUnitID,
	}, nil
}

func (r *PrescriptionSetResolver) RepModifier(ctx context.Context, obj *workout.PrescriptionSet) (*workout.UnitData, error) {

	if obj == nil {
		return nil, nil
	}

	if obj.RepModifierUnitID == nil && obj.RepModifierNumeral == nil && obj.RepModifierText == nil {
		return nil, nil
	} else if obj.RepModifierUnitID == nil && (obj.RepModifierNumeral != nil || obj.RepModifierText != nil) {
		r.logger.Error("Prescription set malformed, has rep modifier numeral and/or rep modifier text but no rep modifier unit ID", logging.UUID("prescriptionSet", obj.ID))
		return nil, gqlerror.Errorf("Prescription set malformed, has rep modifier numeral and/or rep modifier text but no rep modifier unit ID")
	}

	return &workout.UnitData{
		Numeral: obj.RepModifierNumeral,
		Text:    obj.RepModifierText,
		UnitID:  *obj.RepModifierUnitID,
	}, nil
}
