package workoutcategory

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/train-formula/graphcms/calls/tagcall"
	"github.com/train-formula/graphcms/calls/workoutcall"
	"github.com/train-formula/graphcms/database/tagdb"
	"github.com/vektah/gqlparser/gqlerror"
	"go.uber.org/zap"

	"github.com/train-formula/graphcms/calls/organizationcall"
	"github.com/train-formula/graphcms/models/tag"
	"github.com/train-formula/graphcms/models/trainer"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
)

type WorkoutCategoryResolver struct {
	db     *pg.DB
	logger *zap.Logger
}

func NewWorkoutCategoryResolver(db *pg.DB, logger *zap.Logger) *WorkoutCategoryResolver {
	return &WorkoutCategoryResolver{
		db:     db,
		logger: logger.Named("WorkoutCategoryResolver"),
	}
}

func (r *WorkoutCategoryResolver) TrainerOrganization(ctx context.Context, obj *workout.WorkoutCategory) (*trainer.Organization, error) {

	if obj == nil {
		return nil, gqlerror.Errorf("Cannot locate organization from nil workout category")
	}

	g := organizationcall.GetOrganization{
		ID: obj.TrainerOrganizationID,
		DB: r.db,
	}

	if validation.ValidationChain(ctx, g.Validate(ctx)...) {
		return g.Call(ctx)
	}

	return nil, nil
}

func (r *WorkoutCategoryResolver) WorkoutBlocks(ctx context.Context, obj *workout.WorkoutCategory) ([]*workout.WorkoutBlock, error) {
	if obj == nil {
		return nil, gqlerror.Errorf("Cannot locate workout blocks from nil workout category")
	}

	g := workoutcall.NewGetWorkoutCategoryBlocks(obj.ID, r.logger, r.db)

	if validation.ValidationChain(ctx, g.Validate(ctx)...) {
		return g.Call(ctx)
	}

	return nil, nil
}

func (r *WorkoutCategoryResolver) Tags(ctx context.Context, obj *workout.WorkoutCategory) ([]*tag.Tag, error) {
	if obj == nil {
		return nil, gqlerror.Errorf("Cannot locate tags from nil workout category")
	}

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
