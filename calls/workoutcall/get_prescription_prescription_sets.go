package workoutcall

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/dataloader/prescriptionsetsbyprescription"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
	"go.uber.org/zap"
)

func NewGetPrescriptionPrescriptionSets(id uuid.UUID, logger *zap.Logger, db *pg.DB) *GetPrescriptionPrescriptionSets {
	return &GetPrescriptionPrescriptionSets{
		id:     id,
		db:     db,
		logger: logger.Named("GetPrescriptionPrescriptionSets"),
	}
}

type GetPrescriptionPrescriptionSets struct {
	id     uuid.UUID
	db     *pg.DB
	logger *zap.Logger
}

func (g GetPrescriptionPrescriptionSets) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (g GetPrescriptionPrescriptionSets) Call(ctx context.Context) ([]*workout.PrescriptionSet, error) {

	loader := prescriptionsetsbyprescription.GetContextLoader(ctx)

	loaded, err := loader.Load(g.id)
	if err != nil {
		g.logger.Error("Failed to load prescription sets with dataloader", zap.Error(err))
		return nil, err
	}

	if loaded == nil {
		return nil, gqlerror.Errorf("Unknown prescription ID")
	}

	return loaded, nil
}
