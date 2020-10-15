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
func (nb NullBool) MarshalJSON() ([]byte, error) {
	if !nb.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(nb.Bool)
}

// UnmarshalJSON method is called by json.Unmarshal,
// whenever it is of type NullBool
func (nb *NullBool) UnmarshalJSON(b []byte) error {
	var boo *bool
	if err := json.Unmarshal(b, &boo); err != nil {
		return err
	}
	if boo != nil {
		nb.Valid = true
		nb.Bool = *boo
	} else {
		nb.Valid = false
	}
	return nil
}

// Scan satisfies the sql.scanner interface
func (nb *NullBool) Scan(value interface{}) error {
	rt, ok := value.(bool)
	if ok {
		*nb = NullBool{rt, true}
	} else {
		*nb = NullBool{rt, false}
	}
	return nil
}

// Value satifies the driver.Value interface
func (nb NullBool) Value() (driver.Value, error) {
	if nb.Valid {
		return nb.Bool, nil
	}
	return nil, nil
}
