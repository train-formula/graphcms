package workoutcall

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database/workoutdb"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/logging"
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

	var validators []validation.ValidatorFunc

	if c.request.RoundRestDuration != nil {
		validators = append(validators, validation.CheckIntIsNilOrGT(c.request.RoundRestDuration.Value, 0, "If round rest duration is set it must be > 0"))
	}
	if c.request.NumberOfRounds != nil {
		validators = append(validators, validation.CheckIntIsNilOrGT(c.request.NumberOfRounds.Value, 0, "If number of rounds is set it must be > 0"))
	}

	if c.request.DurationSeconds != nil {
		validators = append(validators, validation.CheckIntIsNilOrGT(c.request.DurationSeconds.Value, 0, "If duration seconds is set it must be > 0"))
	}

	if c.request.Round != nil {
		validators = append(validators, validation.CheckUnitDataValidOrNil(ctx, c.db, c.request.Round.Value, "round"))
	}

	return validators
}

func (c EditWorkoutBlock) Call(ctx context.Context) (*workout.WorkoutBlock, error) {

	var finalWorkoutBlock *workout.WorkoutBlock

	err := c.db.RunInTransaction(func(t *pg.Tx) error {

		workoutBlock, err := workoutdb.GetWorkoutBlockForUpdate(ctx, c.db, c.request.ID)
		if err != nil {
			if err == pg.ErrNoRows {
				return gqlerror.Errorf("Workout block does not exist")
			}

			c.logger.Error("Error retrieving workout block", zap.Error(err),
				logging.UUID("workoutBlockID", c.request.ID))
			return err
		}

		if c.request.CategoryOrder != nil {
			workoutBlock.CategoryOrder = *c.request.CategoryOrder
		}

		if c.request.Round != nil {
			var roundNumeral *int
			var roundText *string
			var roundUnitID *uuid.UUID

			if c.request.Round.Value != nil {
				roundNumeral = c.request.Round.Value.Numeral
				roundText = c.request.Round.Value.Text
				roundUnitID = &c.request.Round.Value.UnitID
			}

			workoutBlock.RoundNumeral = roundNumeral
			workoutBlock.RoundText = roundText
			workoutBlock.RoundUnitID = roundUnitID
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
			return gqlerror.Errorf("If round numeral is set, round unit id must also be set")
		}

		finalWorkoutBlock, err = workoutdb.UpdateWorkoutBlock(ctx, c.db, workoutBlock)
		if err != nil {
			c.logger.Error("Error updating workout block", zap.Error(err),
				logging.UUID("workoutBlockID", c.request.ID))
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return finalWorkoutBlock, nil

}
