//go:generate go run github.com/vektah/dataloaden ExerciseIDLoader github.com/gofrs/uuid.UUID *github.com/train-formula/graphcms/models/workout.Exercise
package exerciseid

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database/workoutdb"
	"github.com/train-formula/graphcms/dataloader"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/willtrking/pgxload"
)

func NewLoader(ctx context.Context, db pgxload.PgxLoader) *ExerciseIDLoader {

	return NewExerciseIDLoader(ExerciseIDLoaderConfig{
		Wait:     dataloader.DefaultWaitTime,
		MaxBatch: dataloader.DefaultBatchSize,
		Fetch: func(keys []uuid.UUID) ([]*workout.Exercise, []error) {

			exercises, err := workoutdb.GetExercises(ctx, db, keys)

			result := make([]*workout.Exercise, len(keys))
			errors := make([]error, len(keys))

			if err != nil {
				return result, dataloader.FillErrorSlice(err, errors)
			}

			exerciseMap := make(map[uuid.UUID]*workout.Exercise)

			for _, o := range exercises {
				exerciseMap[o.ID] = o
			}

			for i, k := range keys {
				if tg, hasExercise := exerciseMap[k]; hasExercise {
					result[i] = tg
				} else {
					result[i] = nil
				}
			}

			return result, errors
		},
	})
}
