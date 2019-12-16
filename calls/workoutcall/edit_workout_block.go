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

func NewEditWorkoutBlock(request generated.EditWorkoutBlock, logger *zap.Logger, db *pg.DB) *EditWorkoutBlock {
	return &EditWorkoutBlock{
		Request: request,
		DB:      db,
		Logger:  logger.Named("EditWorkoutBlock"),
	}
}

type EditWorkoutBlock struct {
	Request generated.EditWorkoutBlock
	DB      *pg.DB
	Logger  *zap.Logger
}

func (c EditWorkoutBlock) Validate(ctx context.Context) []validation.ValidatorFunc {

	return []validation.ValidatorFunc{
		func() *gqlerror.Error {
			if c.Request.RoundText != nil {
				return validation.CheckStringNilOrIsNotEmpty(c.Request.RoundText.Value, "If round text is set it must not be empty")()
			}

			return nil
		},
		func() *gqlerror.Error {
			if c.Request.RoundText != nil {
				return validation.CheckIntIsNilOrGTE(c.Request.RoundNumeral.Value, 0, "If round numeral is set it must be >= 0")()
			}

			return nil
		},
		func() *gqlerror.Error {
			if c.Request.RoundText != nil {
				return validation.CheckIntIsNilOrGT(c.Request.RoundRestDuration.Value, 0, "If round rest duration is set it must be > 0")()
			}

			return nil
		},
		func() *gqlerror.Error {
			if c.Request.RoundText != nil {
				return validation.CheckIntIsNilOrGT(c.Request.NumberOfRounds.Value, 0, "If number of rounds is set it must be > 0")()
			}

			return nil
		},
		func() *gqlerror.Error {
			if c.Request.RoundText != nil {
				return validation.CheckIntIsNilOrGT(c.Request.DurationSeconds.Value, 0, "If duration seconds is set it must be > 0")()
			}

			return nil
		},
		func() *gqlerror.Error {
			if c.Request.RoundUnitID != nil {
				return validation.UnitIsNilOrExists(ctx, c.DB, c.Request.RoundUnitID.Value)()
			}

			return nil
		},
	}
}

func (c EditWorkoutBlock) Call(ctx context.Context) (*workout.WorkoutBlock, error) {

	var finalWorkoutBlock *workout.WorkoutBlock

	err := c.DB.RunInTransaction(func(t *pg.Tx) error {

		workoutBlock, err := workoutdb.GetWorkoutBlockForUpdate(ctx, c.DB, c.Request.ID)
		if err != nil {
			if err == pg.ErrNoRows {
				return gqlerror.Errorf("Workout block does not exist")
			}

			c.Logger.Error("Error retrieving workout block", zap.Error(err))
			return err
		}

		if c.Request.CategoryOrder != nil {
			workoutBlock.CategoryOrder = *c.Request.CategoryOrder
		}

		if c.Request.RoundNumeral != nil {
			workoutBlock.RoundNumeral = c.Request.RoundNumeral.Value
		}

		if c.Request.RoundText != nil {
			workoutBlock.RoundText = c.Request.RoundText.Value
		}

		if c.Request.RoundUnitID != nil {
			workoutBlock.RoundUnitID = c.Request.RoundUnitID.Value
		}

		if c.Request.RoundRestDuration != nil {
			workoutBlock.RoundRestDuration = c.Request.RoundRestDuration.Value
		}

		if c.Request.NumberOfRounds != nil {
			workoutBlock.NumberOfRounds = c.Request.NumberOfRounds.Value
		}

		if c.Request.DurationSeconds != nil {
			workoutBlock.DurationSeconds = c.Request.DurationSeconds.Value
		}

		if workoutBlock.RoundNumeral != nil && workoutBlock.RoundUnitID == nil {
			return gqlerror.Errorf("If round numeral is set, round unit ID must also be set")
		}

		finalWorkoutBlock, err = workoutdb.UpdateWorkoutBlock(ctx, c.DB, workoutBlock)
		if err != nil {
			c.Logger.Error("Error updating workout block", zap.Error(err))
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return finalWorkoutBlock, nil

}
