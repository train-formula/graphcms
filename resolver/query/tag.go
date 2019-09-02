package query

import (
	"context"

	uuid "github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/calls/tagcall"
	"github.com/train-formula/graphcms/models/tag"
	"github.com/train-formula/graphcms/validation"
)

func (r *QueryResolver) Tag(ctx context.Context, id uuid.UUID) (*tag.Tag, error) {

	g := tagcall.GetTag{
		ID: id,
		DB: r.db,
	}

	if validation.ValidationChain(ctx, g.Validate(ctx)...) {
		return g.Call(ctx)
	}

	return nil, nil

}

func (r *QueryResolver) TagByTag(ctx context.Context, tag string, trainerOrganizationID uuid.UUID) (*tag.Tag, error) {

	g := tagcall.GetTagByTag{
		Tag:                   tag,
		TrainerOrganizationID: trainerOrganizationID,
		DB:                    r.db,
	}

	if validation.ValidationChain(ctx, g.Validate(ctx)...) {
		return g.Call(ctx)
	}

	return nil, nil

}

/*func (r *QueryResolver) Tag(ctx context.Context, id uuid.UUID) (*Tag, error) {

}*/

/*
Tag(ctx context.Context, id uuid.UUID) (*Tag, error)
	TagByTag(ctx context.Context, tag string) (*Tag, error)
*/
