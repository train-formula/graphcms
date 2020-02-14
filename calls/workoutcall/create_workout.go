package workoutcall

import (
	"context"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/train-formula/graphcms/database/tagdb"
	"github.com/train-formula/graphcms/database/workoutdb"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
	"github.com/willtrking/pgxload"
	"go.uber.org/zap"
)

func NewCreateWorkout(request generated.CreateWorkout, logger *zap.Logger, db pgxload.PgxLoader) *CreateWorkout {
	return &CreateWorkout{
		request: request,
		db:      db,
		logger:  logger.Named("CreateWorkout"),
	}
}

type CreateWorkout struct {
	request generated.CreateWorkout
	db      pgxload.PgxLoader
	logger  *zap.Logger
}

func (c CreateWorkout) Validate(ctx context.Context) []validation.ValidatorFunc {

	return []validation.ValidatorFunc{
		validation.CheckStringIsNotEmpty(c.request.Name, "Name must not be empty", true),
		validation.CheckIntGT(c.request.DaysFromStart, 0, "Days from start must be > 0"),
	}
}

func (c CreateWorkout) Call(ctx context.Context) (*workout.Workout, error) {

	newUuid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	var finalWorkout *workout.Workout

	err = pgxload.RunInTransaction(ctx, c.db, func(ctx context.Context, t pgxload.PgxTxLoader) error {

		program, err := workoutdb.GetWorkoutProgram(ctx, t, c.request.WorkoutProgramID)

		if err != nil {
			if err == pgx.ErrNoRows {
				return gqlerror.Errorf("Workout program does not exist")
			}
			c.logger.Error("Failed to retrieve workout program", zap.Error(err))
			return err
		}

		err = validation.TagsAllExistForTrainer(ctx, t, program.TrainerOrganizationID, c.request.Tags)
		if err != nil {
			c.logger.Error("Failed to check if tag exists", zap.Error(err))
			return err
		}

		new := workout.Workout{
			ID: newUuid,

			TrainerOrganizationID: program.TrainerOrganizationID,
			WorkoutProgramID:      program.ID,

			Name:        c.request.Name,
			Description: strings.TrimSpace(c.request.Description),

			DaysFromStart: c.request.DaysFromStart,
		}

		finalWorkout, err = workoutdb.InsertWorkout(ctx, t, new)

		if err != nil {
			c.logger.Error("Failed to insert workout", zap.Error(err))
			return err
		}

		for _, tagUUID := range c.request.Tags {

			_, err := tagdb.TagWorkout(ctx, t, tagUUID, program.TrainerOrganizationID, finalWorkout.ID)
			if err != nil {
				c.logger.Error("Failed to tag workout", zap.Error(err))
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
