package mutation

import (
	"context"

	"github.com/train-formula/graphcms/calls/tagcall"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models/tag"
	"github.com/train-formula/graphcms/validation"
)

func (m *MutationResolver) CreateTag(ctx context.Context, request generated.CreateTag) (*tag.Tag, error) {

	call := tagcall.NewCreateTag(request, m.logger, m.db)

	if validation.ValidationChain(ctx, call.Validate(ctx)...) {
		return call.Call(ctx)
	}

	return nil, nil

}
