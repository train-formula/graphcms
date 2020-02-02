package query

import (
	"context"

	uuid "github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/calls/tagcall"
	"github.com/train-formula/graphcms/database/cursor"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models/tag"
	"github.com/train-formula/graphcms/validation"
	"go.uber.org/zap"
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

func (r *QueryResolver) TagSearch(ctx context.Context, request generated.TagSearchRequest, first int, after *string) (*generated.TagSearchResults, error) {

	curse, err := cursor.DeserializeCursor(after)
	if err != nil {
		r.logger.Error("Failed to deserialize cursor", zap.Error(err))
		return nil, err
	}

	s := tagcall.NewSearchTags(request, first, curse, r.logger, r.db)

	if validation.ValidationChain(ctx, s.Validate(ctx)...) {

		return s.Call(ctx)
	}

	return nil, nil
}
