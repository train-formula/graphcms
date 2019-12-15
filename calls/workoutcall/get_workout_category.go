package workoutcall

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/dataloader/workoutcategoryid"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
	"go.uber.org/zap"
)

type GetWorkoutCategory struct {
	ID     uuid.UUID
	DB     *pg.DB
	Logger *zap.Logger
}

func (g GetWorkoutCategory) logger() *zap.Logger {

	return g.Logger.Named("GetWorkoutCategory")

}

func (g GetWorkoutCategory) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (g GetWorkoutCategory) Call(ctx context.Context) (*workout.WorkoutCategory, error) {

	loader := workoutcategoryid.GetContextLoader(ctx)

	loaded, err := loader.Load(g.ID)
	if err != nil {
		return nil, err
	}

	if loaded == nil {
		return nil, gqlerror.Errorf("Unknown workout category ID")
	}

	return loaded, nil
}
