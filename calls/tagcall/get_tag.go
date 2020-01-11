package tagcall

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/logging"
	"github.com/train-formula/graphcms/models/tag"
	"go.uber.org/zap"

	"github.com/train-formula/graphcms/dataloader/tagid"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
)

func NewGetTag(id uuid.UUID, logger *zap.Logger, db *pg.DB) *GetTag {
	return &GetTag{
		id:     id,
		db:     db,
		logger: logger.Named("GetTag"),
	}
}

type GetTag struct {
	id     uuid.UUID
	db     *pg.DB
	logger *zap.Logger
}

func (g GetTag) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (g GetTag) Call(ctx context.Context) (*tag.Tag, error) {
	loader := tagid.GetContextLoader(ctx)

	loaded, err := loader.Load(g.id)
	if err != nil {
		g.logger.Error("Failed to retrieve tag with dataloader", zap.Error(err), logging.UUID("tagID", g.id))
		return nil, err
	}

	if loaded == nil {
		return nil, gqlerror.Errorf("Unknown tag ID")
	}

	return loaded, nil
}
