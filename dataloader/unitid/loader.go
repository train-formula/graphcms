//go:generate go run github.com/vektah/dataloaden UnitIDLoader github.com/gofrs/uuid.UUID *github.com/train-formula/graphcms/models/workout.Unit
package unitid

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database/workoutdb"
	"github.com/train-formula/graphcms/dataloader"
	"github.com/train-formula/graphcms/models/workout"
)

func NewLoader(ctx context.Context, db *pg.DB) *UnitIDLoader {

	return NewUnitIDLoader(UnitIDLoaderConfig{
		Wait:     dataloader.DefaultWaitTime,
		MaxBatch: dataloader.DefaultBatchSize,
		Fetch: func(keys []uuid.UUID) ([]*workout.Unit, []error) {

			categories, err := workoutdb.GetUnits(ctx, db, keys)

			result := make([]*workout.Unit, len(keys))
			errors := make([]error, len(keys))

			if err != nil {
				return result, dataloader.FillErrorSlice(err, errors)
			}

			unitMap := make(map[uuid.UUID]*workout.Unit)

			for _, o := range categories {
				unitMap[o.ID] = o
			}

			for i, k := range keys {
				if tg, hasTag := unitMap[k]; hasTag {
					result[i] = tg
				} else {
					result[i] = nil
				}
			}

			return result, errors
		},
	})
}
