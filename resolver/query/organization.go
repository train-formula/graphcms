package query

import (
	"context"

	"github.com/go-pg/pg/v9"
	uuid "github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database/organization"
	"github.com/train-formula/graphcms/models/trainer"
	"github.com/vektah/gqlparser/gqlerror"
	"go.uber.org/zap"
)

func (r *QueryResolver) Organization(ctx context.Context, id uuid.UUID) (*trainer.Organization, error) {

	org, err := organization.GetOrganization(ctx, r.db, id)

	if err == pg.ErrNoRows {
		return nil, gqlerror.Errorf("Unknown organization ID")
	} else if err != nil {
		zap.L().Error("Failed to retrieve organization", zap.Error(err))
		return nil, gqlerror.Errorf("Failed to retrieve organization")
	}

	return &org, nil
}
