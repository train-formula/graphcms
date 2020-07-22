package tag

import (
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database"
)

type TaggedTagJoin struct {
	TagID                    uuid.UUID
	TagCreatedAt             time.Time
	TagUpdatedAt             time.Time
	TagTag                   string
	TagTrainerOrganizationID uuid.UUID
	// TaggedID from Tagged table, NOT id field, hence TaggedTagged
	TaggedTaggedID uuid.UUID
	TaggedTagType  TagType `sql:"type:tag.tag_type"`
}

// Extract a tag struct from this result
func (t *TaggedTagJoin) Tag() *Tag {
	return &Tag{
		ID:                    t.TagID,
		CreatedAt:             t.TagCreatedAt,
		UpdatedAt:             t.TagUpdatedAt,
		Tag:                   t.TagTag,
		TrainerOrganizationID: t.TagTrainerOrganizationID,
	}
}

// Extract the columns to use in a SELECT statement
func (t TaggedTagJoin) SelectColumns(tagTablePrefix, taggedTablePrefix string) string {

	columns := []string{
		database.PGPrefixedColumn("id", tagTablePrefix) + " AS tag_id",
		database.PGPrefixedColumn("created_at", tagTablePrefix) + " AS tag_created_at",
		database.PGPrefixedColumn("updated_at", tagTablePrefix) + " AS tag_updated_at",
		database.PGPrefixedColumn("tag", tagTablePrefix) + " AS tag_tag",
		database.PGPrefixedColumn("trainer_organization_id", tagTablePrefix) + " AS tag_trainer_organization_id",
		database.PGPrefixedColumn("trainer_organization_id", tagTablePrefix) + " AS tag_trainer_organization_id",
		database.PGPrefixedColumn("tagged_id", taggedTablePrefix) + " AS tagged_tagged_id",
		database.PGPrefixedColumn("tag_type", taggedTablePrefix) + " AS tagged_tag_type",
	}

	return strings.Join(columns, ",")
}
