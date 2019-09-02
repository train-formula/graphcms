package tag

import (
	"time"

	"github.com/gofrs/uuid"
)

type Tag struct {
	tableName struct{} `pg:",discard_unknown_columns"`

	ID                    uuid.UUID `json:"id"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	Tag                   string    `json:"tag"`
	Description           string    `json:"description"`
	TrainerOrganizationID uuid.UUID `json:"trainer_organization_id"`
}

func (t Tag) TableName() string {
	return "tag.tags"
}
