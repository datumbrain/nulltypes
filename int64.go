package nulltypes

import (
	"database/sql/driver"
	"encoding/json"
)

//NullInt64 is a wrapper around int64
type NullInt64 struct {
	Int64 int64
	Valid bool
}

// Int64 method to get NullInt64 object from int64
func Int64(Int64 int64) NullInt64 {
	return NullInt64{Int64, true}
}

//MarshalJSON method is called by json.Marshal,
//whenever it is of type NullInt64
func (x *NullInt64) MarshalJSON() ([]byte, error) {
	if !x.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(x.Int64)
}

//UnmarshalJSON method is called by json.Unmarshal,
//whenever it is of type NullInt64
func (this NullInt64) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &this.Int64)
	if err != nil {
		return err
	}
	this.Valid = true
	return nil
}

// satisfy the sql.scanner interface
func (x *NullInt64) Scan(value interface{}) error {
	rt, ok := value.(int64)
	if ok {
		*x = NullInt64{rt, true}
	} else {
		*x = NullInt64{rt, false}
	}
	return nil
}

// satifies the driver.Value interface
func (x NullInt64) Value() (driver.Value, error) {
	if x.Valid {
		return x.Int64, nil
	}
	return nil, nil
}
