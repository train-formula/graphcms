package workout

import (
	"errors"
	"time"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database/cursor"
)

type PrescriptionSet struct {
	tableName             struct{}  `sql:"workout.prescription_set"`
	ID                    uuid.UUID `json:"id"`
	CreatedAt             time.Time `json:"createdAt"`
	UpdatedAt             time.Time `json:"updatedAt"`
	TrainerOrganizationID uuid.UUID `json:"trainerOrganizationID"`
	PrescriptionID        uuid.UUID `json:"prescriptionID"`

	SetNumber int `json:"setNumber"`

	RepNumeral *int
	RepText    *string
	RepUnitID  uuid.UUID

	RepModifierNumeral *int
	RepModifierText    *string
	RepModifierUnitID  *uuid.UUID

	Order int `json:"order"`
}

func (p PrescriptionSet) TableName() string {
	return "workout.prescription_set"
}

// Generate an SQL query with for cursor that paginates with columns from this table
// Also provide the params to go with it
func (p PrescriptionSet) CursorQuery(prefix string, c cursor.Cursor) (string, []interface{}, error) {

	if conv, proper := c.(*cursor.TimeUUIDCursor); proper {

		q, p := conv.GenerateAscendingSQLConditions(prefix, "created_at", "uuid")

		return q, p, nil
	}

	return "", nil, errors.New("cursor must be a Time UUID cursor")

}

// Generated a database query for the particular row represented
func (p *PrescriptionSet) DBCursor() cursor.Cursor {
	return cursor.NewTimeUUIDCursor(p.CreatedAt, p.ID)
}

// Serializes the result of DBCursor
// Necessary for gqlgen, allows us to avoid creating a seperate WorkoutCategoryEdge
func (p *PrescriptionSet) Cursor() string {
	return p.DBCursor().Serialize()
}

// Necessary for gqlgen, allows us to avoid creating a seperate WorkoutCategoryEdge
func (p *PrescriptionSet) Node() *PrescriptionSet {
	return p
}
