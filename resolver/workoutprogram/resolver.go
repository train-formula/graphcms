package workoutprogram

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/calls/trainercall"

	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models/trainer"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
)

func NewResolver(db *pg.DB) *Resolver {
	return &Resolver{
		db: db,
	}
}

type Resolver struct {
	db *pg.DB
}

func (r *Resolver) TrainerOrganization(ctx context.Context, obj *workout.WorkoutProgram) (*trainer.Organization, error) {

	g := trainercall.GetOrganization{
		ID: obj.TrainerOrganizationID,
		DB: r.db,
	}

	if validation.ValidationChain(ctx, g.Validate(ctx)...) {
		return g.Call(ctx)
	}

	return nil, nil
}

func (r *Resolver) Workouts(ctx context.Context, obj *workout.WorkoutProgram, first *int, after uuid.UUID) (*generated.WorkoutConnection, error) {

	return nil, nil
}
