package workoutcall

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/dataloader/blockexercisesbyblock"
	"github.com/train-formula/graphcms/logging"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"github.com/willtrking/pgxload"
	"go.uber.org/zap"
)

func NewGetBlockExercises(workoutBlockID uuid.UUID, logger *zap.Logger, db pgxload.PgxLoader) *GetBlockExercises {

	return &GetBlockExercises{
		workoutBlockID: workoutBlockID,
		db:             db,
		logger:         logger.Named("GetBlockExercises"),
	}
}

type GetBlockExercises struct {
	workoutBlockID uuid.UUID
	db             pgxload.PgxLoader
	logger         *zap.Logger
}

func (g GetBlockExercises) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (g GetBlockExercises) Call(ctx context.Context) ([]*workout.BlockExercise, error) {

	loader := blockexercisesbyblock.GetContextLoader(ctx)

	loaded, err := loader.Load(g.workoutBlockID)
	if err != nil {
		g.logger.Error("Failed to load block exercises with dataloader", zap.Error(err),
			logging.UUID("workoutBlockID", g.workoutBlockID))
		return nil, err
	}

	return loaded, nil
}
