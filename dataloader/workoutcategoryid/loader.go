//go:generate go run github.com/vektah/dataloaden WorkoutCategoryIDLoader github.com/gofrs/uuid.UUID *github.com/train-formula/graphcms/models/workout.WorkoutCategory
package workoutcategoryid

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database/workoutdb"
	"github.com/train-formula/graphcms/dataloader"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/willtrking/pgxload"
)

func NewLoader(ctx context.Context, db pgxload.PgxLoader) *WorkoutCategoryIDLoader {

	return NewWorkoutCategoryIDLoader(WorkoutCategoryIDLoaderConfig{
		Wait:     dataloader.DefaultWaitTime,
		MaxBatch: dataloader.DefaultBatchSize,
		Fetch: func(keys []uuid.UUID) ([]*workout.WorkoutCategory, []error) {

			categories, err := workoutdb.GetWorkoutCategories(ctx, db, keys)

			result := make([]*workout.WorkoutCategory, len(keys))
			errors := make([]error, len(keys))

			if err != nil {
				return result, dataloader.FillErrorSlice(err, errors)
			}

			categoryMap := make(map[uuid.UUID]*workout.WorkoutCategory)

			for _, o := range categories {
				categoryMap[o.ID] = o
			}

			for i, k := range keys {
				if tg, hasTag := categoryMap[k]; hasTag {
					result[i] = tg
				} else {
					result[i] = nil
				}
			}

			return result, errors
		},
	})
}
