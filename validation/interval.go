package validation

import (
	"fmt"

	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/models/interval"
	"github.com/vektah/gqlparser/gqlerror"
)

// Validate a DiurnalIntervalInput
// Allows verifying that the specified interval count is >= minimumInterval
// Allows specification of optional fieldName, which will be appended to any error messages returned by the generated validator
func CheckDiurnalIntervalInput(in generated.DiurnalIntervalInput, minimumInterval int, fieldName *string) ValidatorFunc {

	buildMessage := func(message string) string {
		if fieldName != nil {
			return fmt.Sprintf("%s (%s)", message, *fieldName)
		}

		return message
	}

	return func() *gqlerror.Error {
		if in.Interval == nil {
			return gqlerror.Errorf(buildMessage("Missing interval type"))
		}

		if *in.Interval == interval.UnknownDiurnalInterval {
			return gqlerror.Errorf(buildMessage("Unknown interval type"))
		}

		if in.Count < minimumInterval {
			return gqlerror.Errorf(buildMessage(fmt.Sprintf("Interval must be >= %d", minimumInterval)))
		}

		return nil
	}
}
