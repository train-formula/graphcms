package workoutcall

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/dataloader/workoutprogramid"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
	"go.uber.org/zap"
)

func NewGetWorkoutProgram(id uuid.UUID, logger *zap.Logger, db *pg.DB) *GetWorkoutProgram {
	return &GetWorkoutProgram{
		id:     id,
		db:     db,
		logger: logger.Named("GetWorkoutProgram"),
	}
}

type GetWorkoutProgram struct {
	id     uuid.UUID
	db     *pg.DB
	logger *zap.Logger
}

func (g GetWorkoutProgram) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (g GetWorkoutProgram) Call(ctx context.Context) (*workout.WorkoutProgram, error) {

	loader := workoutprogramid.GetContextLoader(ctx)

	loaded, err := loader.Load(g.id)
	if err != nil {
		g.logger.Error("Failed to load workout program with dataloader", zap.Error(err))
		return nil, err
	}

	if loaded == nil {
		return nil, gqlerror.Errorf("Unknown workout program ID")
	}

	return loaded, nil
}
