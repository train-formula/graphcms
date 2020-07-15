package workoutcall

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database/tagdb"
	"github.com/train-formula/graphcms/database/types"
	"github.com/train-formula/graphcms/database/workoutdb"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"github.com/willtrking/pgxload"
	"go.uber.org/zap"
)

func NewCreateWorkoutProgram(request generated.CreateWorkoutProgram, logger *zap.Logger, db pgxload.PgxLoader) *CreateWorkoutProgram {
	return &CreateWorkoutProgram{
		request: request,
		db:      db,
		logger:  logger.Named("CreateWorkoutProgram"),
	}
}

type CreateWorkoutProgram struct {
	request generated.CreateWorkoutProgram
	db      pgxload.PgxLoader
	logger  *zap.Logger
}

func (c CreateWorkoutProgram) Validate(ctx context.Context) []validation.ValidatorFunc {

	return []validation.ValidatorFunc{
		validation.CheckStringIsNotEmpty(c.request.Name, "Name must not be empty", true),
		validation.OrganizationExists(ctx, c.db, c.request.TrainerOrganizationID),
	}
}

func (c CreateWorkoutProgram) Call(ctx context.Context) (*workout.WorkoutProgram, error) {

	newUuid, err := uuid.NewV4()
	if err != nil {
		c.logger.Error("Failed to generate UUID", zap.Error(err))
		return nil, err
	}

	var description string
	if c.request.Description != nil {
		description = *c.request.Description
	} else {
		description = ""
	}

	var finalProgram *workout.WorkoutProgram

	err = pgxload.RunInTransaction(ctx, c.db, func(ctx context.Context, t pgxload.PgxTxLoader) error {

		err = validation.TagsAllExistForTrainer(ctx, t, c.request.TrainerOrganizationID, c.request.Tags)
		if err != nil {
			return err
		}

		new := workout.WorkoutProgram{
			ID: newUuid,

			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),

			Name:                  c.request.Name,
			Description:           description,
			TrainerOrganizationID: c.request.TrainerOrganizationID,

			ExactStartDate:           c.request.ExactStartDate,
			StartsWhenCustomerStarts: c.request.StartsWhenCustomerStarts,
			NumberOfDays:             types.ReadNullInt(c.request.NumberOfDays),
			ProgramLevel:             c.request.ProgramLevel,
		}

		finalProgram, err = workoutdb.InsertWorkoutProgram(ctx, t, new)
		if err != nil {
			c.logger.Error("Failed to insert workout program", zap.Error(err))
			return err
		}

		for _, tagUUID := range c.request.Tags {

			_, err := tagdb.TagWorkoutProgram(ctx, t, tagUUID, c.request.TrainerOrganizationID, finalProgram.ID)
			if err != nil {
				c.logger.Error("Failed to tag workout program", zap.Error(err))
				return err
			}
		}

		return nil

	})

	if err != nil {
		return nil, err
	}

	return finalProgram, nil
}
