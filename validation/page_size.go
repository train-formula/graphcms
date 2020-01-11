package validation

import "github.com/vektah/gqlparser/gqlerror"

func CheckPageSize(input, min, max int) ValidatorFunc {

	return func() *gqlerror.Error {
		if input > max {
			return gqlerror.Errorf("request page size too large, must be at most %d", max)
		}

		if input < min {
			return gqlerror.Errorf("request page size too small, must be at least %d", min)
		}

		return nil
	}

}

func DefaultCheckPageSize(input int) ValidatorFunc {

	return CheckPageSize(input, 1, 200)

}
