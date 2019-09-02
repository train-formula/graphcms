package trainerdb

import (
	"context"
	"strings"

	uuid "github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database"
	"github.com/train-formula/graphcms/models/trainer"
)

// Retrieves an individual organization by its ID
func GetOrganization(ctx context.Context, conn database.Conn, id uuid.UUID) (trainer.Organization, error) {

	var result trainer.Organization

	_, err := conn.QueryOneContext(ctx, &result, "SELECT * FROM "+database.TableName(result)+" WHERE id = ?", id)

	return result, err
}

// Retrieves individual organizations by their IDs
func GetOrganizations(ctx context.Context, conn database.Conn, ids []uuid.UUID) ([]*trainer.Organization, error) {

	if len(ids) <= 0 {
		return nil, nil
	}

	var result []*trainer.Organization

	query := "SELECT * FROM " + database.TableName(trainer.Organization{}) + " WHERE "

	var params []interface{}

	for _, id := range ids {
		query += "id = ? OR "
		params = append(params, id)
	}

	query = strings.TrimSuffix(query, " OR ")

	_, err := conn.QueryContext(ctx, &result, query, params...)

	return result, err
}
