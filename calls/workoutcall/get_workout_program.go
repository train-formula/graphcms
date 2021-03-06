package workoutcall

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/dataloader/workoutprogramid"
	"github.com/train-formula/graphcms/logging"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
	"github.com/willtrking/pgxload"
	"go.uber.org/zap"
)

func NewGetWorkoutProgram(id uuid.UUID, logger *zap.Logger, db pgxload.PgxLoader) *GetWorkoutProgram {
	return &GetWorkoutProgram{
		id:     id,
		db:     db,
		logger: logger.Named("GetWorkoutProgram"),
	}
}

type GetWorkoutProgram struct {
	id     uuid.UUID
	db     pgxload.PgxLoader
	logger *zap.Logger
}

func (g GetWorkoutProgram) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (g GetWorkoutProgram) Call(ctx context.Context) (*workout.WorkoutProgram, error) {

	loader := workoutprogramid.GetContextLoader(ctx)

	loaded, err := loader.Load(g.id)
	if err != nil {
		g.logger.Error("Failed to load workout program with dataloader", zap.Error(err),
			logging.UUID("workoutProgramID", g.id))
		return nil, err
	}

	if loaded == nil {
		return nil, gqlerror.Errorf("Unknown workout program id")
	}

	return loaded, nil
}
