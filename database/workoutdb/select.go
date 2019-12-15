package workoutdb

import (
	"context"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database"
	"github.com/train-formula/graphcms/models/workout"
)

// Retrieve's all workout unit's from the database
func GetAllUnits(ctx context.Context, conn database.Conn) ([]*workout.Unit, error) {

	var result []*workout.Unit

	query := "SELECT * FROM " + database.TableName(workout.Unit{})

	_, err := conn.QueryContext(ctx, &result, query)

	return result, err
}

// Retrieves individual workout unit's by their IDs
func GetUnits(ctx context.Context, conn database.Conn, ids []uuid.UUID) ([]*workout.Unit, error) {

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

	_, err := conn.QueryContext(ctx, &result, query, params...)

	return result, err
}

// Retrieves an workout category by its ID
func GetWorkoutCategory(ctx context.Context, conn database.Conn, id uuid.UUID) (workout.WorkoutCategory, error) {

	var result workout.WorkoutCategory

	_, err := conn.QueryOneContext(ctx, &result, "SELECT * FROM "+database.TableName(result)+" WHERE id = ?", id)

	return result, err
}

// Retrieves an workout category by its ID, and locks the row
func GetWorkoutCategoryForUpdate(ctx context.Context, conn database.Conn, id uuid.UUID) (workout.WorkoutCategory, error) {

	var result workout.WorkoutCategory

	_, err := conn.QueryOneContext(ctx, &result, "SELECT * FROM "+database.TableName(result)+" WHERE id = ? FOR UPDATE", id)

	return result, err
}

// Retrieves individual workout categories by their IDs
func GetWorkoutCategories(ctx context.Context, conn database.Conn, ids []uuid.UUID) ([]*workout.WorkoutCategory, error) {

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

	_, err := conn.QueryContext(ctx, &result, query, params...)

	return result, err
}

// Retrieves individual prescription's by their IDs
func GetPrescriptions(ctx context.Context, conn database.Conn, ids []uuid.UUID) ([]*workout.Prescription, error) {

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

	_, err := conn.QueryContext(ctx, &result, query, params...)

	return result, err
}

// Retrieves individual workout blocks's by their IDs
func GetWorkoutBlocks(ctx context.Context, conn database.Conn, ids []uuid.UUID) ([]*workout.WorkoutBlock, error) {

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

	_, err := conn.QueryContext(ctx, &result, query, params...)

	return result, err
}

// Retrieves workout blocks by workout category IDs
func GetWorkoutCategoryBlocks(ctx context.Context, conn database.Conn, workoutCategoryIDs []uuid.UUID) (map[uuid.UUID][]*workout.WorkoutBlock, error) {

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

	_, err := conn.QueryContext(ctx, &queryResults, query, params...)

	if err != nil {
		return nil, err
	}

	for _, queryResult := range queryResults {
		results[queryResult.WorkoutCategoryID] = append(results[queryResult.WorkoutCategoryID], queryResult)
	}

	return results, err
}

// Retrieves a workout by its ID, and locks the row
func GetWorkoutForUpdate(ctx context.Context, conn database.Conn, id uuid.UUID) (workout.Workout, error) {

	var result workout.Workout

	_, err := conn.QueryOneContext(ctx, &result, "SELECT * FROM "+database.TableName(result)+" WHERE id = ? FOR UPDATE", id)

	return result, err
}

// Retrieves individual workouts by their IDs
func GetWorkouts(ctx context.Context, conn database.Conn, ids []uuid.UUID) ([]*workout.Workout, error) {
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

	_, err := conn.QueryContext(ctx, &result, query, params...)

	return result, err
}

// Retrieve a workout program by it's ID
func GetWorkoutProgram(ctx context.Context, conn database.Conn, id uuid.UUID) (workout.WorkoutProgram, error) {
	var result workout.WorkoutProgram

	_, err := conn.QueryOneContext(ctx, &result, "SELECT * FROM "+database.TableName(result)+" WHERE id = ?", id)

	return result, err
}

// Retrieves workout categories by workout IDs
func GetWorkoutCategoriesByWorkout(ctx context.Context, conn database.Conn, workoutIDs []uuid.UUID) (map[uuid.UUID][]*workout.WorkoutCategory, error) {

	results := make(map[uuid.UUID][]*workout.WorkoutCategory)

	if len(workoutIDs) <= 0 {
		return results, nil
	}

	query := "SELECT " + (workout.WorkoutWorkoutCategoryJoin{}).SelectColumns("wc", "wwc") + " FROM " + database.TableName(workout.WorkoutWorkoutCategory{}) + " wwc " +
		"INNER JOIN " + database.TableName(workout.WorkoutCategory{}) + " wc " +
		"ON wwc.category_id = wc.id" +
		"WHERE "

	var params []interface{}

	var queryResults []*workout.WorkoutWorkoutCategoryJoin

	for _, id := range workoutIDs {
		query += "wwc.workout_id = ? OR "
		params = append(params, id)
	}

	query = strings.TrimSuffix(query, " OR ") + " ORDER BY wwc.order ASC"

	_, err := conn.QueryContext(ctx, &queryResults, query, params...)

	if err != nil {
		return nil, err
	}

	for _, queryResult := range queryResults {
		results[queryResult.WorkoutWorkoutID] = append(results[queryResult.WorkoutCategoryID], queryResult.WorkoutCategory())
	}

	return results, err
}
