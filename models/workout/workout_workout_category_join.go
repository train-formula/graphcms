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

	var columns []string
	columns = append(columns, database.PGPrefixedColumn("id", workoutCategoryTablePrefix)+" AS workout_category_id")
	columns = append(columns, database.PGPrefixedColumn("created_at", workoutCategoryTablePrefix)+" AS workout_category_created_at")
	columns = append(columns, database.PGPrefixedColumn("updated_at", workoutCategoryTablePrefix)+" AS workout_category_updated_at")
	columns = append(columns, database.PGPrefixedColumn("trainer_organization_id", workoutCategoryTablePrefix)+" AS workout_category_trainer_organization_id")
	columns = append(columns, database.PGPrefixedColumn("name", workoutCategoryTablePrefix)+" AS workout_category_name")
	columns = append(columns, database.PGPrefixedColumn("description", workoutCategoryTablePrefix)+" AS workout_category_description")
	columns = append(columns, database.PGPrefixedColumn("workout_id", workoutWorkoutCategoryID)+" AS workout_workout_id")

	return strings.Join(columns, ",")
}
