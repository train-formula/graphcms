package workout

import (
	"github.com/gofrs/uuid"
	"github.com/train-formula/graphcms/database/types"
)

// Associates a numeral and/or text with a unit ID (at least one must be specified)
// This object is attached to other objects to return unit data
type UnitData struct {
	Numeral types.NullInt64 `json:"numeral"`
	Text    *string         `json:"text"`
	UnitID  uuid.UUID       `json:"unitID"`
}
