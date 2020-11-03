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

// String method to get a pointer of NullString object from string
func String(String string) *NullString {
	return &NullString{String, true}
}

// Set method to set the value
func (this *NullString) Set(value string) {
	this.String = value
	this.Valid = true
}

// MarshalJSON method is called by json.Marshal,
// whenever it is of type NullString
func (ns NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(ns.String)
}

// UnmarshalJSON method is called by json.Unmarshal,
// whenever it is of type NullString
func (ns *NullString) UnmarshalJSON(b []byte) error {
	var s *string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	if s != nil {
		ns.Valid = true
		ns.String = *s
	} else {
		ns.Valid = false
	}
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

// Value satisfies the driver.Value interface
func (ns NullString) Value() (driver.Value, error) {
	if ns.Valid {
		return ns.String, nil
	}
	return nil, nil
}
