package workout

import (
	"fmt"
	"io"
	"strconv"
)

type CategoryType string

const (
	CategoryTypeGeneral    CategoryType = "GENERAL"
	CategoryTypeRound      CategoryType = "ROUND"
	CategoryTypeTimedRound CategoryType = "TIMED_ROUND"
)

var AllCategoryType = []CategoryType{
	CategoryTypeGeneral,
	CategoryTypeRound,
	CategoryTypeTimedRound,
}

func (e CategoryType) IsValid() bool {
	switch e {
	case CategoryTypeGeneral, CategoryTypeRound, CategoryTypeTimedRound:
		return true
	}
	return false
}

func (e CategoryType) String() string {
	return string(e)
}

func (e *CategoryType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CategoryType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CategoryType", str)
	}
	return nil
}

func (e CategoryType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
