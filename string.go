package nulltypes

import (
	"database/sql/driver"
	"encoding/json"
)

// NullString is a wrapper around string
type NullString struct {
	String string
	Valid  bool
}

// String method to get NullString object from string
func String(s string) NullString {
	return NullString{
		String: s,
		Valid:  true,
	}
}

// MarshalJSON method is called by json.Marshal,
// whenever it is of type NullString
func (ns NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

// UnmarshalJSON method is called by json.Unmarshal,
// whenever it is of type NullString
func (ns *NullString) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ns.String)
	if err != nil {
		return err
	}
	ns.Valid = true
	return nil
}

// Scan satisfies the sql.scanner interface
func (ns *NullString) Scan(value interface{}) error {
	rt, ok := value.(string)
	if ok {
		*ns = NullString{rt, true}
	} else {
		*ns = NullString{"", false}
	}
	return nil
}

// Value satifies the driver.Value interface
func (ns NullString) Value() (driver.Value, error) {
	if ns.Valid {
		return ns.String, nil
	}
	return nil, nil
}
