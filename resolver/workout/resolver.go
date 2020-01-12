package workout

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/train-formula/graphcms/calls/workoutcall"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
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
		return nil, nil
	}

	g := workoutcall.NewGetWorkoutWorkoutCategories(obj.ID, r.logger, r.db)

	if validation.ValidationChain(ctx, g.Validate(ctx)...) {
		return g.Call(ctx)
	}

	return nil, nil
}
