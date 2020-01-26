package interval

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/go-pg/pg/v9/types"
)

type DiurnalIntervalInterval uint8

const (
	UnknownDiurnalInterval DiurnalIntervalInterval = iota
	DailyDiurnalInterval
	WeeklyDiurnalInterval
	MonthlyDiurnalInterval
	YearlyDiurnalInterval
)

func (t DiurnalIntervalInterval) String() string {

	switch t {
	case DailyDiurnalInterval:
		return "DAY"
	case WeeklyDiurnalInterval:
		return "WEEK"
	case MonthlyDiurnalInterval:
		return "MONTH"
	case YearlyDiurnalInterval:
		return "YEAR"
	}

	return "UNKNOWN"
}

func ParseDiurnalIntervalInterval(s string) DiurnalIntervalInterval {

	switch strings.ToUpper(strings.TrimSpace(s)) {
	case DailyDiurnalInterval.String():
		return DailyDiurnalInterval
	case WeeklyDiurnalInterval.String():
		return WeeklyDiurnalInterval
	case MonthlyDiurnalInterval.String():
		return MonthlyDiurnalInterval
	case YearlyDiurnalInterval.String():
		return YearlyDiurnalInterval
	}

	return UnknownDiurnalInterval
}

var _ types.ValueAppender = (*DiurnalIntervalInterval)(nil)
var _ types.ValueScanner = (*DiurnalIntervalInterval)(nil)

func (t *DiurnalIntervalInterval) AppendValue(b []byte, flags int) ([]byte, error) {

	if flags == 1 {
		b = append(b, '\'')
	}
	b = append(b, t.String()...)
	if flags == 1 {
		b = append(b, '\'')
	}

	return b, nil
}

func (t *DiurnalIntervalInterval) ScanValue(rd types.Reader, n int) error {
	if n <= 0 {
		return nil
	}

	tmp, err := rd.ReadFull()
	if err != nil {
		return err
	}

	parsed := ParseDiurnalIntervalInterval(string(tmp))

	if parsed == UnknownDiurnalInterval {
		return errors.New("Unknown diurnal interval interval")
	}

	*t = parsed

	return nil
}

func (t *DiurnalIntervalInterval) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	parsed := ParseDiurnalIntervalInterval(str)

	if parsed == UnknownDiurnalInterval {
		return fmt.Errorf("%s is not a valid DiurnalIntervalInterval", str)
	}

	*t = parsed

	return nil
}

func (t DiurnalIntervalInterval) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(t.String()))
}

func (t DiurnalIntervalInterval) IsValid() bool {
	return t != UnknownDiurnalInterval
}
