package workoutcategory

import (
	"context"

	"github.com/train-formula/graphcms/calls/tagcall"
	"github.com/train-formula/graphcms/calls/workoutcall"
	"github.com/train-formula/graphcms/database/tagdb"
	"github.com/willtrking/pgxload"
	"go.uber.org/zap"

	"github.com/train-formula/graphcms/calls/organizationcall"
	"github.com/train-formula/graphcms/models/tag"
	"github.com/train-formula/graphcms/models/trainer"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
)

type WorkoutCategoryResolver struct {
	db     pgxload.PgxLoader
	logger *zap.Logger
}

func NewWorkoutCategoryResolver(db pgxload.PgxLoader, logger *zap.Logger) *WorkoutCategoryResolver {
	return &WorkoutCategoryResolver{
		db:     db,
		logger: logger.Named("WorkoutCategoryResolver"),
	}
}

func (r *WorkoutCategoryResolver) TrainerOrganization(ctx context.Context, obj *workout.WorkoutCategory) (*trainer.Organization, error) {

	if obj == nil {
		return nil, nil
	}

	g := organizationcall.NewGetOrganization(obj.TrainerOrganizationID, r.logger, r.db)

	if validation.ValidationChain(ctx, g.Validate(ctx)...) {
		return g.Call(ctx)
	}

	return nil, nil
}

func (r *WorkoutCategoryResolver) WorkoutBlocks(ctx context.Context, obj *workout.WorkoutCategory) ([]*workout.WorkoutBlock, error) {
	if obj == nil {
		return nil, nil
	}

	g := workoutcall.NewGetWorkoutCategoryBlocks(obj.ID, r.logger, r.db)

	if validation.ValidationChain(ctx, g.Validate(ctx)...) {
		return g.Call(ctx)
	}

	return nil, nil
}

func (r *WorkoutCategoryResolver) Tags(ctx context.Context, obj *workout.WorkoutCategory) ([]*tag.Tag, error) {
	if obj == nil {
		return nil, nil
	}

	request := tagdb.TagsByObject{
		ObjectUUID: obj.ID,
		ObjectType: tag.WorkoutCategoryTagType,
	}

	g := tagcall.NewGetObjectTags(request, r.logger, r.db)

	if validation.ValidationChain(ctx, g.Validate(ctx)...) {
		return g.Call(ctx)
	}

	return nil, nil
}
