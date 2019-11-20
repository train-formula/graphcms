package query

import (
	"context"

	"github.com/train-formula/graphcms/calls/workoutcall"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
)

func (r *QueryResolver) AvailableUnits(ctx context.Context) ([]*workout.Unit, error) {

	g := workoutcall.GetAvailableUnits{
		DB: r.db,
	}

	if validation.ValidationChain(ctx, g.Validate(ctx)...) {
		return g.Call(ctx)
	}

	return nil, nil
}
