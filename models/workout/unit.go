package workout

import (
	"time"

	"github.com/gofrs/uuid"
)

type Unit struct {
	tableName struct{}  `sql:"workout.unit"`
	ID        uuid.UUID `json:"id"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	Name               string `json:"name"`
	NameMedium         string `json:"nameMedium"`
	NameShort          string `json:"nameShort"`
	RepresentsTime     bool   `json:"representsTime"`
	RepresentsWeight   bool   `json:"representsWeight"`
	RepresentsCounter  bool   `json:"representsCounter"`
	RepresentsDistance bool   `json:"representsDistance"`
}

func (u Unit) TableName() string {
	return "workout.unit"
}
