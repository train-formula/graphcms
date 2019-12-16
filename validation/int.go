package validation

import "github.com/vektah/gqlparser/gqlerror"

func CheckIntIsNilOrGT(i *int, gt int, message string) ValidatorFunc {
	if i == nil {
		return EmptyValidatorFunc
	}

	return CheckIntGT(*i, gt, message)
}

func CheckIntIsNilOrGTE(i *int, gt int, message string) ValidatorFunc {
	if i == nil {
		return EmptyValidatorFunc
	}

	return CheckIntGTE(*i, gt, message)
}

func CheckIntGT(i int, gt int, message string) ValidatorFunc {
	return func() *gqlerror.Error {
		if i > gt {
			return nil
		}

		return gqlerror.Errorf(message)
	}
}

func CheckIntGTE(i int, gt int, message string) ValidatorFunc {
	return func() *gqlerror.Error {
		if i >= gt {
			return nil
		}

		return gqlerror.Errorf(message)
	}
}
