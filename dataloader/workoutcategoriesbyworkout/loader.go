//go:generate go run github.com/vektah/dataloaden WorkoutCategoriesByWorkoutLoader github.com/gofrs/uuid.UUID []*github.com/train-formula/graphcms/models/workout.WorkoutCategory
package workoutcategoriesbyworkout

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database/workoutdb"
	"github.com/train-formula/graphcms/dataloader"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/willtrking/pgxload"
)

func NewLoader(ctx context.Context, db pgxload.PgxLoader) *WorkoutCategoriesByWorkoutLoader {

	return NewWorkoutCategoriesByWorkoutLoader(WorkoutCategoriesByWorkoutLoaderConfig{
		Wait:     dataloader.DefaultWaitTime,
		MaxBatch: dataloader.DefaultBatchSize,
		Fetch: func(keys []uuid.UUID) ([][]*workout.WorkoutCategory, []error) {

			categories, err := workoutdb.GetWorkoutCategoriesByWorkout(ctx, db, keys)

			result := make([][]*workout.WorkoutCategory, len(keys))
			errors := make([]error, len(keys))

			if err != nil {
				return result, dataloader.FillErrorSlice(err, errors)
			}

			for i, k := range keys {
				if category, hasCategory := categories[k]; hasCategory {
					result[i] = category
				} else {
					result[i] = nil
				}
			}

			return result, errors
		},
	})
}
