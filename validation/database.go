package validation

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database"
	"github.com/train-formula/graphcms/database/tagdb"
	"github.com/train-formula/graphcms/database/trainerdb"
	"github.com/train-formula/graphcms/database/workoutdb"
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
				return gqlerror.Errorf("tag %s does not belong to the same trainer", id)
			}
		} else {
			return gqlerror.Errorf("tag %s does not exist", id)
		}
	}

	return nil
}

// Validates that a unid ID is either nil, or exists in the database
func UnitIsNilOrExists(ctx context.Context, conn database.Conn, unitID *uuid.UUID) ValidatorFunc {

	return func() *gqlerror.Error {
		if unitID == nil {
			return nil
		}

		searchID := *unitID
		_, err := workoutdb.GetUnit(ctx, conn, searchID)

		if err != nil {
			if err == pg.ErrNoRows {
				return gqlerror.Errorf("Unit ID %s does not exist", searchID)
			}
			return gqlerror.Errorf(err.Error())
		}

		return nil
	}

}

// Validates that a organization ID is either nil, or exists in the database
func OrganizationExists(ctx context.Context, conn database.Conn, organizationID uuid.UUID) ValidatorFunc {

	return func() *gqlerror.Error {

		_, err := trainerdb.GetOrganization(ctx, conn, organizationID)

		if err != nil {
			if err == pg.ErrNoRows {
				return gqlerror.Errorf("Organization ID %s does not exist", organizationID)
			}
			return gqlerror.Errorf(err.Error())
		}

		return nil
	}

}
