package query

import (
	"context"

	"github.com/train-formula/graphcms/calls/workoutcall"
	"github.com/train-formula/graphcms/database/cursor"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/validation"
)

func (r *QueryResolver) WorkoutCategorySearch(ctx context.Context, request generated.WorkoutCategorySearchRequest, first int, after *string) (*generated.WorkoutCategorySearchResults, error) {

	cursor, err := cursor.DeserializeCursor(after)
	if err != nil {
		return nil, err
	}

	s := workoutcall.SearchWorkoutCategory{
		DB:      r.db,
		First:   first,
		After:   cursor,
		Request: request,
		Logger:  r.logger,
	}

	if validation.ValidationChain(ctx, s.Validate(ctx)...) {

		return s.Call(ctx)
	}

	return nil, nil

}
