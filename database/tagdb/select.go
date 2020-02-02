package tagdb

import (
	"context"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database"
	"github.com/train-formula/graphcms/models/tag"
)

// request struct to retrieve a tag by its tag + organization id
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

// request struct to retrieve tags by the type of object they are attached to, the object UUID + the trainer organization UUID
// Grouped together for easy batching
type TagsByObject struct {
	ObjectUUID uuid.UUID
	ObjectType tag.TagType
}

// Retrieves an individual tag by its id
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

// Retrieves an individual tag by its tag and organization id
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

// Retrieves tags by the object's they are attached to
func GetTagsByObject(ctx context.Context, conn database.Conn, byObject []TagsByObject) (map[TagsByObject][]*tag.Tag, error) {

	if len(byObject) <= 0 {
		return nil, nil
	}

	var results []*tag.TaggedTagJoin

	query := "SELECT " + (tag.TaggedTagJoin{}).SelectColumns("t", "tg") +
		" FROM " + database.TableName(tag.Tagged{}) + " tg " +
		" INNER JOIN " + database.TableName(tag.Tag{}) + " t ON tg.tag_uuid = t.id " +
		" WHERE "

	var params []interface{}

	for _, bo := range byObject {
		query += "(tg.tagged_id = ? AND tag_type = ?) OR "
		params = append(params, bo.ObjectUUID, bo.ObjectType)
	}

	query = strings.TrimSuffix(query, " OR ")

	_, err := conn.QueryContext(ctx, &results, query, params...)

	if err != nil {
		return nil, err
	}

	final := make(map[TagsByObject][]*tag.Tag)

	for _, result := range results {
		key := TagsByObject{
			ObjectUUID: result.TaggedTaggedID,
			ObjectType: result.TaggedTagType,
		}

		final[key] = append(final[key], result.Tag())
	}

	return final, nil
}
