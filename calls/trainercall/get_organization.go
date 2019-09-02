package trainercall

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/dataloader/organizationid"
	"github.com/train-formula/graphcms/models/trainer"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
)

type GetOrganization struct {
	ID uuid.UUID
	DB *pg.DB
}

func (g GetOrganization) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (g GetOrganization) Call(ctx context.Context) (*trainer.Organization, error) {
	loader := organizationid.GetContextLoader(ctx)

	loaded, err := loader.Load(g.ID)
	if err != nil {
		return nil, err
	}

	if loaded == nil {
		return nil, gqlerror.Errorf("Unknown organization ID")
	}

	return loaded, nil
}
