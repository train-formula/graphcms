package workout

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

type WorkoutProgram struct {
	ID                    uuid.UUID `json:"id"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
	Name                  string `json:"name"`
	Description           string `json:"description"`
	Public                bool
	Price                 string
	TrainerOrganizationID uuid.UUID `json:"trainerOrganizationID"`
}

func (w WorkoutProgram) TableName() string {
	return "workout.program"
}
