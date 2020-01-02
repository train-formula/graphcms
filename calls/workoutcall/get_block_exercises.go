package workoutcall

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/dataloader/blockexercisesbyblock"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"go.uber.org/zap"
)

func NewGetBlockExercises(id uuid.UUID, logger *zap.Logger, db *pg.DB) *GetBlockExercises {

	return &GetBlockExercises{
		ID:     id,
		DB:     db,
		Logger: logger.Named("GetBlockExercises"),
	}
}

type GetBlockExercises struct {
	ID     uuid.UUID
	DB     *pg.DB
	Logger *zap.Logger
}

func (g GetBlockExercises) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (g GetBlockExercises) Call(ctx context.Context) ([]*workout.BlockExercise, error) {

	loader := blockexercisesbyblock.GetContextLoader(ctx)

	loaded, err := loader.Load(g.ID)
	if err != nil {
		g.Logger.Error("Failed to load block exercises with dataloader", zap.Error(err))
		return nil, err
	}

	return loaded, nil
}
