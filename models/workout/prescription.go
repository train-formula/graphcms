package workout

import (
	"errors"
	"time"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database/cursor"
	"github.com/train-formula/graphcms/database/types"
)

type Prescription struct {
	tableName             struct{}  `sql:"workout.prescription"`
	ID                    uuid.UUID `json:"id"`
	CreatedAt             time.Time `json:"createdAt" pgxload:"defaultZero"`
	UpdatedAt             time.Time `json:"updatedAt" pgxload:"defaultZero"`
	TrainerOrganizationID uuid.UUID `json:"trainerOrganizationID"`
	Name                  string    `json:"name"`
	PrescriptionCategory  string    `json:"prescriptionCategory"`

	DurationSeconds types.NullInt64 `json:"durationSeconds"`
}

func (p Prescription) TableName() string {
	return "workout.prescription"
}

// Generate an SQL query with for cursor that paginates with columns from this table
// Also provide the params to go with it
func (p Prescription) CursorQuery(prefix string, c cursor.Cursor) (string, []interface{}, error) {

	if conv, proper := c.(*cursor.TimeUUIDCursor); proper {

		q, p := conv.GenerateAscendingSQLConditions(prefix, "created_at", "uuid")

		return q, p, nil
	}

	return "", nil, errors.New("cursor must be a Time UUID cursor")

}

// Generated a database query for the particular row represented
func (p *Prescription) DBCursor() cursor.Cursor {
	return cursor.NewTimeUUIDCursor(p.CreatedAt, p.ID)
}

// Serializes the result of DBCursor
// Necessary for gqlgen, allows us to avoid creating a seperate PrescriptionEdge
func (p *Prescription) Cursor() string {
	return p.DBCursor().Serialize()
}

// Necessary for gqlgen, allows us to avoid creating a seperate PrescriptionEdge
func (p *Prescription) Node() *Prescription {
	return p
}
