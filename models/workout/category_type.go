package workout

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/go-pg/pg/v9/types"
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

var _ types.ValueAppender = (*BlockType)(nil)
var _ types.ValueScanner = (*BlockType)(nil)

func (t *BlockType) AppendValue(b []byte, flags int) ([]byte, error) {

	if flags == 1 {
		b = append(b, '\'')
	}
	b = append(b, t.String()...)
	if flags == 1 {
		b = append(b, '\'')
	}

	return b, nil
}

func (t *BlockType) ScanValue(rd types.Reader, n int) error {
	if n <= 0 {
		return nil
	}

	tmp, err := rd.ReadFull()
	if err != nil {
		return err
	}

	parsed := ParseBlockType(string(tmp))

	if parsed == UnknownBlockType {
		return errors.New("Unknown block type")
	}

	*t = parsed

	return nil
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
