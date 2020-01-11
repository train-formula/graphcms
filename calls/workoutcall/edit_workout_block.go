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
		request: request,
		db:      db,
		logger:  logger.Named("EditWorkoutBlock"),
	}
}

type EditWorkoutBlock struct {
	request generated.EditWorkoutBlock
	db      *pg.DB
	logger  *zap.Logger
}

func (c EditWorkoutBlock) Validate(ctx context.Context) []validation.ValidatorFunc {

	return []validation.ValidatorFunc{
		func() *gqlerror.Error {
			if c.request.RoundText != nil {
				return validation.CheckStringNilOrIsNotEmpty(c.request.RoundText.Value, "If round text is set it must not be empty")()
			}

			return nil
		},
		func() *gqlerror.Error {
			if c.request.RoundText != nil {
				return validation.CheckIntIsNilOrGTE(c.request.RoundNumeral.Value, 0, "If round numeral is set it must be >= 0")()
			}

			return nil
		},
		func() *gqlerror.Error {
			if c.request.RoundText != nil {
				return validation.CheckIntIsNilOrGT(c.request.RoundRestDuration.Value, 0, "If round rest duration is set it must be > 0")()
			}

			return nil
		},
		func() *gqlerror.Error {
			if c.request.RoundText != nil {
				return validation.CheckIntIsNilOrGT(c.request.NumberOfRounds.Value, 0, "If number of rounds is set it must be > 0")()
			}

			return nil
		},
		func() *gqlerror.Error {
			if c.request.RoundText != nil {
				return validation.CheckIntIsNilOrGT(c.request.DurationSeconds.Value, 0, "If duration seconds is set it must be > 0")()
			}

			return nil
		},
		func() *gqlerror.Error {
			if c.request.RoundUnitID != nil {
				return validation.UnitIsNilOrExists(ctx, c.db, c.request.RoundUnitID.Value)()
			}

			return nil
		},
	}
}

func (c EditWorkoutBlock) Call(ctx context.Context) (*workout.WorkoutBlock, error) {

	var finalWorkoutBlock *workout.WorkoutBlock

	err := c.db.RunInTransaction(func(t *pg.Tx) error {

		workoutBlock, err := workoutdb.GetWorkoutBlockForUpdate(ctx, c.db, c.request.ID)
		if err != nil {
			if err == pg.ErrNoRows {
				return gqlerror.Errorf("Workout block does not exist")
			}

			c.logger.Error("Error retrieving workout block", zap.Error(err))
			return err
		}

		if c.request.CategoryOrder != nil {
			workoutBlock.CategoryOrder = *c.request.CategoryOrder
		}

		if c.request.RoundNumeral != nil {
			workoutBlock.RoundNumeral = c.request.RoundNumeral.Value
		}

		if c.request.RoundText != nil {
			workoutBlock.RoundText = c.request.RoundText.Value
		}

		if c.request.RoundUnitID != nil {
			workoutBlock.RoundUnitID = c.request.RoundUnitID.Value
		}

		if c.request.RoundRestDuration != nil {
			workoutBlock.RoundRestDuration = c.request.RoundRestDuration.Value
		}

		if c.request.NumberOfRounds != nil {
			workoutBlock.NumberOfRounds = c.request.NumberOfRounds.Value
		}

		if c.request.DurationSeconds != nil {
			workoutBlock.DurationSeconds = c.request.DurationSeconds.Value
		}

		if workoutBlock.RoundNumeral != nil && workoutBlock.RoundUnitID == nil {
			return gqlerror.Errorf("If round numeral is set, round unit ID must also be set")
		}

		finalWorkoutBlock, err = workoutdb.UpdateWorkoutBlock(ctx, c.db, workoutBlock)
		if err != nil {
			c.logger.Error("Error updating workout block", zap.Error(err))
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return finalWorkoutBlock, nil

}
