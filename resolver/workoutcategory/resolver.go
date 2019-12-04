package workoutcategory

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/calls/tagcall"
	"github.com/train-formula/graphcms/database/tagdb"

	"github.com/train-formula/graphcms/calls/organizationcall"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models/tag"
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

func (r *Resolver) TrainerOrganization(ctx context.Context, obj *workout.WorkoutCategory) (*trainer.Organization, error) {

	g := organizationcall.GetOrganization{
		ID: obj.TrainerOrganizationID,
		DB: r.db,
	}

	if validation.ValidationChain(ctx, g.Validate(ctx)...) {
		return g.Call(ctx)
	}

	return nil, nil
}

func (r *Resolver) Exercises(ctx context.Context, obj *workout.WorkoutCategory, first *int, after uuid.UUID) (*generated.ExerciseConnection, error) {
	panic("not implemented")
}

func (r *Resolver) Tags(ctx context.Context, obj *workout.WorkoutCategory) ([]*tag.Tag, error) {

	g := tagcall.GetObjectTags{
		Request: tagdb.TagsByObject{
			ObjectUUID: obj.ID,
			ObjectType: tag.WorkoutCategoryTagType,
		},
		DB: r.db,
	}

	if validation.ValidationChain(ctx, g.Validate(ctx)...) {
		return g.Call(ctx)
	}

	return nil, nil
}
