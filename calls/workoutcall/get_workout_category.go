package workoutcall

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/dataloader/workoutcategoryid"
	"github.com/train-formula/graphcms/logging"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
	"go.uber.org/zap"
)

func NewGetWorkoutCategory(id uuid.UUID, logger *zap.Logger, db *pg.DB) *GetWorkoutCategory {

	return &GetWorkoutCategory{
		id:     id,
		db:     db,
		logger: logger.Named("GetWorkoutCategory"),
	}
}

type GetWorkoutCategory struct {
	id     uuid.UUID
	db     *pg.DB
	logger *zap.Logger
}

func (g GetWorkoutCategory) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (g GetWorkoutCategory) Call(ctx context.Context) (*workout.WorkoutCategory, error) {

	loader := workoutcategoryid.GetContextLoader(ctx)

	loaded, err := loader.Load(g.id)
	if err != nil {
		g.logger.Error("Failed to load workout category with dataloader", zap.Error(err),
			logging.UUID("workoutCategoryID", g.id))
		return nil, err
	}

	if loaded == nil {
		return nil, gqlerror.Errorf("Unknown workout category ID")
	}

	return loaded, nil
}
