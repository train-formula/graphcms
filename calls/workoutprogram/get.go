package workoutprogram

import (
	"context"

	"github.com/go-pg/pg/v9"
	uuid "github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/vektah/gqlparser/gqlerror"
	"go.uber.org/zap"
)

type Get struct {
	DB *pg.DB
}

func (g Get) Validate(ctx context.Context, id uuid.UUID) bool {

	return true
}

func (g Get) Call(ctx context.Context, id uuid.UUID) (*workout.WorkoutProgram, error) {

	var result workout.WorkoutProgram

	_, err := g.DB.QueryOneContext(ctx, &result, "SELECT * FROM "+database.TableName(result)+" WHERE id = ?", id)
	if err == pg.ErrNoRows {
		return nil, gqlerror.Errorf("Unknown workout program ID")
	} else if err != nil {
		zap.L().Error("Failed to retrieve workout program", zap.Error(err))
		return nil, gqlerror.Errorf("Failed to retrieve workout program")
	}

	return &result, nil
}
