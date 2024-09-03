package nulltypes

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
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
		// Return nil to ensure the field is omitted
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

// UnmarshalJSON method is called by json.Unmarshal,
// whenever it is of type NullString
func (ns *NullString) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		ns.Valid = false
		ns.String = ""
		return nil
	}
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	ns.Valid = true
	ns.String = s
	return nil
}

// Scan satisfies the sql.Scanner interface
func (ns *NullString) Scan(value interface{}) error {
	switch v := value.(type) {
	case string:
		*ns = NullString{v, true}
	case nil:
		*ns = NullString{"", false}
	default:
		return fmt.Errorf("unable to scan type %T into NullString", value)
	}
	return nil
}

// Value satisfies the driver.Valuer interface
func (ns NullString) Value() (driver.Value, error) {
	if ns.Valid {
		return ns.String, nil
	}
	return nil, nil
}
