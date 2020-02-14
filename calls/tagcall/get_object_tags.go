package tagcall

import (
	"context"

	"github.com/train-formula/graphcms/database/tagdb"
	"github.com/train-formula/graphcms/dataloader/tagsbyobject"
	"github.com/train-formula/graphcms/logging"
	"github.com/train-formula/graphcms/models/tag"
	"github.com/train-formula/graphcms/validation"
	"github.com/willtrking/pgxload"
	"go.uber.org/zap"
)

func NewGetObjectTags(request tagdb.TagsByObject, logger *zap.Logger, db pgxload.PgxLoader) *GetObjectTags {
	return &GetObjectTags{
		request: request,
		db:      db,
		logger:  logger.Named("GetObjectTags"),
	}
}

type GetObjectTags struct {
	request tagdb.TagsByObject
	db      pgxload.PgxLoader
	logger  *zap.Logger
}

func (g GetObjectTags) Validate(ctx context.Context) []validation.ValidatorFunc {

	return []validation.ValidatorFunc{
		validation.CheckTagTypeKnown(g.request.ObjectType),
	}
}

func (g GetObjectTags) Call(ctx context.Context) ([]*tag.Tag, error) {
	loader := tagsbyobject.GetContextLoader(ctx)

	loaded, err := loader.Load(g.request)
	if err != nil {
		g.logger.Error("Failed to load tags for object", zap.Error(err),
			logging.UUID("objectUUID", g.request.ObjectUUID),
			zap.String("objectType", g.request.ObjectType.String()))
		return nil, err
	}

	return loaded, nil
}
