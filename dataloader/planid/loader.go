//go:generate go run github.com/vektah/dataloaden PlanIDLoader github.com/gofrs/uuid.UUID *github.com/train-formula/graphcms/models/plan.Plan
package planid

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database/plandb"
	"github.com/train-formula/graphcms/dataloader"
	"github.com/train-formula/graphcms/models/plan"
)

func NewLoader(ctx context.Context, db *pg.DB) *PlanIDLoader {

	return NewPlanIDLoader(PlanIDLoaderConfig{
		Wait:     dataloader.DefaultWaitTime,
		MaxBatch: dataloader.DefaultBatchSize,
		Fetch: func(keys []uuid.UUID) ([]*plan.Plan, []error) {

			plans, err := plandb.GetPlans(ctx, db, keys)

			result := make([]*plan.Plan, len(keys))
			errors := make([]error, len(keys))

			if err != nil {
				return result, dataloader.FillErrorSlice(err, errors)
			}

			planMap := make(map[uuid.UUID]*plan.Plan)

			for _, o := range plans {
				planMap[o.ID] = o
			}

			for i, k := range keys {
				if tg, hasPlan := planMap[k]; hasPlan {
					result[i] = tg
				} else {
					result[i] = nil
				}
			}

			return result, errors
		},
	})
}
