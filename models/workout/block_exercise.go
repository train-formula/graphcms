package workout

import (
	"time"

	"github.com/gofrs/uuid"
)

type BlockExercise struct {
	tableName      struct{}  `sql:"workout.block_exercise"`
	ID             uuid.UUID `json:"id"`
	CreatedAt      time.Time `json:"createdAt" pgxload:"defaultZero"`
	UpdatedAt      time.Time `json:"updatedAt" pgxload:"defaultZero"`
	BlockID        uuid.UUID `json:"blockID"`
	ExerciseID     uuid.UUID `json:"exerciseID"`
	PrescriptionID uuid.UUID `json:"prescriptionID"`
	Order          int       `json:"order"`
}

func (b BlockExercise) TableName() string {
	return "workout.block_exercise"
}
