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

type BlockType uint8

const (
	UnknownBlockType BlockType = iota
	GeneralBlockType
	RoundBlockType
	TimedRoundBlockType
)

func (t BlockType) String() string {

	switch t {
	case GeneralBlockType:
		return "GENERAL"
	case RoundBlockType:
		return "ROUND"
	case TimedRoundBlockType:
		return "TIMED_ROUND"
	}

	return "UNKNOWN"
}

func ParseBlockType(s string) BlockType {

	switch strings.ToUpper(strings.TrimSpace(s)) {
	case GeneralBlockType.String():
		return GeneralBlockType
	case RoundBlockType.String():
		return RoundBlockType
	case TimedRoundBlockType.String():
		return TimedRoundBlockType

	}

	return UnknownBlockType
}

var _ sql.Scanner = (*BlockType)(nil)
var _ driver.Valuer = UnknownBlockType

func (t BlockType) Value() (driver.Value, error) {

	return t.String(), nil
}

func (t *BlockType) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	switch src.(type) {
	case string:
		parsed := ParseBlockType(src.(string))
		if parsed == UnknownBlockType {
			return errors.New("Unknown block type")
		}
		*t = parsed
		return nil
	case []byte:
		srcCopy := make([]byte, len(src.([]byte)))
		copy(srcCopy, src.([]byte))
		parsed := ParseBlockType(string(srcCopy))
		if parsed == UnknownBlockType {
			return errors.New("Unknown block type")
		}
		*t = parsed
		return nil
	}

	return fmt.Errorf("cannot scan %T", src)
}

func (e *BlockType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ParseBlockType(str)
	if *e == UnknownBlockType {
		return fmt.Errorf("%s is not a valid BlockType", str)
	}
	return nil
}

func (e BlockType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
