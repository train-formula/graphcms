package organizationcall

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database"
	"github.com/train-formula/graphcms/database/cursor"
	"github.com/train-formula/graphcms/logging"
	"github.com/train-formula/graphcms/models/connections"
	"github.com/train-formula/graphcms/models/tag"
	"github.com/train-formula/graphcms/validation"
	"github.com/willtrking/pgxload"
	"go.uber.org/zap"
)

func NewGetOrganizationAvailableTags(trainerOrganizationID uuid.UUID, first int, after cursor.Cursor, logger *zap.Logger, db pgxload.PgxLoader) *GetOrganizationAvailableTags {
	return &GetOrganizationAvailableTags{
		trainerOrganizationID: trainerOrganizationID,
		first:                 first,
		after:                 after,
		db:                    db,
		logger:                logger.Named("GetOrganizationAvailableTags"),
	}

}

type GetOrganizationAvailableTags struct {
	trainerOrganizationID uuid.UUID
	first                 int
	after                 cursor.Cursor
	db                    pgxload.PgxLoader
	logger                *zap.Logger
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

	return query, []interface{}{g.trainerOrganizationID}
}

func (g GetOrganizationAvailableTags) Validate(ctx context.Context) []validation.ValidatorFunc {
	return []validation.ValidatorFunc{
		validation.DefaultCheckPageSize(g.first),
	}
}

func (g GetOrganizationAvailableTags) Call(ctx context.Context) (*connections.TagConnection, error) {

	var tags []*tag.Tag

	query, params := g.genQuery(false)

	/*query, params, err := database.BasicCursorPaginationQuery(query, "", g.after, tag.tag{}, g.first, params...)
	if err != nil {
		g.logger().Error("Failed to generate pagination query", zap.Error(err))

		return nil, err
	}*/

	rows, err := g.db.Query(ctx, pgxload.RebindPositional(query), params...)

	if err != nil {
		g.logger.Error("Failed to load organization tags", zap.Error(err),
			logging.UUID("trainerOrganizationID", g.trainerOrganizationID), zap.Int("first", g.first),
			logging.Cursor("after", g.after))

		return nil, err
	}

	err = g.db.Scanner(rows).Scan(&tags)
	if err != nil {
		g.logger.Error("Failed to scan organization tags", zap.Error(err),
			logging.UUID("trainerOrganizationID", g.trainerOrganizationID), zap.Int("first", g.first),
			logging.Cursor("after", g.after))

		return nil, err
	}

	return &connections.TagConnection{
		ResolveTotalCount: func(ctx context.Context) (int, error) {
			query, params := g.genQuery(true)

			var count int

			_, err := g.db.Query(ctx, query, params...)
			if err != nil {
				g.logger.Error("Failed to count organization tags", zap.Error(err),
					logging.UUID("trainerOrganizationID", g.trainerOrganizationID), zap.Int("first", g.first),
					logging.Cursor("after", g.after))

				return -1, err
			}

			err = g.db.Scanner(rows).Scan(&count)
			if err != nil {
				g.logger.Error("Failed to scan organization tag count", zap.Error(err),
					logging.UUID("trainerOrganizationID", g.trainerOrganizationID), zap.Int("first", g.first),
					logging.Cursor("after", g.after))

				return -1, err
			}

			return count, nil
		},
		Edges: tags,
	}, nil
}
