package workoutcall

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/dataloader/workoutid"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
	"go.uber.org/zap"
)

type GetWorkout struct {
	ID     uuid.UUID
	DB     *pg.DB
	Logger *zap.Logger
}

func (g GetWorkout) logger() *zap.Logger {

	return g.Logger.Named("GetWorkout")

}

func (g GetWorkout) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (g GetWorkout) Call(ctx context.Context) (*workout.Workout, error) {

	loader := workoutid.GetContextLoader(ctx)

	loaded, err := loader.Load(g.ID)
	if err != nil {
		g.logger().Error("Failed to load workout with dataloader", zap.Error(err))
		return nil, err
	}

	if loaded == nil {
		return nil, gqlerror.Errorf("Unknown workout ID")
	}

	return loaded, nil
}
