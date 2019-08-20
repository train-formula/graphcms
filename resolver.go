package graphcms

import (
	"context"

	"github.com/go-pg/pg/v9"
	uuid "github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models"
	"github.com/train-formula/graphcms/models/connections"
	"github.com/train-formula/graphcms/models/trainer"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/resolver/query"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{ DB *pg.DB }

func (r *Resolver) Exercise() generated.ExerciseResolver {
	return &exerciseResolver{r}
}
func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() generated.QueryResolver {
	return query.NewQueryResolver(r.DB)
}
func (r *Resolver) Unit() generated.UnitResolver {
	return &unitResolver{r}
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

func (r *Resolver) WorkoutProgramConnection() generated.WorkoutProgramConnectionResolver {
	return &connections.WorkoutProgramConnection{}
}

type exerciseResolver struct{ *Resolver }

func (r *exerciseResolver) ID(ctx context.Context, obj *workout.Exercise) (uuid.UUID, error) {
	panic("not implemented")
}
func (r *exerciseResolver) TrainerAccount(ctx context.Context, obj *workout.Exercise) (uuid.UUID, error) {
	panic("not implemented")
}
func (r *exerciseResolver) ClientAccount(ctx context.Context, obj *workout.Exercise) (uuid.UUID, error) {
	panic("not implemented")
}
func (r *exerciseResolver) CategoryID(ctx context.Context, obj *workout.Exercise) (uuid.UUID, error) {
	panic("not implemented")
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) Health(ctx context.Context) (string, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Health(ctx context.Context) (string, error) {
	panic("not implemented")
}
func (r *queryResolver) Organization(ctx context.Context, id uuid.UUID) (*trainer.Organization, error) {
	panic("not implemented")
}

type unitResolver struct{ *Resolver }

func (r *unitResolver) ID(ctx context.Context, obj *models.Unit) (uuid.UUID, error) {
	panic("not implemented")
}

type workoutResolver struct{ *Resolver }

func (r *workoutResolver) ID(ctx context.Context, obj *workout.Workout) (uuid.UUID, error) {
	panic("not implemented")
}
func (r *workoutResolver) TrainerAccount(ctx context.Context, obj *workout.Workout) (uuid.UUID, error) {
	panic("not implemented")
}
func (r *workoutResolver) OrderNumber(ctx context.Context, obj *workout.Workout) (int, error) {
	panic("not implemented")
}
func (r *workoutResolver) OccursOnDate(ctx context.Context, obj *workout.Workout) (*string, error) {
	panic("not implemented")
}
func (r *workoutResolver) Categories(ctx context.Context, obj *workout.Workout, first *int, after uuid.UUID) (*generated.WorkoutCategoryConnection, error) {
	panic("not implemented")
}

type workoutCategoryResolver struct{ *Resolver }

func (r *workoutCategoryResolver) ID(ctx context.Context, obj *workout.WorkoutCategory) (uuid.UUID, error) {
	panic("not implemented")
}
func (r *workoutCategoryResolver) TrainerAccount(ctx context.Context, obj *workout.WorkoutCategory) (uuid.UUID, error) {
	panic("not implemented")
}
func (r *workoutCategoryResolver) WorkoutID(ctx context.Context, obj *workout.WorkoutCategory) (uuid.UUID, error) {
	panic("not implemented")
}
func (r *workoutCategoryResolver) Workout(ctx context.Context, obj *workout.WorkoutCategory) (*workout.Workout, error) {
	panic("not implemented")
}
func (r *workoutCategoryResolver) Exercises(ctx context.Context, obj *workout.WorkoutCategory, first *int, after uuid.UUID) (*generated.ExerciseConnection, error) {
	panic("not implemented")
}

type workoutProgramResolver struct{ *Resolver }

func (r *workoutProgramResolver) ID(ctx context.Context, obj *workout.WorkoutProgram) (uuid.UUID, error) {
	panic("not implemented")
}
func (r *workoutProgramResolver) TrainerAccount(ctx context.Context, obj *workout.WorkoutProgram) (uuid.UUID, error) {
	panic("not implemented")
}
func (r *workoutProgramResolver) Workouts(ctx context.Context, obj *workout.WorkoutProgram, first *int, after uuid.UUID) (*generated.WorkoutConnection, error) {
	panic("not implemented")
}
