package workoutcall

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/dataloader/prescriptionid"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
	"go.uber.org/zap"
)

func NewGetPrescription(id uuid.UUID, logger *zap.Logger, db *pg.DB) *GetPrescription {
	return &GetPrescription{
		id:     id,
		db:     db,
		logger: logger.Named("GetPrescription"),
	}
}

type GetPrescription struct {
	id     uuid.UUID
	db     *pg.DB
	logger *zap.Logger
}

func (g GetPrescription) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (g GetPrescription) Call(ctx context.Context) (*workout.Prescription, error) {

	loader := prescriptionid.GetContextLoader(ctx)

	loaded, err := loader.Load(g.id)
	if err != nil {
		g.logger.Error("Failed to load prescription with dataloader", zap.Error(err))
		return nil, err
	}

	if loaded == nil {
		return nil, gqlerror.Errorf("Unknown prescription ID")
	}

	return loaded, nil
}
