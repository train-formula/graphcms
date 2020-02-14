package workout

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type ProgramLevel uint8

const (
	UnknownProgramLevel ProgramLevel = iota
	BeginnerProgramLevel
	IntermediateProgramLevel
	AdvancedProgramLevel
)

func (t ProgramLevel) String() string {

	switch t {
	case BeginnerProgramLevel:
		return "BEGINNER"
	case IntermediateProgramLevel:
		return "INTERMEDIATE"
	case AdvancedProgramLevel:
		return "ADVANCED"
	}

	return "UNKNOWN"
}

func ParseProgramLevel(s string) ProgramLevel {

	switch strings.ToUpper(strings.TrimSpace(s)) {
	case BeginnerProgramLevel.String():
		return BeginnerProgramLevel
	case IntermediateProgramLevel.String():
		return IntermediateProgramLevel
	case AdvancedProgramLevel.String():
		return AdvancedProgramLevel

	}

	return UnknownProgramLevel
}

var _ sql.Scanner = (*ProgramLevel)(nil)
var _ driver.Valuer = UnknownProgramLevel

func (t ProgramLevel) Value() (driver.Value, error) {

	return t.String(), nil
}

func (t *ProgramLevel) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	switch src.(type) {
	case string:
		parsed := ParseProgramLevel(src.(string))
		if parsed == UnknownProgramLevel {
			return errors.New("Unknown program level")
		}
		*t = parsed
		return nil
	case []byte:
		srcCopy := make([]byte, len(src.([]byte)))
		copy(srcCopy, src.([]byte))
		parsed := ParseProgramLevel(string(srcCopy))
		if parsed == UnknownProgramLevel {
			return errors.New("Unknown program level")
		}
		*t = parsed
		return nil
	}

	return fmt.Errorf("cannot scan %T", src)
}

func (e *ProgramLevel) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ParseProgramLevel(str)
	if *e == UnknownProgramLevel {
		return fmt.Errorf("%s is not a valid program level", str)
	}
	return nil
}

func (e ProgramLevel) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
