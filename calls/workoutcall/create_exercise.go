package workoutcall

import (
	"context"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database/tagdb"
	"github.com/train-formula/graphcms/database/workoutdb"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"github.com/willtrking/pgxload"
	"go.uber.org/zap"
)

func NewCreateExercise(request generated.CreateExercise, logger *zap.Logger, db pgxload.PgxLoader) *CreateExercise {
	return &CreateExercise{
		request: request,
		db:      db,
		logger:  logger.Named("CreateExercise"),
	}
}

type CreateExercise struct {
	request generated.CreateExercise
	db      pgxload.PgxLoader
	logger  *zap.Logger
}

func (c CreateExercise) Validate(ctx context.Context) []validation.ValidatorFunc {

	return []validation.ValidatorFunc{
		validation.CheckStringIsNotEmpty(c.request.Name, "Name must not be empty", true),
		validation.CheckStringNilOrIsURL(c.request.VideoURL, "Invalid video URL"),
		validation.OrganizationExists(ctx, c.db, c.request.TrainerOrganizationID),
	}
}

func (c CreateExercise) Call(ctx context.Context) (*workout.Exercise, error) {

	newUuid, err := uuid.NewV4()
	if err != nil {
		c.logger.Error("Failed to generate UUID", zap.Error(err))
		return nil, err
	}

	var finalExercise *workout.Exercise

	err = pgxload.RunInTransaction(ctx, c.db, func(ctx context.Context, t pgxload.PgxTxLoader) error {

		err = validation.TagsAllExistForTrainer(ctx, t, c.request.TrainerOrganizationID, c.request.Tags)
		if err != nil {
			return err
		}

		new := workout.Exercise{
			ID: newUuid,

			TrainerOrganizationID: c.request.TrainerOrganizationID,

			Name:        strings.TrimSpace(c.request.Name),
			Description: strings.TrimSpace(c.request.Description),

			VideoURL: c.request.VideoURL,
		}

		finalExercise, err = workoutdb.InsertExercise(ctx, t, new)

		if err != nil {
			c.logger.Error("Failed to insert exercise", zap.Error(err))
			return err
		}

		for _, tagUUID := range c.request.Tags {

			_, err := tagdb.TagExercise(ctx, t, tagUUID, c.request.TrainerOrganizationID, finalExercise.ID)
			if err != nil {
				c.logger.Error("Failed to tag exercise", zap.Error(err))
				return err
			}
		}

		return nil

	})

	if err != nil {
		return nil, err
	}

	return finalExercise, nil
}
