package validation

import (
	"context"
	"fmt"

	"github.com/train-formula/graphcms/generated"
	"github.com/willtrking/pgxload"
)

// Returns an array of ValidatorFunc's that are used to validate a CreatePrescriptionSetData
// Takes an optional index which may be specified to provide more specific error messages if the
// CreatePrescriptionSetData was from an array (e.g. from a CreatePrescription request)
func CheckCreatePrescriptionSetData(ctx context.Context, conn pgxload.PgxLoader, request generated.CreatePrescriptionSetData, idx *int) []ValidatorFunc {

	buildMessage := func(message string) string {

		if idx != nil {
			return fmt.Sprintf("%s (at index %d)", message, *idx)
		}

		return message
	}

	return []ValidatorFunc{
		CheckIntGT(request.SetNumber, 0, buildMessage("Set number must be > 0")),
		CheckUnitDataValidAndNotNil(ctx, conn, request.PrimaryParameter, buildMessage("primaryParameter")),
		CheckUnitDataValidOrNil(ctx, conn, request.SecondaryParameter, buildMessage("primaryParameter")),
	}
}
