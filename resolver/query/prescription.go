package query

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/calls/workoutcall"
	"github.com/train-formula/graphcms/database/cursor"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"go.uber.org/zap"
)

func (r *QueryResolver) Prescription(ctx context.Context, id uuid.UUID) (*workout.Prescription, error) {

	g := workoutcall.NewGetPrescription(id, r.logger, r.db)

	if validation.ValidationChain(ctx, g.Validate(ctx)...) {
		return g.Call(ctx)
	}

	return nil, nil
}

func (r *QueryResolver) PrescriptionSearch(ctx context.Context, request generated.PrescriptionSearchRequest, first int, after *string) (*generated.PrescriptionSearchResults, error) {

	curse, err := cursor.DeserializeCursor(after)
	if err != nil {
		r.logger.Error("Failed to deserialize cursor", zap.Error(err))
		return nil, err
	}

	s := workoutcall.NewSearchPrescriptions(request, first, curse, r.logger, r.db)

	if validation.ValidationChain(ctx, s.Validate(ctx)...) {

		return s.Call(ctx)
	}

	return nil, nil
}
