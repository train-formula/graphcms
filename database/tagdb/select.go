package tagdb

import (
	"context"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database"
	"github.com/train-formula/graphcms/models/tag"
)

// Request struct to retrieve a tag by its tag + organization ID
// Grouped together for easy batching
type TagByTag struct {
	Tag                   string
	TrainerOrganizationID uuid.UUID
}

// Get a stable version of this struct (e.g. suitable for map keys)
func (t TagByTag) Stable() TagByTag {

	return TagByTag{
		Tag:                   strings.ToLower(t.Tag),
		TrainerOrganizationID: t.TrainerOrganizationID,
	}
}

// Retrieves an individual tag by its ID
func GetTag(ctx context.Context, conn database.Conn, id uuid.UUID) (tag.Tag, error) {

	var result tag.Tag

	_, err := conn.QueryOneContext(ctx, &result, "SELECT * FROM "+database.TableName(result)+" WHERE id = ?", id)

	return result, err
}

// Retrieves individual tags by their IDs
func GetTags(ctx context.Context, conn database.Conn, ids []uuid.UUID) ([]*tag.Tag, error) {

	if len(ids) <= 0 {
		return nil, nil
	}

	var result []*tag.Tag

	query := "SELECT * FROM " + database.TableName(tag.Tag{}) + " WHERE "

	var params []interface{}

	for _, id := range ids {
		query += "id = ? OR "
		params = append(params, id)
	}

	query = strings.TrimSuffix(query, " OR ")

	_, err := conn.QueryContext(ctx, &result, query, params...)

	return result, err
}

// Retrieves an individual tag by its tag and organization ID
func GetTagByTag(ctx context.Context, conn database.Conn, byTag TagByTag) (tag.Tag, error) {

	var result tag.Tag

	_, err := conn.QueryOneContext(ctx, &result, "SELECT * FROM "+database.TableName(result)+
		" WHERE trainer_organization_id = ? AND LOWER(tag) = LOWER(?)", byTag.TrainerOrganizationID, byTag.Tag)

	return result, err
}

// Retrieves individual tags by their organization IDs and tags
func GetTagsByTag(ctx context.Context, conn database.Conn, byTag []TagByTag) ([]*tag.Tag, error) {

	if len(byTag) <= 0 {
		return nil, nil
	}

	var result []*tag.Tag

	query := "SELECT * FROM " + database.TableName(tag.Tag{}) + " WHERE "

	var params []interface{}

	for _, bt := range byTag {
		query += "(trainer_organization_id = ? AND LOWER(tag) = LOWER(?)) OR "
		params = append(params, bt.TrainerOrganizationID, bt.Tag)
	}

	query = strings.TrimSuffix(query, " OR ")

	_, err := conn.QueryContext(ctx, &result, query, params...)

	return result, err
}
