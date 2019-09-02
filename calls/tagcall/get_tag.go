package tagcall

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/models/tag"

	"github.com/train-formula/graphcms/dataloader/tagid"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
)

type GetTag struct {
	ID uuid.UUID
	DB *pg.DB
}

func (g GetTag) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (g GetTag) Call(ctx context.Context) (*tag.Tag, error) {
	loader := tagid.GetContextLoader(ctx)

	loaded, err := loader.Load(g.ID)
	if err != nil {
		return nil, err
	}

	if loaded == nil {
		return nil, gqlerror.Errorf("Unknown tag ID")
	}

	return loaded, nil
}
