package workoutblock

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/train-formula/graphcms/calls/organizationcall"
	"github.com/train-formula/graphcms/calls/workoutcall"
	"github.com/train-formula/graphcms/logging"
	"github.com/train-formula/graphcms/models/trainer"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
	"go.uber.org/zap"
)

type WorkoutBlockResolver struct {
	db     *pg.DB
	logger *zap.Logger
}

func NewWorkoutBlockResolver(db *pg.DB, logger *zap.Logger) *WorkoutBlockResolver {
	return &WorkoutBlockResolver{
		db:     db,
		logger: logger.Named("WorkoutBlockResolver"),
	}
}

func (r *WorkoutBlockResolver) Round(ctx context.Context, obj *workout.WorkoutBlock) (*workout.UnitData, error) {

	if obj == nil {
		return nil, nil
	}

	if obj.RoundUnitID == nil && obj.RoundNumeral == nil && obj.RoundText == nil {
		return nil, nil
	} else if obj.RoundUnitID == nil && (obj.RoundNumeral != nil || obj.RoundText != nil) {
		r.logger.Error("Workout block malformed, has round numeral and/or round text but no round unit ID", logging.UUID("workoutBlockID", obj.ID))
		return nil, gqlerror.Errorf("Workout block malformed, has round numeral and/or round text but no round unit ID")
	}

	return &workout.UnitData{
		Numeral: obj.RoundNumeral,
		Text:    obj.RoundText,
		UnitID:  *obj.RoundUnitID,
	}, nil
}

func (r *WorkoutBlockResolver) TrainerOrganization(ctx context.Context, obj *workout.WorkoutBlock) (*trainer.Organization, error) {
	if obj == nil {
		return nil, nil
	}

	g := organizationcall.NewGetOrganization(obj.TrainerOrganizationID, r.logger, r.db)

	if validation.ValidationChain(ctx, g.Validate(ctx)...) {
		return g.Call(ctx)
	}

	return nil, nil
}

func (r *WorkoutBlockResolver) WorkoutCategory(ctx context.Context, obj *workout.WorkoutBlock) (*workout.WorkoutCategory, error) {

	if obj == nil {
		return nil, nil
	}

	g := workoutcall.NewGetWorkoutCategory(obj.WorkoutCategoryID, r.logger, r.db)

	if validation.ValidationChain(ctx, g.Validate(ctx)...) {
		return g.Call(ctx)
	}

	return nil, nil
}

func (r *WorkoutBlockResolver) Exercises(ctx context.Context, obj *workout.WorkoutBlock) ([]*workout.BlockExercise, error) {
	if obj == nil {
		return nil, nil
	}

	call := workoutcall.NewGetBlockExercises(obj.ID, r.logger, r.db)

	if validation.ValidationChain(ctx, call.Validate(ctx)...) {
		return call.Call(ctx)
	}

	return nil, nil

}
