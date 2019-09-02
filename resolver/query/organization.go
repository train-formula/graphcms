package query

import (
	"context"

	uuid "github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/calls/trainercall"
	"github.com/train-formula/graphcms/models/trainer"
	"github.com/train-formula/graphcms/validation"
)

func (r *QueryResolver) Organization(ctx context.Context, id uuid.UUID) (*trainer.Organization, error) {

	g := trainercall.GetOrganization{
		ID: id,
		DB: r.db,
	}

	if validation.ValidationChain(ctx, g.Validate(ctx)...) {
		return g.Call(ctx)
	}

	return nil, nil
}
