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

// Validates all tag UUIDs specified exist for a given trainer organization id
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

// Validates that a unit exists in the database
func UnitExists(ctx context.Context, conn database.Conn, unitID uuid.UUID, message string) ValidatorFunc {

	return func() *gqlerror.Error {

		_, err := workoutdb.GetUnit(ctx, conn, unitID)

		if err != nil {
			if err == pg.ErrNoRows {
				return gqlerror.Errorf(message)
			}
			return gqlerror.Errorf(err.Error())
		}

		return nil
	}

}

// Validates that a unit id is either nil, or exists in the database
func UnitIsNilOrExists(ctx context.Context, conn database.Conn, unitID *uuid.UUID, message string) ValidatorFunc {

	if unitID == nil {
		return EmptyValidatorFunc
	}

	return UnitExists(ctx, conn, *unitID, message)

}

// Validates that a organization id exists in the database
func OrganizationExists(ctx context.Context, conn database.Conn, organizationID uuid.UUID) ValidatorFunc {

	return func() *gqlerror.Error {

		_, err := trainerdb.GetOrganization(ctx, conn, organizationID)

		if err != nil {
			if err == pg.ErrNoRows {
				return gqlerror.Errorf("Organization id %s does not exist", organizationID)
			}
			return gqlerror.Errorf(err.Error())
		}

		return nil
	}

}

// Validates that a prescription id  exists in the database
func PrescriptionExists(ctx context.Context, conn database.Conn, prescriptionID uuid.UUID) ValidatorFunc {

	return func() *gqlerror.Error {

		_, err := workoutdb.GetPrescription(ctx, conn, prescriptionID)

		if err != nil {
			if err == pg.ErrNoRows {
				return gqlerror.Errorf("Prescription id %s does not exist", prescriptionID)
			}
			return gqlerror.Errorf(err.Error())
		}

		return nil
	}

}
