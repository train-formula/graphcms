package query

import (
	"context"

	uuid "github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/calls/workoutprogram"
	"github.com/train-formula/graphcms/database/cursor"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models/workout"
)

func (r *QueryResolver) WorkoutProgram(ctx context.Context, id uuid.UUID) (*workout.WorkoutProgram, error) {

	g := workoutprogram.Get{
		DB: r.db,
	}

	if g.Validate(ctx, id) {
		return g.Call(ctx, id)
	}

	return nil, nil
}

func (r *QueryResolver) WorkoutProgramSearch(ctx context.Context, request generated.WorkoutProgramSearchRequest, first int, after *string) (*generated.WorkoutProgramSearchResults, error) {

	curse, err := cursor.DeserializeCursor(after)
	if err != nil {
		return nil, err
	}

	s := workoutprogram.Search{
		DB:      r.db,
		First:   first,
		After:   curse,
		Request: request,
	}

	if s.Validate(ctx) {

		return s.Call(ctx)
	}

	return nil, nil

}
