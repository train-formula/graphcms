package workoutdb

import (
	"context"

	"github.com/train-formula/graphcms/database"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/willtrking/pgxload"
)

// Update workout program, replace all fields with new row
func UpdateWorkoutProgram(ctx context.Context, conn pgxload.PgxTxLoader, new workout.WorkoutProgram) (*workout.WorkoutProgram, error) {

	upd := pgxload.NewStructUpdate(database.TableName(new), new)

	upd = upd.WithReturning("*")

	insStmt, insParams, err := upd.GenerateExactUpdate(conn.Mapper(), "id")
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

// Update workout category, replace all fields with new row
func UpdateWorkoutCategory(ctx context.Context, conn pgxload.PgxTxLoader, new workout.WorkoutCategory) (*workout.WorkoutCategory, error) {

	upd := pgxload.NewStructUpdate(database.TableName(new), new)

	upd = upd.WithReturning("*")

	insStmt, insParams, err := upd.GenerateExactUpdate(conn.Mapper(), "id")
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

// Update workout, replace all fields with new row
func UpdateWorkout(ctx context.Context, conn pgxload.PgxTxLoader, new workout.Workout) (*workout.Workout, error) {

	upd := pgxload.NewStructUpdate(database.TableName(new), new)

	upd = upd.WithReturning("*")

	insStmt, insParams, err := upd.GenerateExactUpdate(conn.Mapper(), "id")
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

// Update workout block, replace all fields with new row
func UpdateWorkoutBlock(ctx context.Context, conn pgxload.PgxTxLoader, new workout.WorkoutBlock) (*workout.WorkoutBlock, error) {

	upd := pgxload.NewStructUpdate(database.TableName(new), new)

	upd = upd.WithReturning("*")

	insStmt, insParams, err := upd.GenerateExactUpdate(conn.Mapper(), "id")
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

// Update exercise, replace all fields with new row
func UpdateExercise(ctx context.Context, conn pgxload.PgxTxLoader, new workout.Exercise) (*workout.Exercise, error) {

	upd := pgxload.NewStructUpdate(database.TableName(new), new)

	upd = upd.WithReturning("*")

	insStmt, insParams, err := upd.GenerateExactUpdate(conn.Mapper(), "id")
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

// Update prescription, replace all fields with new row
func UpdatePrescription(ctx context.Context, conn pgxload.PgxTxLoader, new workout.Prescription) (*workout.Prescription, error) {

	upd := pgxload.NewStructUpdate(database.TableName(new), new)

	upd = upd.WithReturning("*")

	insStmt, insParams, err := upd.GenerateExactUpdate(conn.Mapper(), "id")
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

// Update prescription set, replace all fields with new row
func UpdatePrescriptionSet(ctx context.Context, conn pgxload.PgxTxLoader, new workout.PrescriptionSet) (*workout.PrescriptionSet, error) {

	upd := pgxload.NewStructUpdate(database.TableName(new), new)

	upd = upd.WithReturning("*")

	insStmt, insParams, err := upd.GenerateExactUpdate(conn.Mapper(), "id")
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
