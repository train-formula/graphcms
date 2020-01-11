package query

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/calls/organizationcall"
	"github.com/train-formula/graphcms/database/cursor"
	"github.com/train-formula/graphcms/models/connections"
	"github.com/train-formula/graphcms/models/trainer"
	"github.com/train-formula/graphcms/validation"
	"go.uber.org/zap"
)

func (r *QueryResolver) Organization(ctx context.Context, id uuid.UUID) (*trainer.Organization, error) {

	g := organizationcall.NewGetOrganization(id, r.logger, r.db)

	if validation.ValidationChain(ctx, g.Validate(ctx)...) {
		return g.Call(ctx)
	}

	return nil, nil
}

func (r *QueryResolver) OrganizationAvailableTags(ctx context.Context, id uuid.UUID, first int, after *string) (*connections.TagConnection, error) {

	cursor, err := cursor.DeserializeCursor(after)
	if err != nil {
		r.logger.Error("Failed to deserialize cursor", zap.Error(err))
		return nil, err
	}

	g := organizationcall.NewGetOrganizationAvailableTags(id, first, cursor, r.logger, r.db)

	if validation.ValidationChain(ctx, g.Validate(ctx)...) {
		return g.Call(ctx)
	}

	return nil, nil
}
