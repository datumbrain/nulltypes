package nulltypes

import (
	"database/sql/driver"
	"encoding/json"
)

// NullBool is a wrapper around bool
type NullBool struct {
	Bool  bool
	Valid bool
}

// Bool method to get NullBool object from bool
func Bool(Bool bool) NullBool {
	return NullBool{Bool, true}
}

// MarshalJSON method is called by json.Marshal,
// whenever it is of type NullBool
func (x *NullBool) MarshalJSON() ([]byte, error) {
	if !x.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(x.Bool)
}

// UnmarshalJSON method is called by json.Unmarshal,
// whenever it is of type NullBool
func (this NullBool) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &this.Bool)
	if err != nil {
		return err
	}
	this.Valid = true
	return nil
}

// satisfy the sql.scanner interface
func (x *NullBool) Scan(value interface{}) error {
	rt, ok := value.(bool)
	if ok {
		*x = NullBool{rt, true}
	} else {
		*x = NullBool{rt, false}
	}
	return nil
}

// satifies the driver.Value interface
func (x NullBool) Value() (driver.Value, error) {
	if x.Valid {
		return x.Bool, nil
	}
	return nil, nil
}
