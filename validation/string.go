package validation

import "github.com/vektah/gqlparser/gqlerror"

func CheckStringNilOrIsNotEmpty(s *string, message string) ValidatorFunc {

	if s == nil {
		return EmptyValidatorFunc
	}

	return CheckStringMinimumLength(*s, 1, message)
}

func CheckStringIsNotEmpty(s string, message string) ValidatorFunc {

	return CheckStringMinimumLength(s, 1, message)
}

func CheckStringMinimumLength(s string, minLength int, message string) ValidatorFunc {

	return func() *gqlerror.Error {

		if len(s) < minLength {
			return gqlerror.Errorf(message)
		}

		return nil
	}
}

func CheckIntGT(i int, gt int, message string) ValidatorFunc {
	return func() *gqlerror.Error {
		if i < gt {
			return gqlerror.Errorf(message)
		}

		return nil
	}
}
