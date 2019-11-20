package workoutdb

import (
	"context"

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
