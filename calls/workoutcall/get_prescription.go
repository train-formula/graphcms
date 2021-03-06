package workoutcall

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/dataloader/prescriptionid"
	"github.com/train-formula/graphcms/logging"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
	"github.com/willtrking/pgxload"
	"go.uber.org/zap"
)

func NewGetPrescription(id uuid.UUID, logger *zap.Logger, db pgxload.PgxLoader) *GetPrescription {
	return &GetPrescription{
		id:     id,
		db:     db,
		logger: logger.Named("GetPrescription"),
	}
}

type GetPrescription struct {
	id     uuid.UUID
	db     pgxload.PgxLoader
	logger *zap.Logger
}

func (g GetPrescription) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (g GetPrescription) Call(ctx context.Context) (*workout.Prescription, error) {

	loader := prescriptionid.GetContextLoader(ctx)

	loaded, err := loader.Load(g.id)
	if err != nil {
		g.logger.Error("Failed to load prescription with dataloader", zap.Error(err),
			logging.UUID("prescriptionID", g.id))
		return nil, err
	}

	if loaded == nil {
		return nil, gqlerror.Errorf("Unknown prescription id")
	}

	return loaded, nil
}
