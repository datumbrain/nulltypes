package nulltypes

import (
	"database/sql/driver"
	"encoding/json"
)

//NullString is a wrapper around string
type NullString struct {
	String string
	Valid  bool
}

// String method to get NullString object from string
func String(String string) NullString {
	return NullString{String, true}
}

//MarshalJSON method is called by json.Marshal,
//whenever it is of type NullString
func (x *NullString) MarshalJSON() ([]byte, error) {
	if !x.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(x.String)
}

// satisfy the sql.scanner interface
func (x *NullString) Scan(value interface{}) error {
	rt, ok := value.(string)
	if ok {
		*x = NullString{rt, true}
	} else {
		*x = NullString{"", false}
	}
	return nil
}

// satifies the driver.Value interface
func (x NullString) Value() (driver.Value, error) {
	if x.Valid {
		return x.String, nil
	}
	return nil, nil
}
