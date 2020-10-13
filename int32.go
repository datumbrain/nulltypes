package nulltypes

import (
	"database/sql/driver"
	"encoding/json"
)

// NullInt32 is a wrapper around int32
type NullInt32 struct {
	Int32 int32
	Valid bool
}

// Int32 method to get NullInt32 object from int32
func Int32(Int32 int32) NullInt32 {
	return NullInt32{Int32, true}
}

// MarshalJSON method is called by json.Marshal,
// whenever it is of type NullInt32
func (x *NullInt32) MarshalJSON() ([]byte, error) {
	if !x.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(x.Int32)
}

// UnmarshalJSON method is called by json.Unmarshal,
// whenever it is of type NullInt32
func (this NullInt32) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &this.Int32)
	if err != nil {
		return err
	}
	this.Valid = true
	return nil
}

// satisfy the sql.scanner interface
func (x *NullInt32) Scan(value interface{}) error {
	rt, ok := value.(int32)
	if ok {
		*x = NullInt32{rt, true}
	} else {
		*x = NullInt32{rt, false}
	}
	return nil
}

// satifies the driver.Value interface
func (x NullInt32) Value() (driver.Value, error) {
	if x.Valid {
		return x.Int32, nil
	}
	return nil, nil
}
