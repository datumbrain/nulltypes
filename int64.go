package nulltypes

import (
	"database/sql/driver"
	"encoding/json"
)

// NullInt64 is a wrapper around int64
type NullInt64 struct {
	Int64 int64
	Valid bool
}

// Int64 method to get NullInt64 object from int64
func Int64(Int64 int64) NullInt64 {
	return NullInt64{Int64, true}
}

// MarshalJSON method is called by json.Marshal,
// whenever it is of type NullInt64
func (ni NullInt64) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(ni.Int64)
}

// UnmarshalJSON method is called by json.Unmarshal,
// whenever it is of type NullInt64
func (ni *NullInt64) UnmarshalJSON(b []byte) error {
	var i *int64
	if err := json.Unmarshal(b, &i); err != nil {
		return err
	}
	if i != nil {
		ni.Valid = true
		ni.Int64 = *i
	} else {
		ni.Valid = false
	}
	return nil
}

// Scan satisfies the sql.scanner interface
func (ni *NullInt64) Scan(value interface{}) error {
	rt, ok := value.(int64)
	if ok {
		*ni = NullInt64{rt, true}
	} else {
		*ni = NullInt64{rt, false}
	}
	return nil
}

// Value satifies the driver.Value interface
func (ni NullInt64) Value() (driver.Value, error) {
	if ni.Valid {
		return ni.Int64, nil
	}
	return nil, nil
}
