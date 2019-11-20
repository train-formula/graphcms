package tag

import (
	"time"

	"github.com/gofrs/uuid"
)

type Tagged struct {
	tableName struct{} `pg:",discard_unknown_columns"`

	ID                    uuid.UUID
	CreatedAt             time.Time
	UpdatedAt             time.Time
	TagID                 uuid.UUID
	TrainerOrganizationID uuid.UUID
	TaggedID              uuid.UUID
	TagType               TagType `sql:"type:tag.tag_type"`
}

func (t Tagged) TableName() string {
	return "tag.tagged"
}
