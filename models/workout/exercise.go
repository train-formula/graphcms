package workout

import (
	"time"

	"github.com/gofrs/uuid"
)

type Exercise struct {
	tableName             struct{}  `sql:"workout.exercise"`
	ID                    uuid.UUID `json:"id"`
	CreatedAt             time.Time `json:"createdAt"`
	UpdatedAt             time.Time `json:"updatedAt"`
	TrainerOrganizationID uuid.UUID `json:"trainerOrganizationID"`
	Name                  string    `json:"name"`
	ExerciseText          string    `json:"exerciseText"`
	PrescriptionID        uuid.UUID `json:"prescriptionID"`
}
