package workoutdb

import (
	"context"
	"errors"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/willtrking/pgxload"
)

// Combined exercise + prescription id used, for example, to create BlockExercises
type ExercisePrescription struct {
	ExerciseID     uuid.UUID
	PrescriptionID uuid.UUID
}

// Insert a new workout program
func InsertWorkoutProgram(ctx context.Context, conn pgxload.PgxTxLoader, new workout.WorkoutProgram) (*workout.WorkoutProgram, error) {

	ins := pgxload.NewStructInsert(database.TableName(new), new)

	ins = ins.WithReturning("*")

	insStmt, insParams, err := ins.GenerateInsert(conn.Mapper())
	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, insStmt, insParams...)
	if err != nil {
		return nil, err
	}

	var result workout.WorkoutProgram

	err = conn.Scanner(rows).Scan(&result)

	return &result, err
}

// Insert a new workout category
func InsertWorkoutCategory(ctx context.Context, conn pgxload.PgxTxLoader, new workout.WorkoutCategory) (*workout.WorkoutCategory, error) {

	ins := pgxload.NewStructInsert(database.TableName(new), new)

	ins = ins.WithReturning("*")

	insStmt, insParams, err := ins.GenerateInsert(conn.Mapper())
	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, insStmt, insParams...)
	if err != nil {
		return nil, err
	}

	var result workout.WorkoutCategory

	err = conn.Scanner(rows).Scan(&result)

	return &result, err
}

// Insert a new workout
func InsertWorkout(ctx context.Context, conn pgxload.PgxTxLoader, new workout.Workout) (*workout.Workout, error) {

	ins := pgxload.NewStructInsert(database.TableName(new), new)

	ins = ins.WithReturning("*")

	insStmt, insParams, err := ins.GenerateInsert(conn.Mapper())
	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, insStmt, insParams...)
	if err != nil {
		return nil, err
	}

	var result workout.Workout

	err = conn.Scanner(rows).Scan(&result)

	return &result, err
}

// Insert a new workout block
func InsertWorkoutBlock(ctx context.Context, conn pgxload.PgxTxLoader, new workout.WorkoutBlock) (*workout.WorkoutBlock, error) {

	ins := pgxload.NewStructInsert(database.TableName(new), new)

	ins = ins.WithReturning("*")

	insStmt, insParams, err := ins.GenerateInsert(conn.Mapper())
	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, insStmt, insParams...)
	if err != nil {
		return nil, err
	}

	var result workout.WorkoutBlock

	err = conn.Scanner(rows).Scan(&result)

	return &result, err
}

// Insert a new exercise
func InsertExercise(ctx context.Context, conn pgxload.PgxTxLoader, new workout.Exercise) (*workout.Exercise, error) {

	ins := pgxload.NewStructInsert(database.TableName(new), new)

	ins = ins.WithReturning("*")

	insStmt, insParams, err := ins.GenerateInsert(conn.Mapper())
	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, insStmt, insParams...)
	if err != nil {
		return nil, err
	}

	var result workout.Exercise

	err = conn.Scanner(rows).Scan(&result)

	return &result, err
}

// Insert a new prescription
func InsertPrescription(ctx context.Context, conn pgxload.PgxTxLoader, new workout.Prescription) (*workout.Prescription, error) {

	ins := pgxload.NewStructInsert(database.TableName(new), new)

	ins = ins.WithReturning("*")

	insStmt, insParams, err := ins.GenerateInsert(conn.Mapper())
	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, insStmt, insParams...)
	if err != nil {
		return nil, err
	}

	var result workout.Prescription

	err = conn.Scanner(rows).Scan(&result)

	return &result, err
}

// Insert a new prescription set
func InsertPrescriptionSet(ctx context.Context, conn pgxload.PgxTxLoader, new workout.PrescriptionSet) (*workout.PrescriptionSet, error) {

	ins := pgxload.NewStructInsert(database.TableName(new), new)

	ins = ins.WithReturning("*")

	insStmt, insParams, err := ins.GenerateInsert(conn.Mapper())
	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, insStmt, insParams...)
	if err != nil {
		return nil, err
	}

	var result workout.PrescriptionSet

	err = conn.Scanner(rows).Scan(&result)

	return &result, err
}

// Remove all workout categories from a workout, and set a new list
// The order of the categories will be the order they appear in the list
func SetWorkoutWorkoutCategories(ctx context.Context, conn pgxload.PgxTxLoader, workoutID uuid.UUID, workoutCategoryIDs []uuid.UUID) error {

	if len(workoutCategoryIDs) <= 0 {
		return errors.New("must specify at least one workout category id to set onto workout")
	}

	err := ClearWorkoutWorkoutCategories(ctx, conn, workoutID)
	if err != nil {
		return err
	}

	var toInsert []interface{}

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

	ins := pgxload.NewStructInsert(database.TableName(workout.WorkoutWorkoutCategory{}), toInsert...)

	insStmt, insParams, err := ins.GenerateInsert(conn.Mapper())
	if err != nil {
		return err
	}

	_, err = conn.Exec(ctx, insStmt, insParams...)
	if err != nil {
		return err
	}

	return nil

}

// Remove all BlockExercise's from a workout block, and set a new list
// The order of the BlockExercises will be the order they appear in the list
func SetWorkoutBlockBlockExercises(ctx context.Context, conn pgxload.PgxTxLoader, workoutBlockID uuid.UUID, exercisePrescriptions []ExercisePrescription) error {

	if len(exercisePrescriptions) <= 0 {
		return errors.New("must specify at least exercise + prescription id to set onto workout block")
	}

	err := ClearWorkoutBlockBlockExercises(ctx, conn, workoutBlockID)
	if err != nil {
		return err
	}

	var toInsert []interface{}

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

	ins := pgxload.NewStructInsert(database.TableName(workout.BlockExercise{}), toInsert...)

	insStmt, insParams, err := ins.GenerateInsert(conn.Mapper())
	if err != nil {
		return err
	}

	_, err = conn.Exec(ctx, insStmt, insParams...)
	if err != nil {
		return err
	}

	return nil

}
