package workoutdb

import (
	"context"

	"github.com/train-formula/graphcms/database"
	"github.com/train-formula/graphcms/models/workout"
)

func InsertWorkoutProgram(ctx context.Context, conn database.Conn, new workout.WorkoutProgram) (*workout.WorkoutProgram, error) {

	newModel := &new

	_, err := conn.ModelContext(ctx, newModel).Returning("*").Insert()
	if err != nil {
		return nil, err
	}

	return newModel, nil
}
