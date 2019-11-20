package tag

import (
	"errors"
	"time"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database/cursor"
)

type Tag struct {
	tableName struct{} `sql:"tag.tags"`

	ID                    uuid.UUID `json:"id"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	Tag                   string    `json:"tag"`
	TrainerOrganizationID uuid.UUID `json:"trainer_organization_id"`
}

func (t Tag) TableName() string {
	return "tag.tags"
}

// Generate an SQL query with for cursor that paginates with columns from this table
// Also provide the params to go with it
func (t Tag) CursorQuery(prefix string, c cursor.Cursor) (string, []interface{}, error) {

	if conv, proper := c.(*cursor.TimeUUIDCursor); proper {

		q, p := conv.GenerateAscendingSQLConditions(prefix, "created_at", "uuid")

		return q, p, nil
	}

	return "", nil, errors.New("cursor must be a Time UUID cursor")

}

// Generated a database query for the particular row represented
func (t *Tag) DBCursor() cursor.Cursor {
	return cursor.NewTimeUUIDCursor(t.CreatedAt, t.ID)
}

// Serializes the result of DBCursor
// Necessary for gqlgen, allows us to avoid creating a seperate TagEdge
func (t *Tag) Cursor() string {
	return t.DBCursor().Serialize()
}

// Necessary for gqlgen, allows us to avoid creating a seperate TagEdge
func (w *Tag) Node() *Tag {
	return w
}
