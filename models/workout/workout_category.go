package workout

import (
	"errors"
	"time"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database/cursor"
)

type WorkoutCategory struct {
	tableName             struct{}     `sql:"workout.category"`
	ID                    uuid.UUID    `json:"id"`
	CreatedAt             time.Time    `json:"createdAt"`
	UpdatedAt             time.Time    `json:"updatedAt"`
	TrainerOrganizationID uuid.UUID    `json:"trainerOrganizationID"`
	Name                  string       `json:"name"`
	Description           string       `json:"description"`
	Type                  CategoryType `json:"type"`
	RoundNumeral          *int32       `json:"roundNumeral"`
	RoundText             *string      `json:"roundText"`
	RoundUnitID           *uuid.UUID   `json:"roundUnitID"`
	DurationSeconds       *int32       `json:"durationSeconds"`
}

func (w WorkoutCategory) TableName() string {
	return "workout.category"
}

// Generate an SQL query with for cursor that paginates with columns from this table
// Also provide the params to go with it
func (w WorkoutCategory) CursorQuery(prefix string, c cursor.Cursor) (string, []interface{}, error) {

	if conv, proper := c.(*cursor.TimeUUIDCursor); proper {

		q, p := conv.GenerateAscendingSQLConditions(prefix, "created_at", "uuid")

		return q, p, nil
	}

	return "", nil, errors.New("cursor must be a Time UUID cursor")

}

// Generated a database query for the particular row represented
func (w *WorkoutCategory) DBCursor() cursor.Cursor {
	return cursor.NewTimeUUIDCursor(w.CreatedAt, w.ID)
}

// Serializes the result of DBCursor
// Necessary for gqlgen, allows us to avoid creating a seperate WorkoutCategoryEdge
func (w *WorkoutCategory) Cursor() string {
	return w.DBCursor().Serialize()
}

// Necessary for gqlgen, allows us to avoid creating a seperate WorkoutCategoryEdge
func (w *WorkoutCategory) Node() *WorkoutCategory {
	return w
}
