package query

import (
	"context"
	"errors"

	"github.com/go-pg/pg/v9"
	uuid "github.com/satori/go.uuid"
	"github.com/train-formula/graphcms/database/organization"
	"github.com/train-formula/graphcms/models/trainer"
)

func (r *QueryResolver) Organization(ctx context.Context, id uuid.UUID) (*trainer.Organization, error) {

	org, err := organization.GetOrganization(context.Background(), r.db, id)

	if err == pg.ErrNoRows {
		return nil, errors.New("Unknown organization ID")
	} else if err != nil {
		return nil, errors.New("Failed to retrieve organization")
	}

	return &org, nil
}
