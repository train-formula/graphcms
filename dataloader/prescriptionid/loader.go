//go:generate go run github.com/vektah/dataloaden PrescriptionIDLoader github.com/gofrs/uuid.UUID *github.com/train-formula/graphcms/models/workout.Prescription
package prescriptionid

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database/workoutdb"
	"github.com/train-formula/graphcms/dataloader"
	"github.com/train-formula/graphcms/models/workout"
)

func NewLoader(ctx context.Context, db *pg.DB) *PrescriptionIDLoader {

	return NewPrescriptionIDLoader(PrescriptionIDLoaderConfig{
		Wait:     dataloader.DefaultWaitTime,
		MaxBatch: dataloader.DefaultBatchSize,
		Fetch: func(keys []uuid.UUID) ([]*workout.Prescription, []error) {

			prescriptions, err := workoutdb.GetPrescriptions(ctx, db, keys)

			result := make([]*workout.Prescription, len(keys))
			errors := make([]error, len(keys))

			if err != nil {
				return result, dataloader.FillErrorSlice(err, errors)
			}

			prescriptionMap := make(map[uuid.UUID]*workout.Prescription)

			for _, o := range prescriptions {
				prescriptionMap[o.ID] = o
			}

			for i, k := range keys {
				if tg, hasTag := prescriptionMap[k]; hasTag {
					result[i] = tg
				} else {
					result[i] = nil
				}
			}

			return result, errors
		},
	})
}
