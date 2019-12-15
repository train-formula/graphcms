package workoutcall

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/train-formula/graphcms/database/workoutdb"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"go.uber.org/zap"
)

type EditWorkoutCategory struct {
	Request generated.EditWorkoutCategory
	DB      *pg.DB
	Logger  *zap.Logger
}

func (c EditWorkoutCategory) logger() *zap.Logger {

	return c.Logger.Named("EditWorkoutCategory")

}

func (c EditWorkoutCategory) Validate(ctx context.Context) []validation.ValidatorFunc {

	return []validation.ValidatorFunc{
		validation.CheckStringNilOrIsNotEmpty(c.Request.Name, "Name must not be empty"),
		/*func() *gqlerror.Error {
			if c.Request.Type != nil && *c.Request.Type == workout.UnknownBlockType {
				return gqlerror.Errorf("Unknown category type")
			}

			return nil
		},*/
	}
}

func (c EditWorkoutCategory) Call(ctx context.Context) (*workout.WorkoutCategory, error) {

	category, err := workoutdb.GetWorkoutCategoryForUpdate(ctx, c.DB, c.Request.ID)
	if err != nil {
		c.logger().Error("Error retrieving workout category", zap.Error(err))
		return nil, err
	}

	if c.Request.Name != nil {
		category.Name = *c.Request.Name
	}

	if c.Request.Description != nil {
		category.Description = *c.Request.Description
	}

	/*if c.Request.Type != nil {
		category.Type = *c.Request.Type
	}

	if c.Request.RoundNumeral.ContainsValue() {
		category.RoundNumeral = c.Request.RoundNumeral.Value
	}

	if c.Request.RoundText.ContainsValue() {
		category.RoundText = c.Request.RoundText.Value
	}

	if c.Request.RoundUnitID.ContainsValue() {
		category.RoundUnitID = c.Request.RoundUnitID.Value
	}

	if c.Request.DurationSeconds.ContainsValue() {
		category.DurationSeconds = c.Request.DurationSeconds.Value
	}

	// Do final validation of category after edits here
	if category.Type == workout.GeneralBlockType {
		if category.RoundNumeral != nil || category.RoundText != nil || category.RoundUnitID != nil || category.DurationSeconds != nil {
			return nil, gqlerror.Errorf("General workout categories may not specify round or duration data")
		}
	} else if category.Type == workout.RoundBlockType {
		if category.RoundNumeral == nil || category.RoundUnitID == nil {
			return nil, gqlerror.Errorf("Round workout categories must specify round numeral and unit")
		}

		if category.DurationSeconds != nil {
			return nil, gqlerror.Errorf("Round workout categories may not specify duration, use TimedRound")
		}
	} else if category.Type == workout.TimedRoundBlockType {
		if c.Request.RoundNumeral == nil || c.Request.RoundUnitID == nil {
			return nil, gqlerror.Errorf("Timed round workout categories must specify round numeral and unit")
		}

		if category.DurationSeconds == nil {
			return nil, gqlerror.Errorf("Timed round workout categories must specify duration")
		}
	}*/

	return workoutdb.UpdateWorkoutCategory(ctx, c.DB, category)

}
