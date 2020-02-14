package validation

import (
	"context"
	"fmt"

	"github.com/train-formula/graphcms/generated"
	"github.com/vektah/gqlparser/gqlerror"
	"github.com/willtrking/pgxload"
)

func CheckUnitDataValid(ctx context.Context, conn pgxload.PgxLoader, data generated.AttachUnitData, unitName string) ValidatorFunc {

	return func() *gqlerror.Error {

		if data.Text == nil && data.Numeral == nil {
			return gqlerror.Errorf("%s requires numeral and/or text, missing both", unitName)
		}

		return UnitExists(ctx, conn, data.UnitID, fmt.Sprintf("unit ID %s does not exist for %s", data.UnitID, unitName))()

	}

}

func CheckUnitDataValidAndNotNil(ctx context.Context, conn pgxload.PgxLoader, data *generated.AttachUnitData, unitName string) ValidatorFunc {

	if data == nil {
		return func() *gqlerror.Error {
			return gqlerror.Errorf("%s is required", unitName)
		}
	}

	return CheckUnitDataValid(ctx, conn, *data, unitName)
}

func CheckUnitDataValidOrNil(ctx context.Context, conn pgxload.PgxLoader, data *generated.AttachUnitData, unitName string) ValidatorFunc {
	if data == nil {
		return EmptyValidatorFunc
	}

	return CheckUnitDataValid(ctx, conn, *data, unitName)
}
