package tag

import (
	"errors"
	"strings"

	"github.com/go-pg/pg/v9/types"
)

type TagType uint8

const (
	UnknownTagType TagType = iota
	WorkoutProgramTagType
	WorkoutCategoryTagType
	ExerciseTagType
	WorkoutTagType
)

func (t TagType) String() string {

	switch t {
	case WorkoutProgramTagType:
		return "WORKOUT_PROGRAM"
	case WorkoutCategoryTagType:
		return "WORKOUT_CATEGORY"
	case ExerciseTagType:
		return "EXERCISE"
	case WorkoutTagType:
		return "WORKOUT"
	}

	return "UNKNOWN"
}

func ParseTagType(s string) TagType {

	switch strings.ToUpper(strings.TrimSpace(s)) {
	case WorkoutProgramTagType.String():
		return WorkoutProgramTagType
	case WorkoutCategoryTagType.String():
		return WorkoutCategoryTagType
	case ExerciseTagType.String():
		return ExerciseTagType
	case WorkoutTagType.String():
		return WorkoutTagType

	}

	return UnknownTagType
}

var _ types.ValueAppender = (*TagType)(nil)
var _ types.ValueScanner = (*TagType)(nil)

func (t *TagType) AppendValue(b []byte, flags int) ([]byte, error) {

	if flags == 1 {
		b = append(b, '\'')
	}
	b = append(b, t.String()...)
	if flags == 1 {
		b = append(b, '\'')
	}

	return b, nil
}

func (t *TagType) ScanValue(rd types.Reader, n int) error {
	if n <= 0 {
		return nil
	}

	tmp, err := rd.ReadFull()
	if err != nil {
		return err
	}

	parsed := ParseTagType(string(tmp))

	if parsed == UnknownTagType {
		return errors.New("Unknown tag type")
	}

	*t = parsed

	return nil
}
