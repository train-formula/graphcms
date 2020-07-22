package workout

import (
	"errors"
	"time"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database/cursor"
	"github.com/train-formula/graphcms/database/types"
)

type Exercise struct {
	tableName             struct{}         `sql:"workout.exercise"`
	ID                    uuid.UUID        `json:"id"`
	CreatedAt             time.Time        `json:"createdAt" pgxload:"defaultZero"`
	UpdatedAt             time.Time        `json:"updatedAt" pgxload:"defaultZero"`
	TrainerOrganizationID uuid.UUID        `json:"trainerOrganizationID"`
	Name                  string           `json:"name"`
	Description           string           `json:"description"`
	VideoURL              types.NullString `json:"videoURL"`
}

func (e Exercise) TableName() string {
	return "workout.exercise"
}

// Generate an SQL query with for cursor that paginates with columns from this table
// Also provide the params to go with it
func (e Exercise) CursorQuery(prefix string, c cursor.Cursor) (string, []interface{}, error) {

	if conv, proper := c.(*cursor.TimeUUIDCursor); proper {

		q, p := conv.GenerateAscendingSQLConditions(prefix, "created_at", "uuid")

		return q, p, nil
	}

	return "", nil, errors.New("cursor must be a Time UUID cursor")

}

// Generated a database query for the particular row represented
func (e *Exercise) DBCursor() cursor.Cursor {
	return cursor.NewTimeUUIDCursor(e.CreatedAt, e.ID)
}

// Serializes the result of DBCursor
// Necessary for gqlgen, allows us to avoid creating a seperate ExerciseEdge
func (e *Exercise) Cursor() string {
	return e.DBCursor().Serialize()
}

// Necessary for gqlgen, allows us to avoid creating a seperate ExerciseEdge
func (e *Exercise) Node() *Exercise {
	return e
}
