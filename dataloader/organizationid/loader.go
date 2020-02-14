//go:generate go run github.com/vektah/dataloaden OrganizationIDLoader github.com/gofrs/uuid.UUID *github.com/train-formula/graphcms/models/trainer.Organization

package organizationid

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database/trainerdb"
	"github.com/train-formula/graphcms/dataloader"
	"github.com/train-formula/graphcms/models/trainer"
	"github.com/willtrking/pgxload"
)

func NewLoader(ctx context.Context, db pgxload.PgxLoader) *OrganizationIDLoader {

	return NewOrganizationIDLoader(OrganizationIDLoaderConfig{
		Wait:     dataloader.DefaultWaitTime,
		MaxBatch: dataloader.DefaultBatchSize,
		Fetch: func(keys []uuid.UUID) ([]*trainer.Organization, []error) {

			orgs, err := trainerdb.GetOrganizations(ctx, db, keys)

			result := make([]*trainer.Organization, len(keys))
			errors := make([]error, len(keys))

			if err != nil {
				return result, dataloader.FillErrorSlice(err, errors)
			}

			orgsMap := make(map[uuid.UUID]*trainer.Organization)

			for _, o := range orgs {
				orgsMap[o.ID] = o
			}

			for i, k := range keys {
				if org, hasOrg := orgsMap[k]; hasOrg {
					result[i] = org
				} else {
					result[i] = nil
				}
			}

			return result, errors

		},
	})
}
