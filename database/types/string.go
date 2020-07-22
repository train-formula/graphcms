package types

import (
	"database/sql"
	"database/sql/driver"
	"io"
)

func ReadNullString(str *string) NullString {

	var returnVal NullString

	if str == nil {
		returnVal = NullString{
			delegate: sql.NullString{
				String: "",
				Valid:  false,
			},
		}
	} else {
		returnVal = NullString{
			delegate: sql.NullString{
				String: *str,
				Valid:  true,
			},
		}
	}

	returnVal.assignFromDelegate()
	return returnVal
}

type NullString struct {
	String   string
	Valid    bool // Valid is true if String is not NULL
	delegate sql.NullString
}

func (n *NullString) assignFromDelegate() {
	n.String = n.delegate.String
	n.Valid = n.delegate.Valid
}

// Scan implements the Scanner interface.
func (n *NullString) Scan(value interface{}) error {
	err := n.delegate.Scan(value)
	if err != nil {
		return err
	}

	n.assignFromDelegate()
	return nil

}

// Value implements the driver Valuer interface.
func (n NullString) Value() (driver.Value, error) {
	return n.delegate.Value()
}

// For JSON marshalling
func (n NullString) MarshalJSON() ([]byte, error) {

	if !n.delegate.Valid {
		return nullBytes, nil
	}
	return []byte(n.delegate.String), nil
}

func (y *NullString) UnmarshalGQL(v interface{}) error {

	if v == nil {
		y.delegate = sql.NullString{
			String: "",
			Valid:  false,
		}
	}

	if st, okStr := v.(string); okStr {
		y.delegate = sql.NullString{
			String: st,
			Valid:  true,
		}
	} else if b, okB := v.([]byte); okB {
		y.delegate = sql.NullString{
			String: string(b),
			Valid:  true,
		}
	}

	y.assignFromDelegate()

	return nil
}

func (n NullString) MarshalGQL(w io.Writer) {
	b, _ := n.MarshalJSON()
	w.Write(b)
}
