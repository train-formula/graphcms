package workout

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/go-pg/pg/v9/types"
)

type CategoryType uint8

const (
	UnknownCategoryType CategoryType = iota
	GeneralCategoryType
	RoundCategoryType
	TimedRoundCategoryType
)

func (t CategoryType) String() string {

	switch t {
	case GeneralCategoryType:
		return "GENERAL"
	case RoundCategoryType:
		return "ROUND"
	case TimedRoundCategoryType:
		return "TIMED_ROUND"
	}

	return "UNKNOWN"
}

func ParseCategoryType(s string) CategoryType {

	switch strings.ToUpper(strings.TrimSpace(s)) {
	case GeneralCategoryType.String():
		return GeneralCategoryType
	case RoundCategoryType.String():
		return RoundCategoryType
	case TimedRoundCategoryType.String():
		return TimedRoundCategoryType

	}

	return UnknownCategoryType
}

var _ types.ValueAppender = (*CategoryType)(nil)
var _ types.ValueScanner = (*CategoryType)(nil)

func (t *CategoryType) AppendValue(b []byte, flags int) ([]byte, error) {

	if flags == 1 {
		b = append(b, '\'')
	}
	b = append(b, t.String()...)
	if flags == 1 {
		b = append(b, '\'')
	}

	return b, nil
}

func (t *CategoryType) ScanValue(rd types.Reader, n int) error {
	if n <= 0 {
		return nil
	}

	tmp, err := rd.ReadFull()
	if err != nil {
		return err
	}

	parsed := ParseCategoryType(string(tmp))

	if parsed == UnknownCategoryType {
		return errors.New("Unknown category type")
	}

	*t = parsed

	return nil
}

func (e *CategoryType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ParseCategoryType(str)
	if *e == UnknownCategoryType {
		return fmt.Errorf("%s is not a valid CategoryType", str)
	}
	return nil
}

func (e CategoryType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
