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
	"github.com/vektah/gqlparser/gqlerror"
)

type CreateWorkout struct {
	Request generated.CreateWorkout
	DB      *pg.DB
}

func (c CreateWorkout) Validate(ctx context.Context) []validation.ValidatorFunc {

	return []validation.ValidatorFunc{
		validation.CheckStringIsNotEmpty(c.Request.Name, "Name must not be empty"),
		validation.CheckIntGT(c.Request.DaysFromStart, 0, "Days from start must be > 0"),
	}
}

func (c CreateWorkout) Call(ctx context.Context) (*workout.Workout, error) {

	newUuid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	var finalWorkout *workout.Workout

	err = c.DB.RunInTransaction(func(t *pg.Tx) error {
		program, err := workoutdb.GetWorkoutProgram(ctx, t, c.Request.WorkoutProgramID)

		if err != nil {
			if err == pg.ErrNoRows {
				return gqlerror.Errorf("Workout program does not exist")
			}
			return err
		}

		err = validation.TagsAllExistForTrainer(ctx, t, program.TrainerOrganizationID, c.Request.Tags)
		if err != nil {
			return err
		}

		new := workout.Workout{
			ID: newUuid,

			TrainerOrganizationID: program.TrainerOrganizationID,
			WorkoutProgramID:      program.ID,

			Name:        c.Request.Name,
			Description: strings.TrimSpace(c.Request.Description),

			DaysFromStart: c.Request.DaysFromStart,
		}

		finalWorkout, err = workoutdb.InsertWorkout(ctx, t, new)

		for _, tagUUID := range c.Request.Tags {

			_, err := tagdb.TagWorkout(ctx, t, tagUUID, program.TrainerOrganizationID, finalWorkout.ID)
			if err != nil {
				return err
			}
		}

		return nil

	})

	if err != nil {
		return nil, err
	}

	return finalWorkout, nil
}
