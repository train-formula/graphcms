package workoutcall

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/train-formula/graphcms/database/workoutdb"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
	"go.uber.org/zap"
)

func NewSetWorkoutBlockExercises(request generated.SetWorkoutBlockExercises, logger *zap.Logger, db *pg.DB) *SetWorkoutBlockExercises {
	return &SetWorkoutBlockExercises{
		Request: request,
		DB:      db,
		Logger:  logger.Named("SetWorkoutBlockExercises"),
	}
}

type SetWorkoutBlockExercises struct {
	Request generated.SetWorkoutBlockExercises
	DB      *pg.DB
	Logger  *zap.Logger
}

func (c SetWorkoutBlockExercises) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (c SetWorkoutBlockExercises) Call(ctx context.Context) (*workout.WorkoutBlock, error) {

	var finalWorkout *workout.WorkoutBlock
	err := c.DB.RunInTransaction(func(t *pg.Tx) error {

		var err error
		wrkout, err := workoutdb.GetWorkoutBlockForUpdate(ctx, c.DB, c.Request.WorkoutBlockID)
		if err != nil {
			if err == pg.ErrNoRows {
				return gqlerror.Errorf("Workout block does not exist")
			}

			c.Logger.Error("Error retrieving workout block", zap.Error(err))
			return err
		}

		finalWorkout = &wrkout

		if len(c.Request.BlockExercises) <= 0 {
			return workoutdb.ClearWorkoutBlockBlockExercises(ctx, t, wrkout.ID)
		}

		var toSet []workoutdb.ExercisePrescription

		for _, blockExercise := range c.Request.BlockExercises {
			toSet = append(toSet, workoutdb.ExercisePrescription{
				ExerciseID:     blockExercise.ExerciseID,
				PrescriptionID: blockExercise.PrescriptionID,
			})
		}

		return workoutdb.SetWorkoutBlockBlockExercises(ctx, t, c.Request.WorkoutBlockID, toSet)

	})

	if err != nil {
		return nil, err
	}

	return finalWorkout, nil
}
