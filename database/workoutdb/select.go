package workoutdb

import (
	"context"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/willtrking/pgxload"
)

// Retrieve's all workout unit's from the database
func GetAllUnits(ctx context.Context, conn pgxload.PgxLoader) ([]*workout.Unit, error) {

	var result []*workout.Unit

	query := "SELECT * FROM " + database.TableName(workout.Unit{})

	rows, err := conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	err = conn.Scanner(rows).ScanRow(&result)

	return result, err
}

// Retrieves individual workout unit's by their IDs
func GetUnits(ctx context.Context, conn pgxload.PgxLoader, ids []uuid.UUID) ([]*workout.Unit, error) {

	if len(ids) <= 0 {
		return nil, nil
	}

	var result []*workout.Unit

	query := "SELECT * FROM " + database.TableName(workout.Unit{}) + " WHERE "

	var params []interface{}

	for _, id := range ids {
		query += "id = ? OR "
		params = append(params, id)
	}

	query = strings.TrimSuffix(query, " OR ")

	rows, err := conn.Query(ctx, pgxload.RebindPositional(query), params...)
	if err != nil {
		return nil, err
	}

	err = conn.Scanner(rows).ScanRow(&result)
	if err != nil {
		return nil, err
	}

	return result, err
}

// Retrieves an workout category by its id
func GetUnit(ctx context.Context, conn pgxload.PgxLoader, id uuid.UUID) (workout.Unit, error) {

	var result workout.Unit

	rows, err := conn.Query(ctx, pgxload.RebindPositional("SELECT * FROM "+database.TableName(result)+" WHERE id = ?"), id)
	if err != nil {
		return workout.Unit{}, err
	}

	err = conn.Scanner(rows).ScanRow(&result)

	return result, err
}

// Retrieves an workout category by its id
func GetWorkoutCategory(ctx context.Context, conn pgxload.PgxLoader, id uuid.UUID) (workout.WorkoutCategory, error) {

	var result workout.WorkoutCategory

	rows, err := conn.Query(ctx, pgxload.RebindPositional("SELECT * FROM "+database.TableName(workout.WorkoutCategory{})+" WHERE id = ?"), id)
	if err != nil {
		return workout.WorkoutCategory{}, err
	}

	err = conn.Scanner(rows).ScanRow(&result)

	return result, err
}

// Retrieves an workout category by its id, and locks the row
func GetWorkoutCategoryForUpdate(ctx context.Context, conn pgxload.PgxLoader, id uuid.UUID) (workout.WorkoutCategory, error) {

	var result workout.WorkoutCategory

	rows, err := conn.Query(ctx, pgxload.RebindPositional("SELECT * FROM "+database.TableName(result)+" WHERE id = ? FOR UPDATE"), id)
	if err != nil {
		return workout.WorkoutCategory{}, err
	}

	err = conn.Scanner(rows).ScanRow(&result)

	return result, err
}

// Retrieves individual workout categories by their IDs
func GetWorkoutCategories(ctx context.Context, conn pgxload.PgxLoader, ids []uuid.UUID) ([]*workout.WorkoutCategory, error) {

	if len(ids) <= 0 {
		return nil, nil
	}

	var result []*workout.WorkoutCategory

	query := "SELECT * FROM " + database.TableName(workout.WorkoutCategory{}) + " WHERE "

	var params []interface{}

	for _, id := range ids {
		query += "id = ? OR "
		params = append(params, id)
	}

	query = strings.TrimSuffix(query, " OR ")

	rows, err := conn.Query(ctx, pgxload.RebindPositional(query), params...)
	if err != nil {
		return nil, err
	}

	err = conn.Scanner(rows).Scan(&result)
	if err != nil {
		return nil, err
	}

	return result, err
}

// Retrieve a prescription by it's id
func GetPrescription(ctx context.Context, conn pgxload.PgxLoader, id uuid.UUID) (workout.Prescription, error) {
	var result workout.Prescription

	rows, err := conn.Query(ctx, pgxload.RebindPositional("SELECT * FROM "+database.TableName(result)+" WHERE id = ?"), id)
	if err != nil {
		return workout.Prescription{}, err
	}

	err = conn.Scanner(rows).ScanRow(&result)

	return result, err
}

