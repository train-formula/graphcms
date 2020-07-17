package workoutcall

import (
	"context"

	"github.com/train-formula/graphcms/database"
	"github.com/train-formula/graphcms/database/cursor"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models/connections"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"github.com/willtrking/pgxload"
	"go.uber.org/zap"
)

func NewSearchWorkoutProgram(request generated.WorkoutProgramSearchRequest, first int, after cursor.Cursor, logger *zap.Logger, db pgxload.PgxLoader) *SearchWorkoutProgram {
	return &SearchWorkoutProgram{
		request: request,
		first:   first,
		after:   after,
		db:      db,
		logger:  logger.Named("SearchWorkoutProgram"),
	}
}

type SearchWorkoutProgram struct {
	request generated.WorkoutProgramSearchRequest
	first   int
	after   cursor.Cursor
	db      pgxload.PgxLoader
	logger  *zap.Logger
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

	return query, []interface{}{s.request.TrainerOrganizationID}
}

func (s SearchWorkoutProgram) Validate(ctx context.Context) []validation.ValidatorFunc {

	return []validation.ValidatorFunc{
		validation.DefaultCheckPageSize(s.first),
	}
}

func (s SearchWorkoutProgram) Call(ctx context.Context) (*generated.WorkoutProgramSearchResults, error) {

	var programs []*workout.WorkoutProgram

	query, params := s.genQuery(false)

	/*query, params, err := database.BasicCursorPaginationQuery(query, "", s.after, workout.WorkoutProgram{}, s.first, params...)
	if err != nil {
		s.logger().Error("Failed to generate pagination query", zap.Error(err))

		return nil, err
	}*/

	rows, err := s.db.Query(ctx, pgxload.RebindPositional(query), params...)

	if err != nil {
		s.logger.Error("Failed to search workout programs", zap.Error(err))

		return nil, err
	}

	err = s.db.Scanner(rows).Scan(&programs)
	if err != nil {
		s.logger.Error("Failed to scan workout programs", zap.Error(err))

		return nil, err
	}

	return &generated.WorkoutProgramSearchResults{
		TagFacet: nil,
		Results: &connections.WorkoutProgramConnection{
			ResolveTotalCount: func(ctx context.Context) (int, error) {
				query, params := s.genQuery(true)

				var count int

				rows, err := s.db.Query(ctx, pgxload.RebindPositional(query), params...)
				if err != nil {
					s.logger.Error("Failed to count workout programs", zap.Error(err))

					return -1, err
				}

				err = s.db.Scanner(rows).Scan(&count)
				if err != nil {
					s.logger.Error("Failed to scan workout programs count", zap.Error(err))

					return -1, err
				}

				return count, nil
			},
			Edges: programs,
		},
	}, nil
}
