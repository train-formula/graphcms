package organizationcall

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/dataloader/organizationid"
	"github.com/train-formula/graphcms/logging"
	"github.com/train-formula/graphcms/models/trainer"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
	"go.uber.org/zap"
)

func NewGetOrganization(id uuid.UUID, logger *zap.Logger, db *pg.DB) *GetOrganization {
	return &GetOrganization{
		id:     id,
		db:     db,
		logger: logger.Named("GetOrganization"),
	}
}

type GetOrganization struct {
	id     uuid.UUID
	db     *pg.DB
	logger *zap.Logger
}

func (g GetOrganization) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (g GetOrganization) Call(ctx context.Context) (*trainer.Organization, error) {
	loader := organizationid.GetContextLoader(ctx)

	loaded, err := loader.Load(g.id)
	if err != nil {
		g.logger.Error("Failed to load organization with dataloader", zap.Error(err),
			logging.UUID("organizationID", g.id))
		return nil, err
	}

	if loaded == nil {
		return nil, gqlerror.Errorf("Unknown organization id")
	}

	return loaded, nil
}
