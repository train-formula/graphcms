package query

import (
	"context"

	uuid "github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/calls/tagcall"
	"github.com/train-formula/graphcms/models/tag"
	"github.com/train-formula/graphcms/validation"
)

func (r *QueryResolver) Tag(ctx context.Context, id uuid.UUID) (*tag.Tag, error) {

	g := tagcall.NewGetTag(id, r.logger, r.db)

	if validation.ValidationChain(ctx, g.Validate(ctx)...) {
		return g.Call(ctx)
	}

	return nil, nil

}

func (r *QueryResolver) TagByTag(ctx context.Context, tag string, trainerOrganizationID uuid.UUID) (*tag.Tag, error) {

	g := tagcall.NewGetTagByTag(tag, trainerOrganizationID, r.logger, r.db)

	if validation.ValidationChain(ctx, g.Validate(ctx)...) {
		return g.Call(ctx)
	}

	return nil, nil

}

/*func (r *QueryResolver) tag(ctx context.Context, id uuid.UUID) (*tag, error) {

}*/

/*
tag(ctx context.Context, id uuid.UUID) (*tag, error)
	TagByTag(ctx context.Context, tag string) (*tag, error)
*/
