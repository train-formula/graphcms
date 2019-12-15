//go:generate go run github.com/vektah/dataloaden WorkoutBlocksByCategoryLoader github.com/gofrs/uuid.UUID []*github.com/train-formula/graphcms/models/workout.WorkoutBlock
package workoutblocksbycategory

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database/workoutdb"
	"github.com/train-formula/graphcms/dataloader"
	"github.com/train-formula/graphcms/models/workout"
)

func NewLoader(ctx context.Context, db *pg.DB) *WorkoutBlocksByCategoryLoader {

	return NewWorkoutBlocksByCategoryLoader(WorkoutBlocksByCategoryLoaderConfig{
		Wait:     dataloader.DefaultWaitTime,
		MaxBatch: dataloader.DefaultBatchSize,
		Fetch: func(keys []uuid.UUID) ([][]*workout.WorkoutBlock, []error) {

			blocks, err := workoutdb.GetWorkoutCategoryBlocks(ctx, db, keys)

			result := make([][]*workout.WorkoutBlock, len(keys))
			errors := make([]error, len(keys))

			if err != nil {
				return result, dataloader.FillErrorSlice(err, errors)
			}

			for i, k := range keys {
				if blk, hasBlk := blocks[k]; hasBlk {
					result[i] = blk
				} else {
					result[i] = nil
				}
			}

			return result, errors
		},
	})
}
