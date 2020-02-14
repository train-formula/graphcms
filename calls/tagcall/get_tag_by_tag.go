package tagcall

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database/tagdb"
	"github.com/train-formula/graphcms/dataloader/tagbytag"
	"github.com/train-formula/graphcms/logging"
	"github.com/train-formula/graphcms/models/tag"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
	"github.com/willtrking/pgxload"
	"go.uber.org/zap"
)

func NewGetTagByTag(tag string, trainerOrganizationID uuid.UUID, logger *zap.Logger, db pgxload.PgxLoader) *GetTagByTag {
	return &GetTagByTag{
		tag:                   tag,
		trainerOrganizationID: trainerOrganizationID,
		db:                    db,
		logger:                logger.Named("GetTagByTag"),
	}
}

type GetTagByTag struct {
	tag                   string
	trainerOrganizationID uuid.UUID
	db                    pgxload.PgxLoader
	logger                *zap.Logger
}

func (g GetTagByTag) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (g GetTagByTag) Call(ctx context.Context) (*tag.Tag, error) {
	loader := tagbytag.GetContextLoader(ctx)

	loaded, err := loader.Load(tagdb.TagByTag{
		Tag:                   g.tag,
		TrainerOrganizationID: g.trainerOrganizationID,
	})
	if err != nil {
		g.logger.Error("Failed to load tag with dataloader", zap.Error(err),
			zap.String("tag", g.tag),
			logging.UUID("trainerOrganizationID", g.trainerOrganizationID))
		return nil, err
	}

	if loaded == nil {
		return nil, gqlerror.Errorf("tag does not exist in organization")
	}

	return loaded, nil
}
