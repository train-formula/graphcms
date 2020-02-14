package trainerdb

import (
	"context"
	"strings"

	uuid "github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database"
	"github.com/train-formula/graphcms/models/trainer"
	"github.com/willtrking/pgxload"
)

// Retrieves an individual organization by its id
func GetOrganization(ctx context.Context, conn pgxload.PgxLoader, id uuid.UUID) (trainer.Organization, error) {

	var result trainer.Organization

	rows, err := conn.Query(ctx, pgxload.RebindPositional("SELECT * FROM "+database.TableName(trainer.Organization{})+" WHERE id = ?"), id)
	if err != nil {
		return trainer.Organization{}, err
	}

	err = conn.Scanner(rows).Scan(&result)

	return result, err
}

// Retrieves individual organizations by their IDs
func GetOrganizations(ctx context.Context, conn pgxload.PgxLoader, ids []uuid.UUID) ([]*trainer.Organization, error) {

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

	rows, err := conn.Query(ctx, pgxload.RebindPositional(query), params...)
	if err != nil {
		return nil, err
	}

	err = conn.Scanner(rows).Scan(&result)

	return result, err
}
