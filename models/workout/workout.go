package workout

import (
	"time"

	"github.com/gofrs/uuid"
)

type Workout struct {
	tableName             struct{}  `sql:"workout.workout"`
	ID                    uuid.UUID `json:"id"`
	CreatedAt             time.Time `json:"createdAt"`
	UpdatedAt             time.Time `json:"updatedAt"`
	TrainerOrganizationID uuid.UUID `json:"trainerOrganizationID"`
	Name                  string    `json:"name"`
	Description           string    `json:"description"`
}

func (w Workout) TableName() string {
	return "workout.workout"
}
