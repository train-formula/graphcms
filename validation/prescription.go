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
			return message + fmt.Sprintf(" (at index %d", *idx)
		}

		return message
	}

	repModifierUnitMessage := ""
	if request.RepModifierUnitID != nil {
		repModifierUnitMessage = buildMessage(fmt.Sprintf("Rep modfier unit ID %s does not exist", (*request.RepModifierUnitID).String()))
	}

	return []ValidatorFunc{
		CheckIntGT(request.SetNumber, 0, buildMessage("Set number must be > 0")),
		CheckIntIsNilOrGT(request.RepNumeral, 0, buildMessage("If specified, rep numeral must be > 0")),
		CheckStringNilOrIsNotEmpty(request.RepText, buildMessage("If specified, rep text must not be empty"), true),
		CheckIntIsNilOrGT(request.RepModifierNumeral, 0, buildMessage("If specified, rep modifier numeral must be > 0")),
		CheckStringNilOrIsNotEmpty(request.RepModifierText, buildMessage("If specified, rep modifier text must not be empty"), true),
		UnitExists(ctx, conn, request.RepUnitID, buildMessage(fmt.Sprintf("Rep unit ID %s does not exist", request.RepUnitID.String()))),
		UnitIsNilOrExists(ctx, conn, request.RepModifierUnitID, repModifierUnitMessage),
	}
}

/*




	RepModifierUnitID  *uuid.UUID `json:"repModifierUnitID"`
*/
