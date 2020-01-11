package workoutprogram

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/calls/organizationcall"
	"github.com/train-formula/graphcms/calls/tagcall"
	"github.com/train-formula/graphcms/database/tagdb"
	"github.com/train-formula/graphcms/models/tag"
	"github.com/vektah/gqlparser/gqlerror"
	"go.uber.org/zap"

	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models/trainer"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
)

func NewWorkoutProgramResolver(db *pg.DB, logger *zap.Logger) *WorkoutProgramResolver {
	return &WorkoutProgramResolver{
		db:     db,
		logger: logger.Named("WorkoutProgramResolver"),
	}
}

type WorkoutProgramResolver struct {
	db     *pg.DB
	logger *zap.Logger
}

func (r *WorkoutProgramResolver) TrainerOrganization(ctx context.Context, obj *workout.WorkoutProgram) (*trainer.Organization, error) {

	g := organizationcall.GetOrganization{
		ID: obj.TrainerOrganizationID,
		DB: r.db,
	}

	if validation.ValidationChain(ctx, g.Validate(ctx)...) {
		return g.Call(ctx)
	}

	return nil, nil
}

func (r *WorkoutProgramResolver) Workouts(ctx context.Context, obj *workout.WorkoutProgram, first *int, after uuid.UUID) (*generated.WorkoutConnection, error) {

	return nil, nil
}

func (r *WorkoutProgramResolver) Tags(ctx context.Context, obj *workout.WorkoutProgram) ([]*tag.Tag, error) {

	if obj == nil {
		return nil, gqlerror.Errorf("Cannot locate tags from nil workout program")
	}

	request := tagdb.TagsByObject{
		ObjectUUID: obj.ID,
		ObjectType: tag.WorkoutProgramTagType,
	}

	g := tagcall.NewGetObjectTags(request, r.logger, r.db)

	if validation.ValidationChain(ctx, g.Validate(ctx)...) {
		return g.Call(ctx)
	}

	return nil, nil
}
