package organization

import (
	"context"

	"github.com/satori/go.uuid"
	"github.com/train-formula/graphcms/database"
	"github.com/train-formula/graphcms/models/trainer"
)

func GetOrganization(ctx context.Context, conn database.Conn, id uuid.UUID) (trainer.Organization, error) {

	var result trainer.Organization

	_, err := conn.QueryOneContext(ctx, &result, "SELECT * FROM "+database.TableName(result)+" WHERE id = ?", id)

	return result, err
}