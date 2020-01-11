package query

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/calls/workoutcall"
	"github.com/train-formula/graphcms/database/cursor"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"go.uber.org/zap"
)

func (r *QueryResolver) Exercise(ctx context.Context, id uuid.UUID) (*workout.Exercise, error) {

	call := workoutcall.NewGetExercise(id, r.logger, r.db)

	if validation.ValidationChain(ctx, call.Validate(ctx)...) {
		return call.Call(ctx)
	}

	return nil, nil
}

func (r *QueryResolver) ExerciseSearch(ctx context.Context, request generated.ExerciseSearchRequest, first int, after *string) (*generated.ExerciseSearchResults, error) {

	curse, err := cursor.DeserializeCursor(after)
	if err != nil {
		r.logger.Error("Failed to deserialize cursor", zap.Error(err))
		return nil, err
	}

	s := workoutcall.NewSearchExercises(request, first, curse, r.logger, r.db)

	if validation.ValidationChain(ctx, s.Validate(ctx)...) {

		return s.Call(ctx)
	}

	return nil, nil
}
