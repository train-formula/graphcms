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
	WorkoutProgramID      uuid.UUID `json:"workoutProgramID"`
	Name                  string    `json:"name"`
	Description           string    `json:"description"`
	DaysFromStart         int       `json:"daysFromStart"`
}

func (w Workout) TableName() string {
	return "workout.workout"
}
