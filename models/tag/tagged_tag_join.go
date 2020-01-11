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

	valued := database.ReflectValue(Tag{})

	columns := database.StructColumns(valued, tagTablePrefix)
	columns = append(columns, database.PGPrefixedColumn("tagged_id", taggedTablePrefix))
	columns = append(columns, database.PGPrefixedColumn("tag_type", taggedTablePrefix))

	return strings.Join(columns, ",")
}
