//go:generate go run github.com/vektah/dataloaden PrescriptionSetsByPrescriptionLoader github.com/gofrs/uuid.UUID []*github.com/train-formula/graphcms/models/workout.PrescriptionSet
package prescriptionsetsbyprescription

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database/workoutdb"
	"github.com/train-formula/graphcms/dataloader"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/willtrking/pgxload"
)

func NewLoader(ctx context.Context, db pgxload.PgxLoader) *PrescriptionSetsByPrescriptionLoader {

	return NewPrescriptionSetsByPrescriptionLoader(PrescriptionSetsByPrescriptionLoaderConfig{
		Wait:     dataloader.DefaultWaitTime,
		MaxBatch: dataloader.DefaultBatchSize,
		Fetch: func(keys []uuid.UUID) ([][]*workout.PrescriptionSet, []error) {

			prescriptionSets, err := workoutdb.GetPrescriptionSetsByPrescription(ctx, db, keys)

			result := make([][]*workout.PrescriptionSet, len(keys))
			errors := make([]error, len(keys))

			if err != nil {
				return result, dataloader.FillErrorSlice(err, errors)
			}

			for i, k := range keys {
				if prescriptionSet, hasPrescriptionSet := prescriptionSets[k]; hasPrescriptionSet {
					result[i] = prescriptionSet
				} else {
					result[i] = nil
				}
			}

			return result, errors
		},
	})
}
