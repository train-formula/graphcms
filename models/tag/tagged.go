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
}

func (t Tagged) TableName() string {
	return "tag.tagged"
}

/*
CREATE TABLE "tag"."tagged" (
  "id" uuid NOT NULL,
  "created_at" timestamp without time zone NOT NULL DEFAULT NOW(),
  "updated_at" timestamp without time zone NOT NULL DEFAULT NOW(),
  "tag_uuid" uuid NOT NULL REFERENCES trainer.organization(id) DEFERRABLE INITIALLY DEFERRED,
  "trainer_organization_id" uuid NOT NULL REFERENCES trainer.organization(id) DEFERRABLE INITIALLY DEFERRED,
  "tag_on" tag.taggable NOT NULL,
  "tagged_uuid" uuid NOT NULL,
  PRIMARY KEY ("id")
);
*/
