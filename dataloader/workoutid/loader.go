//go:generate go run github.com/vektah/dataloaden WorkoutIDLoader github.com/gofrs/uuid.UUID *github.com/train-formula/graphcms/models/workout.Workout
package workoutid

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database/workoutdb"
	"github.com/train-formula/graphcms/dataloader"
	"github.com/train-formula/graphcms/models/workout"
)

func NewLoader(ctx context.Context, db *pg.DB) *WorkoutIDLoader {

	return NewWorkoutIDLoader(WorkoutIDLoaderConfig{
		Wait:     dataloader.DefaultWaitTime,
		MaxBatch: dataloader.DefaultBatchSize,
		Fetch: func(keys []uuid.UUID) ([]*workout.Workout, []error) {

			workouts, err := workoutdb.GetWorkouts(ctx, db, keys)

			result := make([]*workout.Workout, len(keys))
			errors := make([]error, len(keys))

			if err != nil {
				return result, dataloader.FillErrorSlice(err, errors)
			}

			workoutMap := make(map[uuid.UUID]*workout.Workout)

			for _, o := range workouts {
				workoutMap[o.ID] = o
			}

			for i, k := range keys {
				if tg, hasWorkout := workoutMap[k]; hasWorkout {
					result[i] = tg
				} else {
					result[i] = nil
				}
			}

			return result, errors
		},
	})
}
