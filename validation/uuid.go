package validation

import (
	"github.com/gofrs/uuid"
	"github.com/vektah/gqlparser/gqlerror"
)

func CheckIsUUID(s string, message string) ValidatorFunc {

	return func() *gqlerror.Error {

		_, err := uuid.FromString(s)
		if err != nil {
			return gqlerror.Errorf(message)
		}

		return nil
	}
}
