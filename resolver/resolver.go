package resolver

import (
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models/connections"
	"github.com/train-formula/graphcms/resolver/blockexercise"
	"github.com/train-formula/graphcms/resolver/exercise"
	"github.com/train-formula/graphcms/resolver/mutation"
	"github.com/train-formula/graphcms/resolver/plan"
	"github.com/train-formula/graphcms/resolver/planschedule"
	"github.com/train-formula/graphcms/resolver/prescription"
	"github.com/train-formula/graphcms/resolver/prescriptionset"
	"github.com/train-formula/graphcms/resolver/query"
	"github.com/train-formula/graphcms/resolver/unitdata"
	"github.com/train-formula/graphcms/resolver/workout"
	"github.com/train-formula/graphcms/resolver/workoutblock"
	"github.com/train-formula/graphcms/resolver/workoutcategory"
	"github.com/train-formula/graphcms/resolver/workoutprogram"
	"github.com/willtrking/pgxload"
	"go.uber.org/zap"
)

func NewResolver(db pgxload.PgxLoader, logger *zap.Logger) *Resolver {
	return &Resolver{
		db:     db,
		logger: logger.Named("graphql"),
	}
}

type Resolver struct {
	db     pgxload.PgxLoader
	logger *zap.Logger
}

var _ generated.ResolverRoot = (*Resolver)(nil)

func (r *Resolver) Mutation() generated.MutationResolver {
	return mutation.NewMutationResolver(r.db, r.logger)
}
func (r *Resolver) Query() generated.QueryResolver {
	return query.NewQueryResolver(r.db, r.logger)
}
func (r *Resolver) Workout() generated.WorkoutResolver {
	return workout.NewWorkoutResolver(r.db, r.logger)
}
func (r *Resolver) WorkoutCategory() generated.WorkoutCategoryResolver {
	return workoutcategory.NewWorkoutCategoryResolver(r.db, r.logger)
}
func (r *Resolver) WorkoutProgram() generated.WorkoutProgramResolver {
	return workoutprogram.NewWorkoutProgramResolver(r.db, r.logger)
}

func (r *Resolver) Prescription() generated.PrescriptionResolver {
	return prescription.NewPrescriptionResolver(r.db, r.logger)
}

func (r *Resolver) PrescriptionSet() generated.PrescriptionSetResolver {
	return prescriptionset.NewPrescriptionSetResolver(r.db, r.logger)
}

func (r *Resolver) WorkoutBlock() generated.WorkoutBlockResolver {
	return workoutblock.NewWorkoutBlockResolver(r.db, r.logger)
}

func (r *Resolver) Exercise() generated.ExerciseResolver {
	return exercise.NewExerciseResolver(r.db, r.logger)
}

func (r *Resolver) BlockExercise() generated.BlockExerciseResolver {
	return blockexercise.NewBlockExerciseResolver(r.db, r.logger)
}

func (r *Resolver) UnitData() generated.UnitDataResolver {
	return unitdata.NewUnitDataResolver(r.db, r.logger)
}

func (r *Resolver) Plan() generated.PlanResolver {
	return plan.NewPlanResolver(r.db, r.logger)
}

func (r *Resolver) PlanSchedule() generated.PlanScheduleResolver {
	return planschedule.NewPlanScheduleResolver(r.db, r.logger)
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

func (r *Resolver) ExerciseConnection() generated.ExerciseConnectionResolver {
	return &connections.ExerciseConnection{}
}

func (r *Resolver) PrescriptionConnection() generated.PrescriptionConnectionResolver {
	return &connections.PrescriptionConnection{}
}
func (r *Resolver) PlanConnection() generated.PlanConnectionResolver {
	return &connections.PlanConnection{}
}
