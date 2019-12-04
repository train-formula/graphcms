package tagcall

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/train-formula/graphcms/database/tagdb"
	"github.com/train-formula/graphcms/dataloader/tagsbyobject"
	"github.com/train-formula/graphcms/models/tag"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
)

type GetObjectTags struct {
	Request tagdb.TagsByObject
	DB      *pg.DB
}

func (g GetObjectTags) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (g GetObjectTags) Call(ctx context.Context) ([]*tag.Tag, error) {
	loader := tagsbyobject.GetContextLoader(ctx)

	loaded, err := loader.Load(g.Request)
	if err != nil {
		return nil, err
	}

	if loaded == nil {
		return nil, gqlerror.Errorf("Failed to load tags by object")
	}

	return loaded, nil
}
