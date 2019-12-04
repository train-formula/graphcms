package organizationcall

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database"
	"github.com/train-formula/graphcms/database/cursor"
	"github.com/train-formula/graphcms/models/connections"
	"github.com/train-formula/graphcms/models/tag"
	"github.com/train-formula/graphcms/validation"
	"go.uber.org/zap"
)

type GetOrganizationAvailableTags struct {
	TrainerOrganizationID uuid.UUID
	First                 int
	After                 cursor.Cursor
	DB                    *pg.DB
	Logger                *zap.Logger
}

func (g GetOrganizationAvailableTags) logger() *zap.Logger {

	return g.Logger.Named("GetOrganizationAvailableTags")

}

func (g GetOrganizationAvailableTags) genQuery(count bool) (string, []interface{}) {

	var query string

	if count {
		query += "SELECT COUNT(1)"
	} else {
		query += "SELECT *"
	}

	query += ` FROM ` + database.TableName(tag.Tag{}) + `
							WHERE trainer_organization_id = ?`

	return query, []interface{}{g.TrainerOrganizationID}
}

func (g GetOrganizationAvailableTags) Validate(ctx context.Context) []validation.ValidatorFunc {
	return []validation.ValidatorFunc{
		validation.DefaultCheckPageSize(g.First),
	}
}

func (g GetOrganizationAvailableTags) Call(ctx context.Context) (*connections.TagConnection, error) {

	var tags []*tag.Tag

	query, params := g.genQuery(false)

	/*query, params, err := database.BasicCursorPaginationQuery(query, "", g.After, tag.Tag{}, g.First, params...)
	if err != nil {
		g.logger().Error("Failed to generate pagination query", zap.Error(err))

		return nil, err
	}*/

	_, err := g.DB.QueryContext(ctx, &tags, query, params...)

	if err != nil {
		g.logger().Error("Failed to organization tags", zap.Error(err))

		return nil, err
	}

	return &connections.TagConnection{
		ResolveTotalCount: func(ctx context.Context) (int, error) {
			query, params := g.genQuery(true)

			var count int

			_, err := g.DB.QueryContext(ctx, pg.Scan(&count), query, params...)
			if err != nil {
				zap.L().Error("Failed to count", zap.Error(err))

				return -1, err
			}

			return count, nil
		},
		Edges: tags,
	}, nil
}
