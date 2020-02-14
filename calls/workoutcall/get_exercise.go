package workoutcall

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/dataloader/exerciseid"
	"github.com/train-formula/graphcms/logging"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
	"github.com/willtrking/pgxload"
	"go.uber.org/zap"
)

func NewGetExercise(id uuid.UUID, logger *zap.Logger, db pgxload.PgxLoader) *GetExercise {
	return &GetExercise{
		id:     id,
		db:     db,
		logger: logger.Named("GetExercise"),
	}
}

type GetExercise struct {
	id     uuid.UUID
	db     pgxload.PgxLoader
	logger *zap.Logger
}

func (g GetExercise) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (g GetExercise) Call(ctx context.Context) (*workout.Exercise, error) {

	loader := exerciseid.GetContextLoader(ctx)

	loaded, err := loader.Load(g.id)
	if err != nil {
		g.logger.Error("Failed to load exercise with dataloader", zap.Error(err),
			logging.UUID("exerciseID", g.id))
		return nil, err
	}

	if loaded == nil {
		return nil, gqlerror.Errorf("Unknown exercise id")
	}

	return loaded, nil
}
