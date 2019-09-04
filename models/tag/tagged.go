package tag

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Tagged struct {
	tableName struct{} `pg:",discard_unknown_columns"`

	ID                    uuid.UUID
	CreatedAt             time.Time
	UpdatedAt             time.Time
	TagUUID               uuid.UUID
	TrainerOrganizationID uuid.UUID
	TaggedUUID            uuid.UUID
	TagType               TagType `sql:"type:tag.tag_type"`
}

func (t Tagged) TableName() string {
	return "tag.tagged"
}
