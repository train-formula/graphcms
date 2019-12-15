//go:generate go run github.com/vektah/dataloaden WorkoutBlockIDLoader github.com/gofrs/uuid.UUID *github.com/train-formula/graphcms/models/workout.WorkoutBlock
package workoutblockid

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database/workoutdb"
	"github.com/train-formula/graphcms/dataloader"
	"github.com/train-formula/graphcms/models/workout"
)

func NewLoader(ctx context.Context, db *pg.DB) *WorkoutBlockIDLoader {

	return NewWorkoutBlockIDLoader(WorkoutBlockIDLoaderConfig{
		Wait:     dataloader.DefaultWaitTime,
		MaxBatch: dataloader.DefaultBatchSize,
		Fetch: func(keys []uuid.UUID) ([]*workout.WorkoutBlock, []error) {

			workoutBlocks, err := workoutdb.GetWorkoutBlocks(ctx, db, keys)

			result := make([]*workout.WorkoutBlock, len(keys))
			errors := make([]error, len(keys))

			if err != nil {
				return result, dataloader.FillErrorSlice(err, errors)
			}

			workoutBlockMap := make(map[uuid.UUID]*workout.WorkoutBlock)

			for _, o := range workoutBlocks {
				workoutBlockMap[o.ID] = o
			}

			for i, k := range keys {
				if tg, hasWorkoutBlock := workoutBlockMap[k]; hasWorkoutBlock {
					result[i] = tg
				} else {
					result[i] = nil
				}
			}

			return result, errors
		},
	})
}
