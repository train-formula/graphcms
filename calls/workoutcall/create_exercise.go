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

func NewCreateExercise(request generated.CreateExercise, logger *zap.Logger, db *pg.DB) *CreateExercise {
	return &CreateExercise{
		Request: request,
		DB:      db,
		Logger:  logger.Named("CreateExercise"),
	}
}

type CreateExercise struct {
	Request generated.CreateExercise
	DB      *pg.DB
	Logger  *zap.Logger
}

func (c CreateExercise) Validate(ctx context.Context) []validation.ValidatorFunc {

	return []validation.ValidatorFunc{
		validation.CheckStringIsNotEmpty(c.Request.Name, "Name must not be empty"),
		validation.CheckStringNilOrIsURL(c.Request.VideoURL, "Invalid video URL"),
		validation.OrganizationExists(ctx, c.DB, c.Request.TrainerOrganizationID),
	}
}

func (c CreateExercise) Call(ctx context.Context) (*workout.Exercise, error) {

	newUuid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	var finalExercise *workout.Exercise

	err = c.DB.RunInTransaction(func(t *pg.Tx) error {

		err = validation.TagsAllExistForTrainer(ctx, t, c.Request.TrainerOrganizationID, c.Request.Tags)
		if err != nil {
			return err
		}

		new := workout.Exercise{
			ID: newUuid,

			TrainerOrganizationID: c.Request.TrainerOrganizationID,

			Name:        strings.TrimSpace(c.Request.Name),
			Description: strings.TrimSpace(c.Request.Description),

			VideoURL: c.Request.VideoURL,
		}

		finalExercise, err = workoutdb.InsertExercise(ctx, t, new)

		for _, tagUUID := range c.Request.Tags {

			_, err := tagdb.TagWorkout(ctx, t, tagUUID, c.Request.TrainerOrganizationID, finalExercise.ID)
			if err != nil {
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
