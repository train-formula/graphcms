package tagcall

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/train-formula/graphcms/database"
	"github.com/train-formula/graphcms/database/cursor"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models/connections"
	"github.com/train-formula/graphcms/models/tag"
	"github.com/train-formula/graphcms/validation"
	"go.uber.org/zap"
)

func NewSearchTags(request generated.TagSearchRequest, first int, after cursor.Cursor, logger *zap.Logger, db *pg.DB) *SearchTags {
	return &SearchTags{
		request: request,
		first:   first,
		after:   after,
		db:      db,
		logger:  logger.Named("SearchTags"),
	}
}

type SearchTags struct {
	request generated.TagSearchRequest
	first   int
	after   cursor.Cursor
	db      *pg.DB
	logger  *zap.Logger
}

func (s SearchTags) genQuery(count bool) (string, []interface{}) {

	var query string

	if count {
		query += "SELECT COUNT(1)"
	} else {
		query += "SELECT *"
	}

	query += ` FROM ` + database.TableName(tag.Tag{}) + `
								WHERE trainer_organization_id = ?`

	params := []interface{}{s.request.TrainerOrganizationID}

	return query, params

}

func (s SearchTags) Validate(ctx context.Context) []validation.ValidatorFunc {

	return []validation.ValidatorFunc{
		validation.DefaultCheckPageSize(s.first),
	}
}

func (s SearchTags) Call(ctx context.Context) (*generated.TagSearchResults, error) {

	var tags []*tag.Tag

	query, params := s.genQuery(false)

	/*query, params, err := database.BasicCursorPaginationQuery(query, "", s.after, workout.WorkoutProgram{}, s.first, params...)
	if err != nil {
		s.logger().Error("Failed to generate pagination query", zap.Error(err))

		return nil, err
	}*/

	_, err := s.db.QueryContext(ctx, &tags, query, params...)

	if err != nil {
		s.logger.Error("Failed to search tags", zap.Error(err))

		return nil, err
	}

	return &generated.TagSearchResults{
		Results: &connections.TagConnection{
			ResolveTotalCount: func(ctx context.Context) (int, error) {
				query, params := s.genQuery(true)

				var count int

				_, err := s.db.QueryContext(ctx, pg.Scan(&count), query, params...)
				if err != nil {
					s.logger.Error("Failed to count tags", zap.Error(err))

					return -1, err
				}

				return count, nil
			},
			Edges: tags,
		},
	}, nil
}
