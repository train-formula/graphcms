package workoutcall

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/train-formula/graphcms/database"
	"github.com/train-formula/graphcms/database/cursor"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models/connections"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"go.uber.org/zap"
)

type SearchWorkoutProgram struct {
	Request generated.WorkoutProgramSearchRequest
	First   int
	After   cursor.Cursor
	DB      *pg.DB
	Logger  *zap.Logger
}

func (s SearchWorkoutProgram) logger() *zap.Logger {

	return s.Logger.Named("SearchWorkoutProgram")

}

func (s SearchWorkoutProgram) genQuery(count bool) (string, []interface{}) {

	var query string

	if count {
		query += "SELECT COUNT(1)"
	} else {
		query += "SELECT *"
	}

	query += ` FROM ` + database.TableName(workout.WorkoutProgram{}) + `
							WHERE trainer_organization_id = ?`

	return query, []interface{}{s.Request.TrainerOrganizationID}
}

func (s SearchWorkoutProgram) Validate(ctx context.Context) []validation.ValidatorFunc {

	return []validation.ValidatorFunc{
		validation.DefaultCheckPageSize(s.First),
	}
}

func (s SearchWorkoutProgram) Call(ctx context.Context) (*generated.WorkoutProgramSearchResults, error) {

	var programs []*workout.WorkoutProgram

	query, params := s.genQuery(false)

	query, params, err := database.BasicCursorPaginationQuery(query, "", s.After, workout.WorkoutProgram{}, s.First, params...)
	if err != nil {
		s.logger().Error("Failed to generate pagination query", zap.Error(err))

		return nil, err
	}

	_, err = s.DB.QueryContext(ctx, &programs, query, params...)

	if err != nil {
		s.logger().Error("Failed to search", zap.Error(err))

		return nil, err
	}

	return &generated.WorkoutProgramSearchResults{
		TagFacet: nil,
		Results: &connections.WorkoutProgramConnection{
			ResolveTotalCount: func(ctx context.Context) (int, error) {
				query, params := s.genQuery(true)

				var count int

				_, err := s.DB.QueryContext(ctx, pg.Scan(&count), query, params...)
				if err != nil {
					zap.L().Error("Failed to count", zap.Error(err))

					return -1, err
				}

				return count, nil
			},
			Edges: programs,
		},
	}, nil
}
