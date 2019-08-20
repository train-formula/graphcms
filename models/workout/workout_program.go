package workout

import (
	"errors"
	"time"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database/cursor"
)

type WorkoutProgram struct {
	ID                    uuid.UUID `json:"id"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
	Name                  string `json:"name"`
	Description           string `json:"description"`
	Public                bool
	Price                 string
	TrainerOrganizationID uuid.UUID `json:"trainerOrganizationID"`
}

func (w WorkoutProgram) TableName() string {
	return "workout.program"
}

// Generate an SQL query with for cursor that paginates with columns from this table
// Also provide the params to go with it
func (w WorkoutProgram) CursorQuery(prefix string, c cursor.Cursor) (string, []interface{}, error) {

	if conv, proper := c.(*cursor.TimeUUIDCursor); proper {

		q, p := conv.GenerateAscendingSQLConditions(prefix, "created_at", "uuid")

		return q, p, nil
	}

	return "", nil, errors.New("cursor must be a Time UUID cursor")

}

func (w *WorkoutProgram) DBCursor() cursor.Cursor {
	return cursor.NewTimeUUIDCursor(w.CreatedAt, w.ID)
}

func (w *WorkoutProgram) Cursor() string {
	return w.DBCursor().Serialize()
}

func (w *WorkoutProgram) Node() *WorkoutProgram {
	return w
}
