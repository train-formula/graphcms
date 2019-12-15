package workout

import (
	"time"

	"github.com/gofrs/uuid"
)

// Links a Workout to a Workout Category
type WorkoutWorkoutCategory struct {
	tableName  struct{} `sql:"workout.workout_category"`
	ID         uuid.UUID
	CreatedAt  time.Time
	UpdatedAt  time.Time
	WorkoutID  uuid.UUID
	CategoryID uuid.UUID
	Order      int
}

func (w WorkoutWorkoutCategory) TableName() string {
	return "workout.workout_category"
}