// Retrieves a prescription by its id, and locks the row
func GetPrescriptionForUpdate(ctx context.Context, conn pgxload.PgxLoader, id uuid.UUID) (workout.Prescription, error) {

	var result workout.Prescription

	rows, err := conn.Query(ctx, pgxload.RebindPositional("SELECT * FROM "+database.TableName(result)+" WHERE id = ? FOR UPDATE"), id)
	if err != nil {
		return workout.Prescription{}, err
	}

	err = conn.Scanner(rows).ScanRow(&result)

	return result, err
}

// Retrieves a prescription set by its id, and locks the row
func GetPrescriptionSetForUpdate(ctx context.Context, conn pgxload.PgxLoader, id uuid.UUID) (workout.PrescriptionSet, error) {

	var result workout.PrescriptionSet

	rows, err := conn.Query(ctx, pgxload.RebindPositional("SELECT * FROM "+database.TableName(result)+" WHERE id = ? FOR UPDATE"), id)
	if err != nil {
		return workout.PrescriptionSet{}, err
	}

	err = conn.Scanner(rows).ScanRow(&result)

	return result, err
}

// Retrieves individual prescription's by their IDs
func GetPrescriptions(ctx context.Context, conn pgxload.PgxLoader, ids []uuid.UUID) ([]*workout.Prescription, error) {

	if len(ids) <= 0 {
		return nil, nil
	}

	var result []*workout.Prescription

	query := "SELECT * FROM " + database.TableName(workout.Prescription{}) + " WHERE "

	var params []interface{}

	for _, id := range ids {
		query += "id = ? OR "
		params = append(params, id)
	}

	query = strings.TrimSuffix(query, " OR ")

	rows, err := conn.Query(ctx, pgxload.RebindPositional(query), params...)
	if err != nil {
		return nil, err
	}

	err = conn.Scanner(rows).Scan(&result)

	return result, err
}

// Retrieves individual workout blocks's by their IDs
func GetWorkoutBlocks(ctx context.Context, conn pgxload.PgxLoader, ids []uuid.UUID) ([]*workout.WorkoutBlock, error) {

	if len(ids) <= 0 {
		return nil, nil
	}

	var result []*workout.WorkoutBlock

	query := "SELECT * FROM " + database.TableName(workout.WorkoutBlock{}) + " WHERE "

	var params []interface{}

	for _, id := range ids {
		query += "id = ? OR "
		params = append(params, id)
	}

	query = strings.TrimSuffix(query, " OR ")

	rows, err := conn.Query(ctx, pgxload.RebindPositional(query), params...)
	if err != nil {
		return nil, err
	}

	err = conn.Scanner(rows).Scan(&result)
	if err != nil {
		return nil, err
	}

	return result, err
}

// Retrieves workout blocks by workout category IDs
func GetWorkoutCategoryBlocks(ctx context.Context, conn pgxload.PgxLoader, workoutCategoryIDs []uuid.UUID) (map[uuid.UUID][]*workout.WorkoutBlock, error) {

	results := make(map[uuid.UUID][]*workout.WorkoutBlock)

	if len(workoutCategoryIDs) <= 0 {
		return results, nil
	}

	query := "SELECT * FROM " + database.TableName(workout.WorkoutBlock{}) + " WHERE "

	var params []interface{}

	var queryResults []*workout.WorkoutBlock

	for _, id := range workoutCategoryIDs {
		query += "workout_category_id = ? OR "
		params = append(params, id)
	}

	query = strings.TrimSuffix(query, " OR ")

	rows, err := conn.Query(ctx, pgxload.RebindPositional(query), params...)
	if err != nil {
		return nil, err
	}

	err = conn.Scanner(rows).Scan(&queryResults)
	if err != nil {
		return nil, err
	}

	for _, queryResult := range queryResults {
		results[queryResult.WorkoutCategoryID] = append(results[queryResult.WorkoutCategoryID], queryResult)
	}

	return results, err
}

// Retrieves a workout block by its id, and locks the row
func GetWorkoutBlockForUpdate(ctx context.Context, conn pgxload.PgxLoader, id uuid.UUID) (workout.WorkoutBlock, error) {

	var result workout.WorkoutBlock

	rows, err := conn.Query(ctx, pgxload.RebindPositional("SELECT * FROM "+database.TableName(result)+" WHERE id = ? FOR UPDATE"), id)
	if err != nil {
		return workout.WorkoutBlock{}, err
	}

	err = conn.Scanner(rows).Scan(&result)

	return result, err
}

