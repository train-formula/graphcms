package workoutcall

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database/tagdb"
	"github.com/train-formula/graphcms/database/trainerdb"
	"github.com/train-formula/graphcms/database/workoutdb"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
)

type CreateWorkoutProgram struct {
	Request generated.CreateWorkoutProgram
	DB      *pg.DB
}

func (c CreateWorkoutProgram) Validate(ctx context.Context) []validation.ValidatorFunc {

	return []validation.ValidatorFunc{
		validation.CheckStringIsNotEmpty(c.Request.Name, "Name must not be empty"),
	}
}

func (c CreateWorkoutProgram) Call(ctx context.Context) (*workout.WorkoutProgram, error) {

	newUuid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	var description string
	if c.Request.Description != nil {
		description = *c.Request.Description
	} else {
		description = ""
	}

	var finalProgram *workout.WorkoutProgram

	err = c.DB.RunInTransaction(func(t *pg.Tx) error {
		_, err := trainerdb.GetOrganization(ctx, c.DB, c.Request.TrainerOrganizationID)

		if err != nil {
			if err == pg.ErrNoRows {
				return gqlerror.Errorf("Organization does not exist")
			}
			return err
		}

		err = validation.TagsAllExistForTrainer(ctx, t, c.Request.TrainerOrganizationID, c.Request.Tags)
		if err != nil {
			return err
		}

		new := workout.WorkoutProgram{
			ID: newUuid,

			Name:                  c.Request.Name,
			Description:           description,
			TrainerOrganizationID: c.Request.TrainerOrganizationID,
		}

		finalProgram, err = workoutdb.InsertWorkoutProgram(ctx, c.DB, new)

		for _, tagUUID := range c.Request.Tags {

			_, err := tagdb.TagWorkoutProgram(ctx, t, tagUUID, c.Request.TrainerOrganizationID, finalProgram.ID)
			if err != nil {
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
