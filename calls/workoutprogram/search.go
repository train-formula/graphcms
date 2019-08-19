package workoutprogram

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v9"
	"github.com/train-formula/graphcms/database"
	"github.com/train-formula/graphcms/database/cursor"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"go.uber.org/zap"
)

type Search struct {
	DB *pg.DB
}

func (s Search) Validate(ctx context.Context, request generated.WorkoutProgramSearchRequest, first int, after *cursor.TimeUUIDCursor) bool {

	return validation.ValidationChain(ctx,
		validation.CheckPageSize(first, 1, 50),
	)
}

func (s Search) Call(ctx context.Context, request generated.WorkoutProgramSearchRequest, first int, after *cursor.TimeUUIDCursor) (*generated.WorkoutProgramSearchResults, error) {

	var programs []*workout.WorkoutProgram

	_, err := s.DB.QueryContext(ctx, &programs, `SELECT * FROM `+database.TableName(workout.WorkoutProgram{})+`
							WHERE trainer_organization_id = ? LIMIT ?`, request.TrainerOrganizationID, first)

	if err != nil {
		zap.L().Error("Failed to search", zap.Error(err))

		return nil, err
	}

	fmt.Println(cursor.NewTimeUUIDCursor(programs[0].CreatedAt, programs[0].ID).Serialize())

	return &generated.WorkoutProgramSearchResults{
		TagFacet: nil,
		Results:  programs,
	}, nil
}
