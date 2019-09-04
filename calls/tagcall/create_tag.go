package tagcall

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database/tagdb"

	"github.com/train-formula/graphcms/database/trainerdb"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models/tag"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
)

type CreateTag struct {
	Request generated.CreateTag
	DB      *pg.DB
}

func (g CreateTag) Validate(ctx context.Context) []validation.ValidatorFunc {

	return []validation.ValidatorFunc{
		validation.CheckStringIsNotEmpty(g.Request.Tag, "Tag must not be empty"),
	}
}

func (g CreateTag) Call(ctx context.Context) (*tag.Tag, error) {

	newUuid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	_, err = trainerdb.GetOrganization(ctx, g.DB, g.Request.TrainerOrganizationID)

	if err != nil {
		if err == pg.ErrNoRows {
			return nil, gqlerror.Errorf("Organization does not exist")
		}
		return nil, err
	}

	_, errTag := tagdb.GetTagByTag(ctx, g.DB, tagdb.TagByTag{
		Tag:                   g.Request.Tag,
		TrainerOrganizationID: g.Request.TrainerOrganizationID,
	})

	// Retrieving means tag already exists
	// If its ErrNoRows no tag exists
	if errTag == nil {
		return nil, gqlerror.Errorf("Tag '" + g.Request.Tag + "' already exists")

	} else if errTag != nil && errTag != pg.ErrNoRows {
		return nil, errTag
	}

	new := tag.Tag{
		ID:                    newUuid,
		Tag:                   g.Request.Tag,
		TrainerOrganizationID: g.Request.TrainerOrganizationID,
	}

	return tagdb.InsertTag(ctx, g.DB, new)
}
