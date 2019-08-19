package cursor

import (
	"bytes"
	"encoding/base64"
	"time"

	uuid "github.com/gofrs/uuid"
	"github.com/vektah/gqlparser/gqlerror"
)

const timeFormat = time.RFC3339Nano

var cursorSep = []byte(":")
var b64Encoding = base64.URLEncoding

func DeserializeCursor(cursor *string) (*TimeUUIDCursor, error) {
	if cursor == nil {
		return nil, nil
	}

	decoded, err := b64Encoding.DecodeString(*cursor)
	if err != nil {
		return nil, gqlerror.Errorf("Invalid cursor")
	}

	split := bytes.SplitN(decoded, cursorSep, 2)

	if len(split) != 2 {
		return nil, gqlerror.Errorf("Invalid cursor size")
	}

	id, err := uuid.FromBytes(split[0])
	if err != nil {
		return nil, gqlerror.Errorf("Invalid cursor UUID")
	}

	tme, err := time.Parse(timeFormat, string(split[1]))
	if err != nil {
		return nil, gqlerror.Errorf("Invalid cursor time")
	}

	return &TimeUUIDCursor{
		Time: tme,
		UUID: id,
	}, nil
}

func NewTimeUUIDCursor(t time.Time, u uuid.UUID) *TimeUUIDCursor {
	return &TimeUUIDCursor{
		Time: t,
		UUID: u,
	}
}

type TimeUUIDCursor struct {
	Time time.Time
	UUID uuid.UUID
}

func (t *TimeUUIDCursor) Serialize() string {

	str := append(t.UUID.Bytes(), cursorSep...)

	str = t.Time.AppendFormat(str, timeFormat)

	return b64Encoding.EncodeToString(str)
}
