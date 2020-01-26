package validation

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/gqlerror"
)

type ValidatorFunc func() *gqlerror.Error

// A ValidatorFunc that returns nothing
func EmptyValidatorFunc() *gqlerror.Error {
	return nil
}

// Generates a ValidatorFunc that always return the given message as an error
func ImmediateErrorValidator(message string, args ...interface{}) ValidatorFunc {
	return func() *gqlerror.Error {
		return gqlerror.Errorf(message, args...)
	}
}

func ValidationChain(ctx context.Context, funcs ...ValidatorFunc) bool {

	ok := true
	for _, f := range funcs {

		err := f()

		if err != nil {
			graphql.AddError(ctx, err)
			ok = false
		}

	}

	return ok
}
