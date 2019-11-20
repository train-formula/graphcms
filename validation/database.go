package validation

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database"
	"github.com/train-formula/graphcms/database/tagdb"
	"github.com/vektah/gqlparser/gqlerror"
)

// Validates all tag UUIDs specified exist for a given trainer organization ID
// Returns an error if any don't exist
func TagsAllExistForTrainer(ctx context.Context, conn database.Conn, trainerOrganizationID uuid.UUID, ids []uuid.UUID) error {

	tags, err := tagdb.GetTags(ctx, conn, ids)
	if err != nil {
		return err
	}

	mappedTagTrainers := make(map[uuid.UUID]uuid.UUID)

	for _, t := range tags {
		mappedTagTrainers[t.ID] = t.TrainerOrganizationID
	}

	for _, id := range ids {
		if foundTrainer, hasTag := mappedTagTrainers[id]; hasTag {
			if foundTrainer != trainerOrganizationID {
				return gqlerror.Errorf("Tag %s does not belong to the same trainer", id)
			}
		} else {
			return gqlerror.Errorf("Tag %s does not exist", id)
		}
	}

	return nil
}
