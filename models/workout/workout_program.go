package workout

import (
	"errors"
	"time"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database/cursor"
)

type WorkoutProgram struct {
	tableName                struct{}   `sql:"workout.program"`
	ID                       uuid.UUID  `json:"id"`
	CreatedAt                time.Time  `json:"createdAt"`
	UpdatedAt                time.Time  `json:"updatedAt"`
	TrainerOrganizationID    uuid.UUID  `json:"trainerOrganizationID"`
	Name                     string     `json:"name"`
	Description              string     `json:"description" pg:",use_zero"`
	ExactStartDate           *time.Time `json:"exactStartDate"`
	StartsWhenCustomerStarts bool       `json:"startsWhenCustomerStarts"`
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

// Generated a database query for the particular row represented
func (w *WorkoutProgram) DBCursor() cursor.Cursor {
	return cursor.NewTimeUUIDCursor(w.CreatedAt, w.ID)
}

// Serializes the result of DBCursor
// Necessary for gqlgen, allows us to avoid creating a seperate WorkoutProgramEdge
func (w *WorkoutProgram) Cursor() string {
	return w.DBCursor().Serialize()
}

// Necessary for gqlgen, allows us to avoid creating a seperate WorkoutProgramEdge
func (w *WorkoutProgram) Node() *WorkoutProgram {
	return w
}
