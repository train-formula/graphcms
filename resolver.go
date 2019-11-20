package graphcms

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models/connections"
	"github.com/train-formula/graphcms/models/tag"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/resolver/mutation"
	"github.com/train-formula/graphcms/resolver/query"
	"github.com/train-formula/graphcms/resolver/workoutcategory"
	"github.com/train-formula/graphcms/resolver/workoutprogram"
	"go.uber.org/zap"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{ DB *pg.DB }

func (r *Resolver) Exercise() generated.ExerciseResolver {
	return &exerciseResolver{r}
}
func (r *Resolver) Mutation() generated.MutationResolver {
	return mutation.NewMutationResolver(r.DB, zap.L())
}
func (r *Resolver) Query() generated.QueryResolver {
	return query.NewQueryResolver(r.DB, zap.L())
}
func (r *Resolver) Workout() generated.WorkoutResolver {
	return &workoutResolver{r}
}
func (r *Resolver) WorkoutCategory() generated.WorkoutCategoryResolver {
	return workoutcategory.NewResolver(r.DB)
}
func (r *Resolver) WorkoutProgram() generated.WorkoutProgramResolver {
	return workoutprogram.NewResolver(r.DB)
}

func (r *Resolver) WorkoutProgramConnection() generated.WorkoutProgramConnectionResolver {
	return &connections.WorkoutProgramConnection{}
}

func (r *Resolver) TagConnection() generated.TagConnectionResolver {
	return &connections.TagConnection{}
}

func (r *Resolver) WorkoutCategoryConnection() generated.WorkoutCategoryConnectionResolver {
	return &connections.WorkoutCategoryConnection{}
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
func (r *workoutResolver) Categories(ctx context.Context, obj *workout.Workout, first *int, after uuid.UUID) (*connections.WorkoutCategoryConnection, error) {
	panic("not implemented")
}

type workoutCategoryResolver struct{ *Resolver }

func (r *workoutCategoryResolver) Exercises(ctx context.Context, obj *workout.WorkoutCategory, first *int, after uuid.UUID) (*generated.ExerciseConnection, error) {
	panic("not implemented")
}

func (r *workoutCategoryResolver) Tags(ctx context.Context, obj *workout.WorkoutCategory) ([]*tag.Tag, error) {
	panic("not implemented")
}
