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

func NewSearchPrescriptions(request generated.PrescriptionSearchRequest, first int, after cursor.Cursor, logger *zap.Logger, db *pg.DB) *SearchPrescriptions {
	return &SearchPrescriptions{
		request: request,
		first:   first,
		after:   after,
		db:      db,
		logger:  logger.Named("SearchPrescriptions"),
	}
}

type SearchPrescriptions struct {
	request generated.PrescriptionSearchRequest
	first   int
	after   cursor.Cursor
	db      *pg.DB
	logger  *zap.Logger
}

func (s SearchPrescriptions) genQuery(count bool) (string, []interface{}) {

	var query string

	if count {
		query += "SELECT COUNT(1)"
	} else {
		query += "SELECT *"
	}

	var params []interface{}

	if len(s.request.TagUUIDs) > 0 {
		query += ` FROM ` + database.TableName(workout.Prescription{}) + ` p ` +
			` INNER JOIN ` + database.TableName(tag.Tagged{}) + ` t ` +
			` ON p.id = t.tagged_id WHERE p.trainer_organization_id = ? AND `

		params = []interface{}{s.request.TrainerOrganizationID}

		for _, tagUUID := range s.request.TagUUIDs {
			params = append(params, tagUUID)
		}

	} else {
		query += ` FROM ` + database.TableName(workout.Prescription{}) + `
								WHERE trainer_organization_id = ?`

		params = []interface{}{s.request.TrainerOrganizationID}
	}

	return query, params

}

func (s SearchPrescriptions) Validate(ctx context.Context) []validation.ValidatorFunc {

	return []validation.ValidatorFunc{
		validation.DefaultCheckPageSize(s.first),
	}
}

func (s SearchPrescriptions) Call(ctx context.Context) (*generated.PrescriptionSearchResults, error) {

	var prescriptions []*workout.Prescription

	query, params := s.genQuery(false)

	/*query, params, err := database.BasicCursorPaginationQuery(query, "", s.after, workout.WorkoutProgram{}, s.first, params...)
	if err != nil {
		s.logger().Error("Failed to generate pagination query", zap.Error(err))

		return nil, err
	}*/

	_, err := s.db.QueryContext(ctx, &prescriptions, query, params...)

	if err != nil {
		s.logger.Error("Failed to search prescriptions", zap.Error(err))

		return nil, err
	}

	return &generated.PrescriptionSearchResults{
		TagFacet: nil,
		Results: &connections.PrescriptionConnection{
			ResolveTotalCount: func(ctx context.Context) (int, error) {
				query, params := s.genQuery(true)

				var count int

				_, err := s.db.QueryContext(ctx, pg.Scan(&count), query, params...)
				if err != nil {
					s.logger.Error("Failed to count prescriptions", zap.Error(err))

					return -1, err
				}

				return count, nil
			},
			Edges: prescriptions,
		},
	}, nil
}
