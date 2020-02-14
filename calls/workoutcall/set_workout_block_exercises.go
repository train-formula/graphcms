package workoutcall

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/train-formula/graphcms/database/workoutdb"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
	"github.com/willtrking/pgxload"
	"go.uber.org/zap"
)

func NewSetWorkoutBlockExercises(request generated.SetWorkoutBlockExercises, logger *zap.Logger, db pgxload.PgxLoader) *SetWorkoutBlockExercises {
	return &SetWorkoutBlockExercises{
		request: request,
		db:      db,
		logger:  logger.Named("SetWorkoutBlockExercises"),
	}
}

type SetWorkoutBlockExercises struct {
	request generated.SetWorkoutBlockExercises
	db      pgxload.PgxLoader
	logger  *zap.Logger
}

func (c SetWorkoutBlockExercises) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (c SetWorkoutBlockExercises) Call(ctx context.Context) (*workout.WorkoutBlock, error) {

	var finalWorkout *workout.WorkoutBlock
	err := pgxload.RunInTransaction(ctx, c.db, func(ctx context.Context, t pgxload.PgxTxLoader) error {

		var err error
		wrkout, err := workoutdb.GetWorkoutBlockForUpdate(ctx, t, c.request.WorkoutBlockID)
		if err != nil {
			if err == pgx.ErrNoRows {
				return gqlerror.Errorf("Workout block does not exist")
			}

			c.logger.Error("Error retrieving workout block", zap.Error(err))
			return err
		}

		finalWorkout = &wrkout

		if len(c.request.BlockExercises) <= 0 {
			err = workoutdb.ClearWorkoutBlockBlockExercises(ctx, t, wrkout.ID)
			if err != nil {
				c.logger.Error("Failed to clear workout block block exercises", zap.Error(err))
				return err
			}
			return nil
		}

		var toSet []workoutdb.ExercisePrescription

		for _, blockExercise := range c.request.BlockExercises {
			toSet = append(toSet, workoutdb.ExercisePrescription{
				ExerciseID:     blockExercise.ExerciseID,
				PrescriptionID: blockExercise.PrescriptionID,
			})
		}

		err = workoutdb.SetWorkoutBlockBlockExercises(ctx, t, c.request.WorkoutBlockID, toSet)

		if err != nil {
			c.logger.Error("Failed to set workout block block exercises", zap.Error(err))
			return err
		}

		return nil

	})

	if err != nil {
		return nil, err
	}

	return finalWorkout, nil
}
