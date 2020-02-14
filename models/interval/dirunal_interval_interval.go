package interval

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
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

var _ sql.Scanner = (*DiurnalIntervalInterval)(nil)
var _ driver.Valuer = UnknownDiurnalInterval

func (t DiurnalIntervalInterval) Value() (driver.Value, error) {

	return t.String(), nil
}

func (t *DiurnalIntervalInterval) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	switch src.(type) {
	case string:
		parsed := ParseDiurnalIntervalInterval(src.(string))
		if parsed == UnknownDiurnalInterval {
			return errors.New("Unknown diurnal interval interval")
		}
		*t = parsed
		return nil
	case []byte:
		srcCopy := make([]byte, len(src.([]byte)))
		copy(srcCopy, src.([]byte))
		parsed := ParseDiurnalIntervalInterval(string(srcCopy))
		if parsed == UnknownDiurnalInterval {
			return errors.New("Unknown diurnal interval interval")
		}
		*t = parsed
		return nil
	}

	return fmt.Errorf("cannot scan %T", src)
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
