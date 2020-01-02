package validation

import (
	"github.com/asaskevich/govalidator"
	"github.com/vektah/gqlparser/gqlerror"
)

func CheckStringNilOrIsNotEmpty(s *string, message string) ValidatorFunc {

	if s == nil {
		return EmptyValidatorFunc
	}

	return CheckStringMinimumLength(*s, 1, message)
}

func CheckStringNilOrIsURL(s *string, message string) ValidatorFunc {
	if s == nil {
		return EmptyValidatorFunc
	}

	return CheckStringIsURL(*s, message)
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

func CheckStringIsURL(s string, message string) ValidatorFunc {

	return func() *gqlerror.Error {
		if len(s) == 0 {
			return gqlerror.Errorf(message)
		}

		if !govalidator.IsURL(s) {
			return gqlerror.Errorf(message)
		}

		return nil
	}
}
