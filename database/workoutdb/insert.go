package workoutdb

import (
	"context"
	"errors"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database"
	"github.com/train-formula/graphcms/models/workout"
)

// Insert a new workout program
func InsertWorkoutProgram(ctx context.Context, conn database.Conn, new workout.WorkoutProgram) (*workout.WorkoutProgram, error) {

	newModel := &new

	_, err := conn.ModelContext(ctx, newModel).Returning("*").Insert()
	if err != nil {
		return nil, err
	}

	return newModel, nil
}

// Insert a new workout category
func InsertWorkoutCategory(ctx context.Context, conn database.Conn, new workout.WorkoutCategory) (*workout.WorkoutCategory, error) {

	newModel := &new

	_, err := conn.ModelContext(ctx, newModel).Returning("*").Insert()
	if err != nil {
		return nil, err
	}

	return newModel, nil
}

// Insert a new workout
func InsertWorkout(ctx context.Context, conn database.Conn, new workout.Workout) (*workout.Workout, error) {

	newModel := &new

	_, err := conn.ModelContext(ctx, newModel).Returning("*").Insert()
	if err != nil {
		return nil, err
	}

	return newModel, nil
}

// Remove all workout categories from a workout, and set a new list
// The order of the categories will be the order they appear in the list
func SetWorkoutWorkoutCategories(ctx context.Context, conn database.Conn, workoutID uuid.UUID, workoutCategoryIDs []uuid.UUID) error {

	if len(workoutCategoryIDs) <= 0 {
		return errors.New("must specify at least one workout category ID to set onto workout")
	}

	err := ClearWorkoutWorkoutCategories(ctx, conn, workoutID)
	if err != nil {
		return err
	}

	var toInsert []*workout.WorkoutWorkoutCategory

	for idx, categoryID := range workoutCategoryIDs {
		newUuid, err := uuid.NewV4()
		if err != nil {
			return err
		}

		toInsert = append(toInsert, &workout.WorkoutWorkoutCategory{
			ID:         newUuid,
			WorkoutID:  workoutID,
			CategoryID: categoryID,
			Order:      idx,
		})
	}

	_, err = conn.ModelContext(ctx, &toInsert).Insert()
	if err != nil {
		return err
	}

	return nil

}
