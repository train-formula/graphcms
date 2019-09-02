package workoutcall

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
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

	return nil
}

func (c CreateWorkoutProgram) Call(ctx context.Context) (*workout.WorkoutProgram, error) {

	newUuid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	_, err = trainerdb.GetOrganization(ctx, c.DB, c.Request.TrainerOrganizationID)

	if err != nil {
		if err == pg.ErrNoRows {
			return nil, gqlerror.Errorf("Organization does not exist")
		}
		return nil, err
	}

	var description string
	if c.Request.Description != nil {
		description = *c.Request.Description
	} else {
		description = ""
	}

	new := workout.WorkoutProgram{
		ID: newUuid,

		Name:                  c.Request.Name,
		Description:           description,
		Public:                true,
		Price:                 "19.99",
		TrainerOrganizationID: c.Request.TrainerOrganizationID,
	}

	return workoutdb.InsertWorkoutProgram(ctx, c.DB, new)
}
