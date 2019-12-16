package workoutcall

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database/workoutdb"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
	"go.uber.org/zap"
)

func NewCreateWorkoutBlock(request generated.CreateWorkoutBlock, logger *zap.Logger, db *pg.DB) *CreateWorkoutBlock {
	return &CreateWorkoutBlock{
		Request: request,
		Logger:  logger.Named("CreateWorkoutBlock"),
		DB:      db,
	}
}

type CreateWorkoutBlock struct {
	Request generated.CreateWorkoutBlock
	Logger  *zap.Logger
	DB      *pg.DB
}

func (c CreateWorkoutBlock) Validate(ctx context.Context) []validation.ValidatorFunc {

	return []validation.ValidatorFunc{
		validation.CheckStringNilOrIsNotEmpty(c.Request.RoundText, "If round text is set it must not be empty"),
		validation.CheckIntIsNilOrGTE(c.Request.RoundNumeral, 0, "If round numeral is set it must be >= 0"),
		validation.CheckIntIsNilOrGT(c.Request.RoundRestDuration, 0, "If round rest duration is set it must be > 0"),
		validation.CheckIntIsNilOrGT(c.Request.NumberOfRounds, 0, "If number of rounds is set it must be > 0"),
		validation.CheckIntIsNilOrGT(c.Request.DurationSeconds, 0, "If duration seconds is set it must be > 0"),
		func() *gqlerror.Error {
			if c.Request.RoundNumeral != nil && c.Request.RoundUnitID == nil {
				return gqlerror.Errorf("If round numeral is set, round unit ID must also be set")
			}

			return nil
		},

		validation.UnitIsNilOrExists(ctx, c.DB, c.Request.RoundUnitID),
	}
}

func (c CreateWorkoutBlock) Call(ctx context.Context) (*workout.WorkoutBlock, error) {

	newUuid, err := uuid.NewV4()
	if err != nil {
		c.Logger.Error("Failed to generate UUID", zap.Error(err))
		return nil, err
	}

	var finalWorkoutBlock *workout.WorkoutBlock

	err = c.DB.RunInTransaction(func(t *pg.Tx) error {
		workoutCategory, err := workoutdb.GetWorkoutCategory(ctx, t, c.Request.WorkoutCategoryID)

		if err != nil {
			if err == pg.ErrNoRows {
				return gqlerror.Errorf("Workout category does not exist")
			}
			c.Logger.Error("Failed to load workout category", zap.Error(err))
			return err
		}

		new := workout.WorkoutBlock{
			ID:                    newUuid,
			TrainerOrganizationID: workoutCategory.TrainerOrganizationID,
			WorkoutCategoryID:     workoutCategory.ID,
			CategoryOrder:         c.Request.CategoryOrder,

			RoundNumeral:      c.Request.RoundNumeral,
			RoundText:         c.Request.RoundText,
			RoundUnitID:       c.Request.RoundUnitID,
			DurationSeconds:   c.Request.DurationSeconds,
			RoundRestDuration: c.Request.RoundRestDuration,
			NumberOfRounds:    c.Request.NumberOfRounds,
		}

		finalWorkoutBlock, err = workoutdb.InsertWorkoutBlock(ctx, t, new)

		return nil

	})

	if err != nil {
		return nil, err
	}

	return finalWorkoutBlock, nil
}
