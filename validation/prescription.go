package validation

import (
	"context"
	"fmt"

	"github.com/train-formula/graphcms/database"
	"github.com/train-formula/graphcms/generated"
)

// Returns an array of ValidatorFunc's that are used to validate a CreatePrescriptionSetWithPrescription
// Takes an optional index which may be specified to provide more specific error messages if the
// CreatePrescriptionSetWithPrescription was from an array (e.g. from a CreatePrescription request)
func CheckCreatePrescriptionSetWithPrescription(ctx context.Context, conn database.Conn, request generated.CreatePrescriptionSetWithPrescription, idx *int) []ValidatorFunc {

	buildMessage := func(message string) string {

		if idx != nil {
			return message + fmt.Sprintf(" (at index %d)", *idx)
		}

		return message
	}

	return []ValidatorFunc{
		CheckIntGT(request.SetNumber, 0, buildMessage("Set number must be > 0")),
		CheckUnitDataValidAndNotNil(ctx, conn, request.PrimaryParameter, buildMessage("primaryParameter")),
		CheckUnitDataValidOrNil(ctx, conn, request.SecondaryParameter, buildMessage("primaryParameter")),
	}
}
