package organization

import (
	"context"

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