// Retrieves a workout by its id, and locks the row
func GetWorkoutForUpdate(ctx context.Context, conn pgxload.PgxLoader, id uuid.UUID) (workout.Workout, error) {

	var result workout.Workout

	rows, err := conn.Query(ctx, pgxload.RebindPositional("SELECT * FROM "+database.TableName(result)+" WHERE id = ? FOR UPDATE"), id)
	if err != nil {
		return workout.Workout{}, err
	}

	err = conn.Scanner(rows).ScanRow(&result)

	return result, err
}

// Retrieves individual workouts by their IDs
func GetWorkouts(ctx context.Context, conn pgxload.PgxLoader, ids []uuid.UUID) ([]*workout.Workout, error) {
	if len(ids) <= 0 {
		return nil, nil
	}

	var result []*workout.Workout

	query := "SELECT * FROM " + database.TableName(workout.Workout{}) + " WHERE "

	var params []interface{}

	for _, id := range ids {
		query += "id = ? OR "
		params = append(params, id)
	}

	query = strings.TrimSuffix(query, " OR ")

	rows, err := conn.Query(ctx, pgxload.RebindPositional(query), params...)
	if err != nil {
		return nil, err
	}

	err = conn.Scanner(rows).Scan(&result)

	return result, err
}

// Retrieve a workout program by it's id
func GetWorkoutProgram(ctx context.Context, conn pgxload.PgxLoader, id uuid.UUID) (workout.WorkoutProgram, error) {
	var result workout.WorkoutProgram

	rows, err := conn.Query(ctx, pgxload.RebindPositional("SELECT * FROM "+database.TableName(result)+" WHERE id = ?"), id)
	if err != nil {
		return workout.WorkoutProgram{}, err
	}

	err = conn.Scanner(rows).ScanRow(&result)

	return result, err
}

// Retrieves individual workout programs by their IDs
func GetWorkoutPrograms(ctx context.Context, conn pgxload.PgxLoader, ids []uuid.UUID) ([]*workout.WorkoutProgram, error) {
	if len(ids) <= 0 {
		return nil, nil
	}

	var result []*workout.WorkoutProgram

	query := "SELECT * FROM " + database.TableName(workout.WorkoutProgram{}) + " WHERE "

	var params []interface{}

	for _, id := range ids {
		query += "id = ? OR "
		params = append(params, id)
	}

	query = strings.TrimSuffix(query, " OR ")

	rows, err := conn.Query(ctx, pgxload.RebindPositional(query), params...)
	if err != nil {
		return nil, err
	}

	err = conn.Scanner(rows).Scan(&result)
	if err != nil {
		return nil, err
	}

	return result, err
}

// Retrieves workout categories by workout IDs
func GetWorkoutCategoriesByWorkout(ctx context.Context, conn pgxload.PgxLoader, workoutIDs []uuid.UUID) (map[uuid.UUID][]*workout.WorkoutCategory, error) {

	results := make(map[uuid.UUID][]*workout.WorkoutCategory)

	if len(workoutIDs) <= 0 {
		return results, nil
	}

	query := "SELECT " + (workout.WorkoutWorkoutCategoryJoin{}).SelectColumns("wc", "wwc") + " FROM " + database.TableName(workout.WorkoutWorkoutCategory{}) + " wwc " +
		"INNER JOIN " + database.TableName(workout.WorkoutCategory{}) + " wc " +
		"ON wwc.category_id = wc.id " +
		"WHERE "

	var params []interface{}

	var queryResults []*workout.WorkoutWorkoutCategoryJoin

	for _, id := range workoutIDs {
		query += "wwc.workout_id = ? OR "
		params = append(params, id)
	}

	query = strings.TrimSuffix(query, " OR ") + " ORDER BY wwc.order ASC"

	rows, err := conn.Query(ctx, pgxload.RebindPositional(query), params...)
	if err != nil {
		return nil, err
	}

	err = conn.Scanner(rows).Scan(&queryResults)
	if err != nil {
		return nil, err
	}

	for _, queryResult := range queryResults {

		results[queryResult.WorkoutWorkoutID] = append(results[queryResult.WorkoutWorkoutID], queryResult.WorkoutCategory())
	}

	return results, err
}

