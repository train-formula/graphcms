//go:generate go run github.com/vektah/dataloaden PlanScheduleIDLoader github.com/gofrs/uuid.UUID *github.com/train-formula/graphcms/models/plan.PlanSchedule
package planscheduleid

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database/plandb"
	"github.com/train-formula/graphcms/dataloader"
	"github.com/train-formula/graphcms/models/plan"
)

func NewLoader(ctx context.Context, db *pg.DB) *PlanScheduleIDLoader {

	return NewPlanScheduleIDLoader(PlanScheduleIDLoaderConfig{
		Wait:     dataloader.DefaultWaitTime,
		MaxBatch: dataloader.DefaultBatchSize,
		Fetch: func(keys []uuid.UUID) ([]*plan.PlanSchedule, []error) {

			schedules, err := plandb.GetPlanSchedules(ctx, db, keys)

			result := make([]*plan.PlanSchedule, len(keys))
			errors := make([]error, len(keys))

			if err != nil {
				return result, dataloader.FillErrorSlice(err, errors)
			}

			scheduleMap := make(map[uuid.UUID]*plan.PlanSchedule)

			for _, o := range schedules {
				scheduleMap[o.ID] = o
			}

			for i, k := range keys {
				if tg, hasSchedule := scheduleMap[k]; hasSchedule {
					result[i] = tg
				} else {
					result[i] = nil
				}
			}

			return result, errors
		},
	})
}
