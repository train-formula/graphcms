package prescription

import (
	"context"

	"github.com/train-formula/graphcms/calls/tagcall"
	"github.com/train-formula/graphcms/calls/workoutcall"
	"github.com/train-formula/graphcms/database/tagdb"
	"github.com/train-formula/graphcms/models/tag"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"github.com/willtrking/pgxload"
	"go.uber.org/zap"
)

type PrescriptionResolver struct {
	db     pgxload.PgxLoader
	logger *zap.Logger
}

func NewPrescriptionResolver(db pgxload.PgxLoader, logger *zap.Logger) *PrescriptionResolver {
	return &PrescriptionResolver{
		db:     db,
		logger: logger.Named("PrescriptionResolver"),
	}
}

func (r *PrescriptionResolver) HasReps(ctx context.Context, obj *workout.Prescription) (bool, error) {
	return true, nil
}

func (r *PrescriptionResolver) HasRepModifier(ctx context.Context, obj *workout.Prescription) (bool, error) {
	return true, nil
}

func (r *PrescriptionResolver) Sets(ctx context.Context, obj *workout.Prescription) ([]*workout.PrescriptionSet, error) {

	if obj == nil {
		return nil, nil
	}

	call := workoutcall.NewGetPrescriptionPrescriptionSets(obj.ID, r.logger, r.db)

	if validation.ValidationChain(ctx, call.Validate(ctx)...) {
		return call.Call(ctx)
	}

	return nil, nil

}

func (r *PrescriptionResolver) Tags(ctx context.Context, obj *workout.Prescription) ([]*tag.Tag, error) {
	if obj == nil {
		return nil, nil
	}

	request := tagdb.TagsByObject{
		ObjectUUID: obj.ID,
		ObjectType: tag.PrescriptionTagType,
	}

	g := tagcall.NewGetObjectTags(request, r.logger, r.db)

	if validation.ValidationChain(ctx, g.Validate(ctx)...) {
		return g.Call(ctx)
	}

	return nil, nil
}
