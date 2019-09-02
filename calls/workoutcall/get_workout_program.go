package workoutcall

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"github.com/vektah/gqlparser/gqlerror"
	"go.uber.org/zap"
)

type GetWorkoutProgram struct {
	ID uuid.UUID
	DB *pg.DB
}

func (g GetWorkoutProgram) Validate(ctx context.Context) []validation.ValidatorFunc {

	return nil
}

func (g GetWorkoutProgram) Call(ctx context.Context) (*workout.WorkoutProgram, error) {

	var result workout.WorkoutProgram

	_, err := g.DB.QueryOneContext(ctx, &result, "SELECT * FROM "+database.TableName(result)+" WHERE id = ?", g.ID)
	if err == pg.ErrNoRows {
		return nil, gqlerror.Errorf("Unknown workout program ID")
	} else if err != nil {
		zap.L().Error("Failed to retrieve workout program", zap.Error(err))
		return nil, gqlerror.Errorf("Failed to retrieve workout program")
	}

	return &result, nil
}
