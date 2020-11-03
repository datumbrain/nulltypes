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

// Int32 method to get a pointer of NullInt32 object from int32
func Int32(Int32 int32) *NullInt32 {
	return &NullInt32{Int32, true}
}

// Set method to set the value
func (this *NullInt32) Set(value int32) {
	this.Int32 = value
	this.Valid = true
}

// MarshalJSON method is called by json.Marshal,
// whenever it is of type NullInt32
func (ni NullInt32) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(ni.Int32)
}

// UnmarshalJSON method is called by json.Unmarshal,
// whenever it is of type NullInt32
func (ni *NullInt32) UnmarshalJSON(b []byte) error {
	var i *int32
	if err := json.Unmarshal(b, &i); err != nil {
		return err
	}
	if i != nil {
		ni.Valid = true
		ni.Int32 = *i
	} else {
		ni.Valid = false
	}
	return nil
}

// Scan satisfies the sql.scanner interface
func (ni *NullInt32) Scan(value interface{}) error {
	rt, ok := value.(int32)
	if ok {
		*ni = NullInt32{rt, true}
	} else {
		*ni = NullInt32{rt, false}
	}
	return nil
}

// Value satisfies the driver.Value interface
func (ni NullInt32) Value() (driver.Value, error) {
	if ni.Valid {
		return ni.Int32, nil
	}
	return nil, nil
}
