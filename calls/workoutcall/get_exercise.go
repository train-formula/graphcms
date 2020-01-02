package workoutcall

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/dataloader/exerciseid"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
	"go.uber.org/zap"
)

func NewGetExercise(id uuid.UUID, logger *zap.Logger, db *pg.DB) *GetExercise {
	return &GetExercise{
		ID:     id,
		DB:     db,
		Logger: logger.Named("GetExercise"),
	}
}

type GetExercise struct {
	ID     uuid.UUID
	DB     *pg.DB
	Logger *zap.Logger
}

func (g GetExercise) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (g GetExercise) Call(ctx context.Context) (*workout.Exercise, error) {

	loader := exerciseid.GetContextLoader(ctx)

	loaded, err := loader.Load(g.ID)
	if err != nil {
		g.Logger.Error("Failed to load exercise with dataloader", zap.Error(err))
		return nil, err
	}

	if loaded == nil {
		return nil, gqlerror.Errorf("Unknown exercise ID")
	}

	return loaded, nil
}
