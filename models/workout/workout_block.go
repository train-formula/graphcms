package workout

import (
	"errors"
	"time"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database/cursor"
	"github.com/train-formula/graphcms/database/types"
)

type WorkoutBlock struct {
	tableName             struct{}  `sql:"workout.category"`
	ID                    uuid.UUID `json:"id"`
	CreatedAt             time.Time `json:"createdAt" pgxload:"defaultZero"`
	UpdatedAt             time.Time `json:"updatedAt" pgxload:"defaultZero"`
	TrainerOrganizationID uuid.UUID `json:"trainerOrganizationID"`
	WorkoutCategoryID     uuid.UUID `json:"workoutCategoryID"`
	CategoryOrder         int       `json:"categoryOrder"`
	RoundNumeral          types.NullInt64
	RoundText             types.NullString
	RoundUnitID           *uuid.UUID
	DurationSeconds       types.NullInt64 `json:"durationSeconds"`
	RoundRestDuration     types.NullInt64 `json:"roundRestDuration"`
	NumberOfRounds        types.NullInt64 `json:"numberOfRounds"`
}

func (w WorkoutBlock) TableName() string {
	return "workout.block"
}

// Generate an SQL query with for cursor that paginates with columns from this table
// Also provide the params to go with it
func (w WorkoutBlock) CursorQuery(prefix string, c cursor.Cursor) (string, []interface{}, error) {

	if conv, proper := c.(*cursor.TimeUUIDCursor); proper {

		q, p := conv.GenerateAscendingSQLConditions(prefix, "created_at", "uuid")

		return q, p, nil
	}

	return "", nil, errors.New("cursor must be a Time UUID cursor")

}

// Generated a database query for the particular row represented
func (w *WorkoutBlock) DBCursor() cursor.Cursor {
	return cursor.NewTimeUUIDCursor(w.CreatedAt, w.ID)
}
