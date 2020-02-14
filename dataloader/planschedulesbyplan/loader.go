//go:generate go run github.com/vektah/dataloaden PlanSchedulesByPlanLoader github.com/gofrs/uuid.UUID []*github.com/train-formula/graphcms/models/plan.PlanSchedule
package planschedulesbyplan

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database/plandb"
	"github.com/train-formula/graphcms/dataloader"
	"github.com/train-formula/graphcms/models/plan"
	"github.com/willtrking/pgxload"
)

func NewLoader(ctx context.Context, db pgxload.PgxLoader) *PlanSchedulesByPlanLoader {

	return NewPlanSchedulesByPlanLoader(PlanSchedulesByPlanLoaderConfig{
		Wait:     dataloader.DefaultWaitTime,
		MaxBatch: dataloader.DefaultBatchSize,
		Fetch: func(keys []uuid.UUID) ([][]*plan.PlanSchedule, []error) {

			planSchedules, err := plandb.GetSchedulesForPlans(ctx, db, keys)

			result := make([][]*plan.PlanSchedule, len(keys))
			errors := make([]error, len(keys))

			if err != nil {
				return result, dataloader.FillErrorSlice(err, errors)
			}

			for i, k := range keys {
				if planSchedule, hasPlanSchedule := planSchedules[k]; hasPlanSchedule {
					result[i] = planSchedule
				} else {
					result[i] = nil
				}
			}

			return result, errors
		},
	})
}
