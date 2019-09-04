package validation

import "github.com/vektah/gqlparser/gqlerror"

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
