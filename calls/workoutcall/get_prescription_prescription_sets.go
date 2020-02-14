package workoutcall

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/dataloader/prescriptionsetsbyprescription"
	"github.com/train-formula/graphcms/logging"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
	"github.com/willtrking/pgxload"
	"go.uber.org/zap"
)

func NewGetPrescriptionPrescriptionSets(prescriptionID uuid.UUID, logger *zap.Logger, db pgxload.PgxLoader) *GetPrescriptionPrescriptionSets {
	return &GetPrescriptionPrescriptionSets{
		prescriptionID: prescriptionID,
		db:             db,
		logger:         logger.Named("GetPrescriptionPrescriptionSets"),
	}
}

type GetPrescriptionPrescriptionSets struct {
	prescriptionID uuid.UUID
	db             pgxload.PgxLoader
	logger         *zap.Logger
}

func (g GetPrescriptionPrescriptionSets) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (g GetPrescriptionPrescriptionSets) Call(ctx context.Context) ([]*workout.PrescriptionSet, error) {

	loader := prescriptionsetsbyprescription.GetContextLoader(ctx)

	loaded, err := loader.Load(g.prescriptionID)
	if err != nil {
		g.logger.Error("Failed to load prescription sets with dataloader", zap.Error(err),
			logging.UUID("prescriptionID", g.prescriptionID))
		return nil, err
	}

	if loaded == nil {
		return nil, gqlerror.Errorf("Unknown prescription id")
	}

	return loaded, nil
}
