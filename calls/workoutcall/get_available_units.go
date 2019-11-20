package workoutcall

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/train-formula/graphcms/database/workoutdb"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
)

type GetAvailableUnits struct {
	DB *pg.DB
}

func (g GetAvailableUnits) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (g GetAvailableUnits) Call(ctx context.Context) ([]*workout.Unit, error) {

	return workoutdb.GetAllUnits(ctx, g.DB)
}
