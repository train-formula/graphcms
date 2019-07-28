package workout

import (
	"fmt"
	"io"
	"strconv"
)

type WorkoutProgramType string

const (
	WorkoutProgramTypeRelative WorkoutProgramType = "RELATIVE"
	WorkoutProgramTypeStatic   WorkoutProgramType = "STATIC"
)

var AllWorkoutProgramType = []WorkoutProgramType{
	WorkoutProgramTypeRelative,
	WorkoutProgramTypeStatic,
}

func (e WorkoutProgramType) IsValid() bool {
	switch e {
	case WorkoutProgramTypeRelative, WorkoutProgramTypeStatic:
		return true
	}
	return false
}

func (e WorkoutProgramType) String() string {
	return string(e)
}

func (e *WorkoutProgramType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = WorkoutProgramType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid WorkoutProgramType", str)
	}
	return nil
}

func (e WorkoutProgramType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
