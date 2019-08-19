package trainer

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

type Organization struct {
	tableName struct{} `pg:",discard_unknown_columns"`

	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (o Organization) TableName() string {
	return "trainer.organization"
}
