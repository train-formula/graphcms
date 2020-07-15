package types

import (
	"database/sql"
	"database/sql/driver"
	"io"
	"strconv"
)

var nullBytes = []byte("null")

func ReadNullInt64(i *int64) NullInt64 {

	if i == nil {
		return NullInt64{
			delegate: sql.NullInt64{
				Int64: 0,
				Valid: false,
			},
		}
	}

	return NullInt64{
		delegate: sql.NullInt64{
			Int64: *i,
			Valid: true,
		},
	}
}

func ReadNullInt(i *int) NullInt64 {

	if i == nil {
		return NullInt64{
			delegate: sql.NullInt64{
				Int64: 0,
				Valid: false,
			},
		}
	}

	return NullInt64{
		delegate: sql.NullInt64{
			Int64: int64(*i),
			Valid: true,
		},
	}
}

// NullInt64 represents an int64 that may be null.
// NullInt64 implements the Scanner interface so
// it can be used as a scan destination, similar to NullString.
type NullInt64 struct {
	Int64    int64
	Valid    bool // Valid is true if Int64 is not NULL
	delegate sql.NullInt64
}

func (n *NullInt64) assignFromDelegate() {
	n.Int64 = n.delegate.Int64
	n.Valid = n.delegate.Valid
}

// Scan implements the Scanner interface.
func (n *NullInt64) Scan(value interface{}) error {
	err := n.delegate.Scan(value)
	if err != nil {
		return err
	}

	n.assignFromDelegate()
	return nil

}

// Value implements the driver Valuer interface.
func (n NullInt64) Value() (driver.Value, error) {
	return n.delegate.Value()
}

// For JSON marshalling
func (n NullInt64) MarshalJSON() ([]byte, error) {

	if !n.delegate.Valid {
		return nullBytes, nil
	}
	return []byte(strconv.FormatInt(n.delegate.Int64, 10)), nil
}

func (y *NullInt64) UnmarshalGQL(v interface{}) error {

	if v == nil {
		y.delegate = sql.NullInt64{
			Int64: 0,
			Valid: false,
		}
	}

	i, _ := v.(int64)

	y.delegate = sql.NullInt64{
		Int64: i,
		Valid: false,
	}

	y.assignFromDelegate()

	return nil
}

func (n NullInt64) MarshalGQL(w io.Writer) {
	b, _ := n.MarshalJSON()
	w.Write(b)
}
