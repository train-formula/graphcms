package workoutdb

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database"
	"github.com/train-formula/graphcms/models/workout"
)

// Delete a workout, and un-link all of the workout categories associated
func DeleteWorkout(ctx context.Context, conn database.Tx, workoutID uuid.UUID) error {

	err := ClearWorkoutWorkoutCategories(ctx, conn, workoutID)
	if err != nil {
		return err
	}

	_, err = conn.ExecContext(ctx, "DELETE FROM "+database.TableName(workout.Workout{})+" WHERE id = ?", workoutID)

	if err != nil {
		return err
	}

	return nil
}

// Remove all workout categories from a workout
func ClearWorkoutWorkoutCategories(ctx context.Context, conn database.Tx, workoutID uuid.UUID) error {
	_, err := conn.ExecContext(ctx, "DELETE FROM "+database.TableName(workout.WorkoutWorkoutCategory{})+" WHERE workout_id = ?", workoutID)
	if err != nil {
		return err
	}

	return nil
}

// Delete a workout block, and un-link all of the exercises associated
func DeleteWorkoutBlock(ctx context.Context, conn database.Tx, workoutBlockID uuid.UUID) error {

	_, err := conn.ExecContext(ctx, "DELETE FROM "+database.TableName(workout.WorkoutBlock{})+" WHERE id = ?", workoutBlockID)

	if err != nil {
		return err
	}

	return nil
}

// Delete an exercise
func DeleteExercise(ctx context.Context, conn database.Tx, exerciseID uuid.UUID) error {

	_, err := conn.ExecContext(ctx, "DELETE FROM "+database.TableName(workout.Exercise{})+" WHERE id = ?", exerciseID)

	if err != nil {
		return err
	}

	return nil
}

// Remove all block exercises from a workout block
func ClearWorkoutBlockBlockExercises(ctx context.Context, conn database.Tx, workoutBlockID uuid.UUID) error {
	_, err := conn.ExecContext(ctx, "DELETE FROM "+database.TableName(workout.BlockExercise{})+" WHERE block_id = ?", workoutBlockID)
	if err != nil {
		return err
	}

	return nil
}

// Delete all prescription sets that are attached to a prescription
func DeleteAllSetsFromPrescription(ctx context.Context, conn database.Tx, prescriptionID uuid.UUID) error {

	_, err := conn.ExecContext(ctx, "DELETE FROM "+database.TableName(workout.PrescriptionSet{})+" WHERE prescription_id = ?", prescriptionID)

	if err != nil {
		return err
	}

	return nil

}

// Delete a prescription and all attached prescription sets
func DeletePrescription(ctx context.Context, conn database.Tx, prescriptionID uuid.UUID) error {

	err := DeleteAllSetsFromPrescription(ctx, conn, prescriptionID)

	if err != nil {
		return err
	}

	_, err = conn.ExecContext(ctx, "DELETE FROM "+database.TableName(workout.Prescription{})+" WHERE id = ?", prescriptionID)

	if err != nil {
		return err
	}

	return nil
}

// Delete a prescription set
func DeletePrescriptionSet(ctx context.Context, conn database.Tx, prescriptionSetID uuid.UUID) error {

	_, err := conn.ExecContext(ctx, "DELETE FROM "+database.TableName(workout.PrescriptionSet{})+" WHERE id = ?", prescriptionSetID)

	if err != nil {
		return err
	}

	return nil
}
