package plan

import (
	"errors"
	"time"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database/cursor"
)

type Plan struct {
	tableName struct{} `sql:"plan.plan"`

	ID                    uuid.UUID `json:"id"`
	CreatedAt             time.Time `json:"created_at" pgxload:"omitZero"`
	UpdatedAt             time.Time `json:"updated_at" pgxload:"omitZero"`
	TrainerOrganizationID uuid.UUID `json:"trainerOrganizationID"`
	Name                  string    `json:"name"`
	Description           *string   `json:"description"`
	RegistrationAvailable bool      `json:"registrationAvailable"`
	Archived              bool      `json:"archived"`
}

func (t Plan) TableName() string {
	return "plan.plan"
}

// Generate an SQL query with for cursor that paginates with columns from this table
// Also provide the params to go with it
func (t Plan) CursorQuery(prefix string, c cursor.Cursor) (string, []interface{}, error) {

	if conv, proper := c.(*cursor.TimeUUIDCursor); proper {

		q, p := conv.GenerateAscendingSQLConditions(prefix, "created_at", "uuid")

		return q, p, nil
	}

	return "", nil, errors.New("cursor must be a Time UUID cursor")

}

// Generated a database query for the particular row represented
func (t *Plan) DBCursor() cursor.Cursor {
	return cursor.NewTimeUUIDCursor(t.CreatedAt, t.ID)
}

// Serializes the result of DBCursor
// Necessary for gqlgen, allows us to avoid creating a seperate PlanEdge
func (t *Plan) Cursor() string {
	return t.DBCursor().Serialize()
}

// Necessary for gqlgen, allows us to avoid creating a seperate PlanEdge
func (w *Plan) Node() *Plan {
	return w
}
