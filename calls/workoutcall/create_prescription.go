package workoutcall

import (
	"context"
	"strings"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database/tagdb"
	"github.com/train-formula/graphcms/database/workoutdb"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"go.uber.org/zap"
)

func NewCreatePrescription(request generated.CreatePrescription, logger *zap.Logger, db *pg.DB) *CreatePrescription {
	return &CreatePrescription{
		request: request,
		db:      db,
		logger:  logger.Named("CreatePrescription"),
	}
}

type CreatePrescription struct {
	request generated.CreatePrescription
	db      *pg.DB
	logger  *zap.Logger
}

func (c CreatePrescription) Validate(ctx context.Context) []validation.ValidatorFunc {

	funcs := []validation.ValidatorFunc{
		validation.CheckStringIsNotEmpty(c.request.Name, "Name must not be empty", true),
		validation.CheckStringIsNotEmpty(c.request.PrescriptionCategory, "Prescription category must not be empty", true),
		validation.CheckIntIsNilOrGT(c.request.DurationSeconds, 0, "If duration seconds is set it must be > 0"),
		validation.OrganizationExists(ctx, c.db, c.request.TrainerOrganizationID),
	}

	for setIdx, set := range c.request.Sets {

		if set != nil {
			funcs = append(funcs, validation.CheckCreatePrescriptionSetWithPrescription(ctx, c.db, *set, &setIdx)...)
		}

	}

	return funcs
}

func (c CreatePrescription) Call(ctx context.Context) (*workout.Prescription, error) {

	newUuid, err := uuid.NewV4()
	if err != nil {
		c.logger.Error("Failed to generate UUID", zap.Error(err))
		return nil, err
	}

	var finalPrescription *workout.Prescription

	err = c.db.RunInTransaction(func(t *pg.Tx) error {

		err = validation.TagsAllExistForTrainer(ctx, t, c.request.TrainerOrganizationID, c.request.Tags)
		if err != nil {
			c.logger.Error("Failed to check if tag exists", zap.Error(err))
			return err
		}

		newPrescription := workout.Prescription{
			ID:                    newUuid,
			TrainerOrganizationID: c.request.TrainerOrganizationID,
			Name:                  strings.TrimSpace(c.request.Name),
			PrescriptionCategory:  strings.TrimSpace(c.request.PrescriptionCategory),
			DurationSeconds:       c.request.DurationSeconds,
		}

		finalPrescription, err = workoutdb.InsertPrescription(ctx, t, newPrescription)
		if err != nil {
			c.logger.Error("Failed to insert prescription", zap.Error(err))
			return err
		}

		for _, set := range c.request.Sets {
			newSetUuid, err := uuid.NewV4()
			if err != nil {
				c.logger.Error("Failed to generate UUID for prescription set", zap.Error(err))
				return err
			}

			newPrescriptionSet := workout.PrescriptionSet{
				ID:                    newSetUuid,
				TrainerOrganizationID: finalPrescription.TrainerOrganizationID,
				PrescriptionID:        finalPrescription.ID,
				SetNumber:             set.SetNumber,
				RepNumeral:            set.RepNumeral,
				RepText:               set.RepText,
				RepUnitID:             set.RepUnitID,
				RepModifierNumeral:    set.RepModifierNumeral,
				RepModifierText:       set.RepModifierText,
				RepModifierUnitID:     set.RepModifierUnitID,
			}

			_, err = workoutdb.InsertPrescriptionSet(ctx, t, newPrescriptionSet)
			if err != nil {
				c.logger.Error("Failed to create prescription set", zap.Error(err))
				return err
			}
		}

		for _, tagUUID := range c.request.Tags {

			_, err := tagdb.TagPrescription(ctx, t, tagUUID, c.request.TrainerOrganizationID, finalPrescription.ID)
			if err != nil {
				c.logger.Error("Failed to tag prescription", zap.Error(err))
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return finalPrescription, nil
}
