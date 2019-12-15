package workoutdb

import (
	"context"

	"github.com/train-formula/graphcms/database"
	"github.com/train-formula/graphcms/models/workout"
)

// Update workout program, replace all fields with new row
func UpdateWorkoutProgram(ctx context.Context, conn database.Conn, new workout.WorkoutProgram) (*workout.WorkoutProgram, error) {

	newModel := &new

	_, err := conn.ModelContext(ctx, newModel).Returning("*").Update()
	if err != nil {
		return nil, err
	}

	return newModel, nil
}

// Update workout category, replace all fields with new row
func UpdateWorkoutCategory(ctx context.Context, conn database.Conn, new workout.WorkoutCategory) (*workout.WorkoutCategory, error) {

	newModel := &new

	_, err := conn.ModelContext(ctx, newModel).Returning("*").Update()
	if err != nil {
		return nil, err
	}

	return newModel, nil
}

// Update workout, replace all fields with new row
func UpdateWorkout(ctx context.Context, conn database.Conn, new workout.Workout) (*workout.Workout, error) {

	newModel := &new

	_, err := conn.ModelContext(ctx, newModel).Returning("*").Update()
	if err != nil {
		return nil, err
	}

	return newModel, nil
}
