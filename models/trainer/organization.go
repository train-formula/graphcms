package trainer

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

type Organization struct {
	tableName struct{} `pg:",discard_unknown_columns"`

	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at" pgxload:"defaultZero"`
	UpdatedAt   time.Time `json:"updated_at" pgxload:"defaultZero"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

func (o Organization) TableName() string {
	return "trainer.organization"
}