// Retrieves individual exercises by their IDs
func GetExercises(ctx context.Context, conn pgxload.PgxLoader, ids []uuid.UUID) ([]*workout.Exercise, error) {
	if len(ids) <= 0 {
		return nil, nil
	}

	var result []*workout.Exercise

	query := "SELECT * FROM " + database.TableName(workout.Exercise{}) + " WHERE "

	var params []interface{}

	for _, id := range ids {
		query += "id = ? OR "
		params = append(params, id)
	}

	query = strings.TrimSuffix(query, " OR ")

	rows, err := conn.Query(ctx, pgxload.RebindPositional(query), params...)
	if err != nil {
		return nil, err
	}

	err = conn.Scanner(rows).Scan(&result)

	return result, err
}

// Retrieves an exercise by its id, and locks the row
func GetExerciseForUpdate(ctx context.Context, conn pgxload.PgxLoader, id uuid.UUID) (workout.Exercise, error) {

	var result workout.Exercise

	rows, err := conn.Query(ctx, pgxload.RebindPositional("SELECT * FROM "+database.TableName(result)+" WHERE id = ? FOR UPDATE"), id)
	if err != nil {
		return workout.Exercise{}, err
	}

	err = conn.Scanner(rows).ScanRow(&result)

	return result, err
}

// Check if an exercise is connected to any blocks
func ExerciseConnectedToBlocks(ctx context.Context, conn pgxload.PgxLoader, id uuid.UUID) (bool, error) {

	query := pgxload.RebindPositional("SELECT COUNT(1) FROM " + database.TableName(workout.BlockExercise{}) + " WHERE exercise_id = ?")

	var count int

	rows, err := conn.Query(ctx, query, id)
	if err != nil {

		return false, err
	}

	err = conn.Scanner(rows).Scan(&count)
	if err != nil {

		return false, err
	}

	return count > 0, nil

}

// Check if an prescription is connected to any blocks
func PrescriptionConnectedToBlocks(ctx context.Context, conn pgxload.PgxLoader, id uuid.UUID) (bool, error) {

	query := pgxload.RebindPositional("SELECT COUNT(1) FROM " + database.TableName(workout.BlockExercise{}) + " WHERE prescription_id = ?")

	var count int

	rows, err := conn.Query(ctx, query, id)
	if err != nil {

		return false, err
	}

	err = conn.Scanner(rows).Scan(&count)
	if err != nil {

		return false, err
	}

	return count > 0, nil

}

// Retrieves workout block exercise + prescription combinations by block id
func GetBlockExercisesByBlock(ctx context.Context, conn pgxload.PgxLoader, blockIDs []uuid.UUID) (map[uuid.UUID][]*workout.BlockExercise, error) {

	results := make(map[uuid.UUID][]*workout.BlockExercise)

	if len(blockIDs) <= 0 {
		return results, nil
	}

	query := "SELECT * FROM " + database.TableName(workout.BlockExercise{}) + " WHERE "

	var params []interface{}

	var queryResults []*workout.BlockExercise

	for _, id := range blockIDs {
		query += "block_id = ? OR "
		params = append(params, id)
	}

	query = strings.TrimSuffix(query, " OR ") + " ORDER BY order ASC"

	rows, err := conn.Query(ctx, query, params...)
	if err != nil {
		return nil, err
	}

	err = conn.Scanner(rows).Scan(&queryResults)
	if err != nil {
		return nil, err
	}

	for _, queryResult := range queryResults {

		results[queryResult.BlockID] = append(results[queryResult.BlockID], queryResult)
	}

	return results, err
}

// Retrieves PrescriptionSet's by their Prescription.
func GetPrescriptionSetsByPrescription(ctx context.Context, conn pgxload.PgxLoader, prescriptionIDs []uuid.UUID) (map[uuid.UUID][]*workout.PrescriptionSet, error) {

	results := make(map[uuid.UUID][]*workout.PrescriptionSet)

	if len(prescriptionIDs) <= 0 {
		return results, nil
	}

	query := "SELECT * FROM " + database.TableName(workout.PrescriptionSet{}) + " WHERE "

	var params []interface{}
	var queryResults []*workout.PrescriptionSet

	for _, id := range prescriptionIDs {
		query += "prescription_id OR "
		params = append(params, id)
	}

	query = strings.TrimSuffix(query, " OR ") + " ORDER BY set_number ASC"

	rows, err := conn.Query(ctx, pgxload.RebindPositional(query), params...)
	if err != nil {
		return nil, err
	}

	err = conn.Scanner(rows).Scan(&queryResults)
	if err != nil {
		return nil, err
	}

	for _, queryResult := range queryResults {

		results[queryResult.PrescriptionID] = append(results[queryResult.PrescriptionID], queryResult)
	}

	return results, err

}
