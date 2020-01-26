package tagcall

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database/tagdb"
	"go.uber.org/zap"

	"github.com/train-formula/graphcms/database/trainerdb"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models/tag"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
)

func NewCreateTag(request generated.CreateTag, logger *zap.Logger, db *pg.DB) *CreateTag {
	return &CreateTag{
		request: request,
		db:      db,
		logger:  logger.Named("CreateTag"),
	}
}

type CreateTag struct {
	request generated.CreateTag
	db      *pg.DB
	logger  *zap.Logger
}

func (g CreateTag) Validate(ctx context.Context) []validation.ValidatorFunc {

	return []validation.ValidatorFunc{
		validation.CheckStringIsNotEmpty(g.request.Tag, "tag must not be empty", true),
	}
}

func (g CreateTag) Call(ctx context.Context) (*tag.Tag, error) {

	newUuid, err := uuid.NewV4()
	if err != nil {
		g.logger.Error("Failed to generate UUID", zap.Error(err))
		return nil, err
	}

	var finalTag *tag.Tag

	err = g.db.RunInTransaction(func(t *pg.Tx) error {
		_, err = trainerdb.GetOrganization(ctx, t, g.request.TrainerOrganizationID)

		if err != nil {
			if err == pg.ErrNoRows {
				return gqlerror.Errorf("Organization does not exist")
			}
			g.logger.Error("Failed to get organization", zap.Error(err))
			return err
		}

		_, err = tagdb.GetTagByTag(ctx, t, tagdb.TagByTag{
			Tag:                   g.request.Tag,
			TrainerOrganizationID: g.request.TrainerOrganizationID,
		})

		// Retrieving means tag already exists
		// If its ErrNoRows no tag exists
		if err == nil {
			return gqlerror.Errorf("tag '" + g.request.Tag + "' already exists")

		} else if err != nil && err != pg.ErrNoRows {
			g.logger.Error("Failed to check if tag already exists", zap.Error(err))
			return err
		}

		newTag := tag.Tag{
			ID:                    newUuid,
			Tag:                   g.request.Tag,
			TrainerOrganizationID: g.request.TrainerOrganizationID,
		}

		finalTag, err = tagdb.InsertTag(ctx, t, newTag)

		if err != nil {
			g.logger.Error("Failed to insert tag", zap.Error(err))
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return finalTag, nil
}
