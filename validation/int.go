package validation

import (
	"math"

	"github.com/vektah/gqlparser/gqlerror"
)

func CheckIntIsNilOrGT(i *int, gt int, message string) ValidatorFunc {
	if i == nil {
		return EmptyValidatorFunc
	}

	return CheckIntGT(*i, gt, message)
}

func CheckIntIsNilOrGTE(i *int, gte int, message string) ValidatorFunc {
	if i == nil {
		return EmptyValidatorFunc
	}

	return CheckIntGTE(*i, gte, message)
}

// Check that an int fits in uint8 bounds (0, 255) or is nil
func CheckIntIsNilOrUint8(i *int, message string) ValidatorFunc {
	if i == nil {
		return EmptyValidatorFunc
	}

	return CheckIntUint8(*i, message)
}

func CheckIntGT(i int, gt int, message string) ValidatorFunc {
	return func() *gqlerror.Error {
		if i > gt {
			return nil
		}

		return gqlerror.Errorf(message)
	}
}

func CheckIntGTE(i int, gte int, message string) ValidatorFunc {
	return func() *gqlerror.Error {
		if i >= gte {
			return nil
		}

		return gqlerror.Errorf(message)
	}
}

// Check that an int fits in uint8 bounds (0, 255)
func CheckIntUint8(i int, message string) ValidatorFunc {
	return func() *gqlerror.Error {

		if i < 0 || i > math.MaxUint8 {
			return gqlerror.Errorf(message)
		}

		return nil
	}
}
