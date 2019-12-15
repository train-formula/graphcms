package graphcms

import (
	"github.com/go-pg/pg/v9"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models/connections"
	"github.com/train-formula/graphcms/resolver/exercise"
	"github.com/train-formula/graphcms/resolver/mutation"
	"github.com/train-formula/graphcms/resolver/prescription"
	"github.com/train-formula/graphcms/resolver/query"
	"github.com/train-formula/graphcms/resolver/workout"
	"github.com/train-formula/graphcms/resolver/workoutblock"
	"github.com/train-formula/graphcms/resolver/workoutcategory"
	"github.com/train-formula/graphcms/resolver/workoutprogram"
	"go.uber.org/zap"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{ DB *pg.DB }

/*func (r *WorkoutCategoryResolver) Exercise() generated.ExerciseResolver {
	return &exerciseResolver{r}
}*/
func (r *Resolver) Mutation() generated.MutationResolver {
	return mutation.NewMutationResolver(r.DB, zap.L())
}
func (r *Resolver) Query() generated.QueryResolver {
	return query.NewQueryResolver(r.DB, zap.L())
}
func (r *Resolver) Workout() generated.WorkoutResolver {
	return workout.NewWorkoutResolver(r.DB, zap.L())
}
func (r *Resolver) WorkoutCategory() generated.WorkoutCategoryResolver {
	return workoutcategory.NewWorkoutCategoryResolver(r.DB, zap.L())
}
func (r *Resolver) WorkoutProgram() generated.WorkoutProgramResolver {
	return workoutprogram.NewResolver(r.DB)
}

func (r *Resolver) Prescription() generated.PrescriptionResolver {
	return prescription.NewPrescriptionResolver(r.DB, zap.L())
}

func (r *Resolver) WorkoutBlock() generated.WorkoutBlockResolver {
	return workoutblock.NewWorkoutBlockResolver(r.DB, zap.L())
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

func (r *Resolver) Exercise() generated.ExerciseResolver {
	return exercise.NewExerciseResolver(r.DB, zap.L())
}
