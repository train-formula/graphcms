//go:generate go run github.com/vektah/dataloaden BlockExercisesByBlockLoader github.com/gofrs/uuid.UUID []*github.com/train-formula/graphcms/models/workout.BlockExercise
package blockexercisesbyblock

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database/workoutdb"
	"github.com/train-formula/graphcms/dataloader"
	"github.com/train-formula/graphcms/models/workout"
)

func NewLoader(ctx context.Context, db *pg.DB) *BlockExercisesByBlockLoader {

	return NewBlockExercisesByBlockLoader(BlockExercisesByBlockLoaderConfig{
		Wait:     dataloader.DefaultWaitTime,
		MaxBatch: dataloader.DefaultBatchSize,
		Fetch: func(keys []uuid.UUID) ([][]*workout.BlockExercise, []error) {

			blockExercises, err := workoutdb.GetBlockExercisesByBlock(ctx, db, keys)

			result := make([][]*workout.BlockExercise, len(keys))
			errors := make([]error, len(keys))

			if err != nil {
				return result, dataloader.FillErrorSlice(err, errors)
			}

			for i, k := range keys {
				if blockExercise, hasBlockExercise := blockExercises[k]; hasBlockExercise {
					result[i] = blockExercise
				} else {
					result[i] = nil
				}
			}

			return result, errors
		},
	})
}
