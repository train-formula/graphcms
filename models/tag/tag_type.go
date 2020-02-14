package tag

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

type TagType uint8

const (
	UnknownTagType TagType = iota
	WorkoutProgramTagType
	WorkoutCategoryTagType
	ExerciseTagType
	WorkoutTagType
	PrescriptionTagType
	PlanTagType
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
	case PrescriptionTagType:
		return "PRESCRIPTION"
	case PlanTagType:
		return "PLAN"
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
	case PrescriptionTagType.String():
		return PrescriptionTagType
	case PlanTagType.String():
		return PlanTagType
	}

	return UnknownTagType
}

var _ sql.Scanner = (*TagType)(nil)
var _ driver.Valuer = UnknownTagType

func (t TagType) Value() (driver.Value, error) {

	return t.String(), nil
}

func (t *TagType) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	switch src.(type) {
	case string:
		parsed := ParseTagType(src.(string))
		if parsed == UnknownTagType {
			return errors.New("Unknown tag type")
		}
		*t = parsed
		return nil
	case []byte:
		srcCopy := make([]byte, len(src.([]byte)))
		copy(srcCopy, src.([]byte))
		parsed := ParseTagType(string(srcCopy))
		if parsed == UnknownTagType {
			return errors.New("Unknown tag type")
		}
		*t = parsed
		return nil
	}

	return fmt.Errorf("cannot scan %T", src)
}
