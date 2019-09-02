package tagcall

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database/tagdb"
	"github.com/train-formula/graphcms/dataloader/tagbytag"
	"github.com/train-formula/graphcms/models/tag"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
)

type GetTagByTag struct {
	Tag                   string
	TrainerOrganizationID uuid.UUID
	DB                    *pg.DB
}

func (g GetTagByTag) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (g GetTagByTag) Call(ctx context.Context) (*tag.Tag, error) {
	loader := tagbytag.GetContextLoader(ctx)

	loaded, err := loader.Load(tagdb.TagByTag{
		Tag:                   g.Tag,
		TrainerOrganizationID: g.TrainerOrganizationID,
	})
	if err != nil {
		return nil, err
	}

	if loaded == nil {
		return nil, gqlerror.Errorf("Tag does not exist in organization")
	}

	return loaded, nil
}
