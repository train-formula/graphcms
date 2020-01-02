package workoutdb

import (
	"context"
	"errors"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database"
	"github.com/train-formula/graphcms/models/workout"
)

// Combined exercise + prescription ID used, for example, to create BlockExercises
type ExercisePrescription struct {
	ExerciseID     uuid.UUID
	PrescriptionID uuid.UUID
}

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

// Insert a new workout block
func InsertWorkoutBlock(ctx context.Context, conn database.Conn, new workout.WorkoutBlock) (*workout.WorkoutBlock, error) {

	newModel := &new

	_, err := conn.ModelContext(ctx, newModel).Returning("*").Insert()
	if err != nil {
		return nil, err
	}

	return newModel, nil
}

// Insert a new exercise
func InsertExercise(ctx context.Context, conn database.Conn, new workout.Exercise) (*workout.Exercise, error) {

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

// Remove all BlockExercise's from a workout block, and set a new list
// The order of the BlockExercises will be the order they appear in the list
func SetWorkoutBlockBlockExercises(ctx context.Context, conn database.Conn, workoutBlockID uuid.UUID, exercisePrescriptions []ExercisePrescription) error {

	if len(exercisePrescriptions) <= 0 {
		return errors.New("must specify at least exercise + prescription ID to set onto workout block")
	}

	err := ClearWorkoutBlockBlockExercises(ctx, conn, workoutBlockID)
	if err != nil {
		return err
	}

	var toInsert []*workout.BlockExercise

	for idx, exercisePrescription := range exercisePrescriptions {

		newUuid, err := uuid.NewV4()
		if err != nil {
			return err
		}

		toInsert = append(toInsert, &workout.BlockExercise{
			ID:             newUuid,
			BlockID:        workoutBlockID,
			ExerciseID:     exercisePrescription.ExerciseID,
			PrescriptionID: exercisePrescription.PrescriptionID,
			Order:          idx,
		})
	}
	_, err = conn.ModelContext(ctx, &toInsert).Insert()
	if err != nil {
		return err
	}

	return nil

}
