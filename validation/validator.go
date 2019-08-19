package validation

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/gqlerror"
)

type ValidatorFunc func() *gqlerror.Error

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
