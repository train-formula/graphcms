package workoutcall

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/train-formula/graphcms/database"
	"github.com/train-formula/graphcms/database/cursor"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models/connections"
	"github.com/train-formula/graphcms/models/tag"
	"github.com/train-formula/graphcms/models/workout"
	"github.com/train-formula/graphcms/validation"
	"go.uber.org/zap"
)

func NewSearchWorkoutCategory(request generated.WorkoutCategorySearchRequest, first int, after cursor.Cursor, logger *zap.Logger, db *pg.DB) *SearchWorkoutCategory {
	return &SearchWorkoutCategory{
		request: request,
		first:   first,
		after:   after,
		db:      db,
		logger:  logger.Named("SearchWorkoutCategory"),
	}
}

type SearchWorkoutCategory struct {
	request generated.WorkoutCategorySearchRequest
	first   int
	after   cursor.Cursor
	db      *pg.DB
	logger  *zap.Logger
}

func (s SearchWorkoutCategory) genQuery(count bool) (string, []interface{}) {

	var query string

	if count {
		query += "SELECT COUNT(1)"
	} else {
		query += "SELECT *"
	}

	var params []interface{}

	if len(s.request.TagUUIDs) > 0 {
		query += ` FROM ` + database.TableName(workout.WorkoutCategory{}) + ` wc ` +
			` INNER JOIN ` + database.TableName(tag.Tagged{}) + ` t ` +
			` ON wc.workoutBlockID = t.tagged_id WHERE wc.trainer_organization_id = ? AND `

		params = []interface{}{s.request.TrainerOrganizationID}

		for _, tagUUID := range s.request.TagUUIDs {
			params = append(params, tagUUID)
		}

	} else {
		query += ` FROM ` + database.TableName(workout.WorkoutCategory{}) + `
								WHERE trainer_organization_id = ?`

		params = []interface{}{s.request.TrainerOrganizationID}
	}

	return query, params

}

func (s SearchWorkoutCategory) Validate(ctx context.Context) []validation.ValidatorFunc {

	return []validation.ValidatorFunc{
		validation.DefaultCheckPageSize(s.first),
	}
}

func (s SearchWorkoutCategory) Call(ctx context.Context) (*generated.WorkoutCategorySearchResults, error) {

	var programs []*workout.WorkoutCategory

	query, params := s.genQuery(false)

	/*query, params, err := database.BasicCursorPaginationQuery(query, "", s.after, workout.WorkoutProgram{}, s.first, params...)
	if err != nil {
		s.logger().Error("Failed to generate pagination query", zap.Error(err))

		return nil, err
	}*/

	_, err := s.db.QueryContext(ctx, &programs, query, params...)

	if err != nil {
		s.logger.Error("Failed to search workout categories", zap.Error(err))

		return nil, err
	}

	return &generated.WorkoutCategorySearchResults{
		TagFacet: nil,
		Results: &connections.WorkoutCategoryConnection{
			ResolveTotalCount: func(ctx context.Context) (int, error) {
				query, params := s.genQuery(true)

				var count int

				_, err := s.db.QueryContext(ctx, pg.Scan(&count), query, params...)
				if err != nil {
					s.logger.Error("Failed to count workout categories", zap.Error(err))

					return -1, err
				}

				return count, nil
			},
			Edges: programs,
		},
	}, nil
}
