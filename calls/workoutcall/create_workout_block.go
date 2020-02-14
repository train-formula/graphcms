package workoutcall

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/train-formula/graphcms/database/workoutdb"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
	"github.com/willtrking/pgxload"
	"go.uber.org/zap"
)

func NewCreateWorkoutBlock(request generated.CreateWorkoutBlock, logger *zap.Logger, db pgxload.PgxLoader) *CreateWorkoutBlock {
	return &CreateWorkoutBlock{
		request: request,
		logger:  logger.Named("CreateWorkoutBlock"),
		db:      db,
	}
}

type CreateWorkoutBlock struct {
	request generated.CreateWorkoutBlock
	logger  *zap.Logger
	db      pgxload.PgxLoader
}

func (c CreateWorkoutBlock) Validate(ctx context.Context) []validation.ValidatorFunc {

	return []validation.ValidatorFunc{
		validation.CheckIntIsNilOrGT(c.request.RoundRestDuration, 0, "If round rest duration is set it must be > 0"),
		validation.CheckIntIsNilOrGT(c.request.NumberOfRounds, 0, "If number of rounds is set it must be > 0"),
		validation.CheckIntIsNilOrGT(c.request.DurationSeconds, 0, "If duration seconds is set it must be > 0"),
		validation.CheckUnitDataValidOrNil(ctx, c.db, c.request.Round, "round"),
	}
}

func (c CreateWorkoutBlock) Call(ctx context.Context) (*workout.WorkoutBlock, error) {

	newUuid, err := uuid.NewV4()
	if err != nil {
		c.logger.Error("Failed to generate UUID", zap.Error(err))
		return nil, err
	}

	var finalWorkoutBlock *workout.WorkoutBlock

	err = pgxload.RunInTransaction(ctx, c.db, func(ctx context.Context, t pgxload.PgxTxLoader) error {

		workoutCategory, err := workoutdb.GetWorkoutCategory(ctx, t, c.request.WorkoutCategoryID)

		if err != nil {
			if err == pgx.ErrNoRows {
				return gqlerror.Errorf("Workout category does not exist")
			}
			c.logger.Error("Failed to load workout category", zap.Error(err))
			return err
		}

		var roundNumeral *int
		var roundText *string
		var roundUnitID *uuid.UUID

		if c.request.Round != nil {
			roundNumeral = c.request.Round.Numeral
			roundText = c.request.Round.Text
			roundUnitID = &c.request.Round.UnitID
		}

		new := workout.WorkoutBlock{
			ID:                    newUuid,
			TrainerOrganizationID: workoutCategory.TrainerOrganizationID,
			WorkoutCategoryID:     workoutCategory.ID,
			CategoryOrder:         c.request.CategoryOrder,

			RoundNumeral:      roundNumeral,
			RoundText:         roundText,
			RoundUnitID:       roundUnitID,
			DurationSeconds:   c.request.DurationSeconds,
			RoundRestDuration: c.request.RoundRestDuration,
			NumberOfRounds:    c.request.NumberOfRounds,
		}

		finalWorkoutBlock, err = workoutdb.InsertWorkoutBlock(ctx, t, new)

		if err != nil {
			c.logger.Error("Failed to insert workout block", zap.Error(err))
			return err
		}

		return nil

	})

	if err != nil {
		return nil, err
	}

	return finalWorkoutBlock, nil
}
