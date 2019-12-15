package workout

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/train-formula/graphcms/calls/workoutcall"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
	"go.uber.org/zap"
)

type WorkoutResolver struct {
	db     *pg.DB
	logger *zap.Logger
}

func NewWorkoutResolver(db *pg.DB, logger *zap.Logger) *WorkoutResolver {
	return &WorkoutResolver{
		db:     db,
		logger: logger.Named("WorkoutResolver"),
	}
}

func (r *WorkoutResolver) Categories(ctx context.Context, obj *workout.Workout) ([]*workout.WorkoutCategory, error) {

	if obj == nil {
		return nil, gqlerror.Errorf("Cannot locate workout categories from nil workout")
	}

	g := workoutcall.GetWorkoutWorkoutCategories{
		Logger: r.logger,
		ID:     obj.ID,
		DB:     r.db,
	}

	if validation.ValidationChain(ctx, g.Validate(ctx)...) {
		return g.Call(ctx)
	}

	return nil, nil
}
