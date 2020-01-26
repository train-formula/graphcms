package exercise

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/train-formula/graphcms/calls/tagcall"
	"github.com/train-formula/graphcms/database/tagdb"
	"github.com/train-formula/graphcms/models/tag"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"go.uber.org/zap"
)

type ExerciseResolver struct {
	db     *pg.DB
	logger *zap.Logger
}

func NewExerciseResolver(db *pg.DB, logger *zap.Logger) *ExerciseResolver {
	return &ExerciseResolver{
		db:     db,
		logger: logger.Named("ExerciseResolver"),
	}
}

func (r *ExerciseResolver) Tags(ctx context.Context, obj *workout.Exercise) ([]*tag.Tag, error) {
	if obj == nil {
		return nil, nil
	}

	request := tagdb.TagsByObject{
		ObjectUUID: obj.ID,
		ObjectType: tag.ExerciseTagType,
	}

	g := tagcall.NewGetObjectTags(request, r.logger, r.db)

	if validation.ValidationChain(ctx, g.Validate(ctx)...) {
		return g.Call(ctx)
	}

	return nil, nil
}
