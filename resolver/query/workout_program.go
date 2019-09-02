package query

import (
	"context"

	uuid "github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/calls/workoutcall"
	"github.com/train-formula/graphcms/database/cursor"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
)

func (r *QueryResolver) WorkoutProgram(ctx context.Context, id uuid.UUID) (*workout.WorkoutProgram, error) {

	g := workoutcall.GetWorkoutProgram{
		ID: id,
		DB: r.db,
	}

	if validation.ValidationChain(ctx, g.Validate(ctx)...) {
		return g.Call(ctx)
	}

	return nil, nil
}

func (r *QueryResolver) WorkoutProgramSearch(ctx context.Context, request generated.WorkoutProgramSearchRequest, first int, after *string) (*generated.WorkoutProgramSearchResults, error) {

	curse, err := cursor.DeserializeCursor(after)
	if err != nil {
		return nil, err
	}

	s := workoutcall.SearchWorkoutProgram{
		DB:      r.db,
		First:   first,
		After:   curse,
		Request: request,
	}

	if validation.ValidationChain(ctx, s.Validate(ctx)...) {

		return s.Call(ctx)
	}

	return nil, nil

}
