package workout

import "github.com/gofrs/uuid"

// Associates a numeral and/or text with a unit ID (at least one must be specified)
// This object is attached to other objects to return unit data
type UnitData struct {
	Numeral *int      `json:"numeral"`
	Text    *string   `json:"text"`
	UnitID  uuid.UUID `json:"unitID"`
}
