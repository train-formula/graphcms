package workout

import (
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database"
)

type WorkoutWorkoutCategoryJoin struct {
	WorkoutCategoryID                    uuid.UUID
	WorkoutCategoryCreatedAt             time.Time
	WorkoutCategoryUpdatedAt             time.Time
	WorkoutCategoryTrainerOrganizationID uuid.UUID
	WorkoutCategoryName                  string
	WorkoutCategoryDescription           string

	// WorkoutWorkoutID from WorkoutWorkoutCategory (workout.workout_category) table, NOT ID field, hence WorkoutWorkoutID
	WorkoutWorkoutID uuid.UUID
}

// Extract a Tag struct from this result
func (t *WorkoutWorkoutCategoryJoin) WorkoutCategory() *WorkoutCategory {
	return &WorkoutCategory{
		ID:                    t.WorkoutCategoryID,
		CreatedAt:             t.WorkoutCategoryCreatedAt,
		UpdatedAt:             t.WorkoutCategoryUpdatedAt,
		TrainerOrganizationID: t.WorkoutCategoryTrainerOrganizationID,
		Name:                  t.WorkoutCategoryName,
		Description:           t.WorkoutCategoryDescription,
	}
}

// Extract the columns to use in a SELECT statement
func (t WorkoutWorkoutCategoryJoin) SelectColumns(workoutCategoryTablePrefix, workoutWorkoutCategoryID string) string {

	valued := database.ReflectValue(WorkoutCategory{})

	columns := database.StructColumns(valued, workoutCategoryTablePrefix)
	columns = append(columns, database.PGPrefixedColumn("workout_id", workoutWorkoutCategoryID))

	return strings.Join(columns, ",")
}
