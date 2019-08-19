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

	cursor, err := cursor.DeserializeCursor(after)
	if err != nil {
		return nil, err
	}

	s := workoutprogram.Search{
		DB: r.db,
	}

	if s.Validate(ctx, request, first, cursor) {

		return s.Call(ctx, request, first, cursor)
	}

	return nil, nil

}
