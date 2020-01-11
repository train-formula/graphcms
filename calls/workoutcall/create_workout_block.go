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
		request: request,
		logger:  logger.Named("CreateWorkoutBlock"),
		db:      db,
	}
}

type CreateWorkoutBlock struct {
	request generated.CreateWorkoutBlock
	logger  *zap.Logger
	db      *pg.DB
}

func (c CreateWorkoutBlock) Validate(ctx context.Context) []validation.ValidatorFunc {

	return []validation.ValidatorFunc{
		validation.CheckStringNilOrIsNotEmpty(c.request.RoundText, "If round text is set it must not be empty"),
		validation.CheckIntIsNilOrGTE(c.request.RoundNumeral, 0, "If round numeral is set it must be >= 0"),
		validation.CheckIntIsNilOrGT(c.request.RoundRestDuration, 0, "If round rest duration is set it must be > 0"),
		validation.CheckIntIsNilOrGT(c.request.NumberOfRounds, 0, "If number of rounds is set it must be > 0"),
		validation.CheckIntIsNilOrGT(c.request.DurationSeconds, 0, "If duration seconds is set it must be > 0"),
		func() *gqlerror.Error {
			if c.request.RoundNumeral != nil && c.request.RoundUnitID == nil {
				return gqlerror.Errorf("If round numeral is set, round unit ID must also be set")
			}

			return nil
		},

		validation.UnitIsNilOrExists(ctx, c.db, c.request.RoundUnitID),
	}
}

func (c CreateWorkoutBlock) Call(ctx context.Context) (*workout.WorkoutBlock, error) {

	newUuid, err := uuid.NewV4()
	if err != nil {
		c.logger.Error("Failed to generate UUID", zap.Error(err))
		return nil, err
	}

	var finalWorkoutBlock *workout.WorkoutBlock

	err = c.db.RunInTransaction(func(t *pg.Tx) error {
		workoutCategory, err := workoutdb.GetWorkoutCategory(ctx, t, c.request.WorkoutCategoryID)

		if err != nil {
			if err == pg.ErrNoRows {
				return gqlerror.Errorf("Workout category does not exist")
			}
			c.logger.Error("Failed to load workout category", zap.Error(err))
			return err
		}

		new := workout.WorkoutBlock{
			ID:                    newUuid,
			TrainerOrganizationID: workoutCategory.TrainerOrganizationID,
			WorkoutCategoryID:     workoutCategory.ID,
			CategoryOrder:         c.request.CategoryOrder,

			RoundNumeral:      c.request.RoundNumeral,
			RoundText:         c.request.RoundText,
			RoundUnitID:       c.request.RoundUnitID,
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
