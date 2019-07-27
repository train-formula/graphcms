package graphcms

import (
	"context"

	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models/workout"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Exercise() generated.ExerciseResolver {
	return &exerciseResolver{r}
}
func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Workout() generated.WorkoutResolver {
	return &workoutResolver{r}
}
func (r *Resolver) WorkoutCategory() generated.WorkoutCategoryResolver {
	return &workoutCategoryResolver{r}
}
func (r *Resolver) WorkoutProgram() generated.WorkoutProgramResolver {
	return &workoutProgramResolver{r}
}

type exerciseResolver struct{ *Resolver }

func (r *exerciseResolver) TrainerAccount(ctx context.Context, obj *workout.Exercise) (string, error) {
	panic("not implemented")
}
func (r *exerciseResolver) ClientAccount(ctx context.Context, obj *workout.Exercise) (string, error) {
	panic("not implemented")
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) Health(ctx context.Context) (*string, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Health(ctx context.Context) (*string, error) {
	panic("not implemented")
}
func (r *queryResolver) Workout(ctx context.Context, id string) (*workout.Workout, error) {
	panic("not implemented")
}

type workoutResolver struct{ *Resolver }

func (r *workoutResolver) TrainerAccount(ctx context.Context, obj *workout.Workout) (string, error) {
	panic("not implemented")
}
func (r *workoutResolver) OrderNumber(ctx context.Context, obj *workout.Workout) (int, error) {
	panic("not implemented")
}
func (r *workoutResolver) OccursOnDate(ctx context.Context, obj *workout.Workout) (*string, error) {
	panic("not implemented")
}
func (r *workoutResolver) Categories(ctx context.Context, obj *workout.Workout, first *int, after string) (*generated.WorkoutCategoryConnection, error) {
	panic("not implemented")
}

type workoutCategoryResolver struct{ *Resolver }

func (r *workoutCategoryResolver) TrainerAccount(ctx context.Context, obj *workout.WorkoutCategory) (string, error) {
	panic("not implemented")
}
func (r *workoutCategoryResolver) Workout(ctx context.Context, obj *workout.WorkoutCategory) (*workout.Workout, error) {
	panic("not implemented")
}
func (r *workoutCategoryResolver) Exercises(ctx context.Context, obj *workout.WorkoutCategory, first *int, after string) (*generated.ExerciseConnection, error) {
	panic("not implemented")
}

type workoutProgramResolver struct{ *Resolver }

func (r *workoutProgramResolver) Workouts(ctx context.Context, obj *workout.WorkoutProgram, first *int, after string) (*generated.WorkoutConnection, error) {
	panic("not implemented")
}
