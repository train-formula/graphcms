package validation

import (
	"github.com/train-formula/graphcms/models/tag"
	"github.com/vektah/gqlparser/gqlerror"
)

func CheckTagTypeKnown(tg tag.TagType) ValidatorFunc {

	return func() *gqlerror.Error {
		if tg == tag.UnknownTagType {
			return gqlerror.Errorf("Unknown tag type")
		}

		return nil
	}
}
