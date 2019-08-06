package trainer

import (
	"time"

	"github.com/satori/go.uuid"
)

type Organization struct {
	tableName struct{} `pg:",discard_unknown_columns"`

	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name string `json:"name"`
	Description  string `json:"description"`
}

func (o Organization) TableName() string {
	return "trainer.organization"
}