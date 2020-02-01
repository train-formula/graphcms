package workout

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/go-pg/pg/v9/types"
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

var _ types.ValueAppender = (*ProgramLevel)(nil)
var _ types.ValueScanner = (*ProgramLevel)(nil)

func (t *ProgramLevel) AppendValue(b []byte, flags int) ([]byte, error) {

	if flags == 1 {
		b = append(b, '\'')
	}
	b = append(b, t.String()...)
	if flags == 1 {
		b = append(b, '\'')
	}

	return b, nil
}

func (t *ProgramLevel) ScanValue(rd types.Reader, n int) error {
	if n <= 0 {
		return nil
	}

	tmp, err := rd.ReadFull()
	if err != nil {
		return err
	}

	parsed := ParseProgramLevel(string(tmp))

	if parsed == UnknownProgramLevel {
		return errors.New("Unknown program level")
	}

	*t = parsed

	return nil
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
