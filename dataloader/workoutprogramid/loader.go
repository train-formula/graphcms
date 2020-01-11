//go:generate go run github.com/vektah/dataloaden WorkoutProgramIDLoader github.com/gofrs/uuid.UUID *github.com/train-formula/graphcms/models/workout.WorkoutProgram
package workoutprogramid

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database/workoutdb"
	"github.com/train-formula/graphcms/dataloader"
	"github.com/train-formula/graphcms/models/workout"
)

func NewLoader(ctx context.Context, db *pg.DB) *WorkoutProgramIDLoader {

	return NewWorkoutProgramIDLoader(WorkoutProgramIDLoaderConfig{
		Wait:     dataloader.DefaultWaitTime,
		MaxBatch: dataloader.DefaultBatchSize,
		Fetch: func(keys []uuid.UUID) ([]*workout.WorkoutProgram, []error) {

			workoutPrograms, err := workoutdb.GetWorkoutPrograms(ctx, db, keys)

			result := make([]*workout.WorkoutProgram, len(keys))
			errors := make([]error, len(keys))

			if err != nil {
				return result, dataloader.FillErrorSlice(err, errors)
			}

			workoutProgramMap := make(map[uuid.UUID]*workout.WorkoutProgram)

			for _, o := range workoutPrograms {
				workoutProgramMap[o.ID] = o
			}

			for i, k := range keys {
				if tg, hasWorkoutProgram := workoutProgramMap[k]; hasWorkoutProgram {
					result[i] = tg
				} else {
					result[i] = nil
				}
			}

			return result, errors
		},
	})
}
