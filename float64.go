package nulltypes

import (
	"database/sql/driver"
	"encoding/json"
)

// NullFloat64 is a wrapper around float64
type NullFloat64 struct {
	Float64 float64
	Valid   bool
}

// Float64 method to get NullFloat64 object from float64
func Float64(Float64 float64) NullFloat64 {
	return NullFloat64{Float64, true}
}

// MarshalJSON method is called by json.Marshal,
// whenever it is of type NullFloat64
func (x NullFloat64) MarshalJSON() ([]byte, error) {
	if !x.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(x.Float64)
}

// UnmarshalJSON method is called by json.Unmarshal,
// whenever it is of type NullFloat64
func (this *NullFloat64) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &this.Float64)
	if err != nil {
		return err
	}
	this.Valid = true
	return nil
}

// satisfy the sql.scanner floaterface
func (x *NullFloat64) Scan(value interface{}) error {
	rt, ok := value.(float64)
	if ok {
		*x = NullFloat64{rt, true}
	} else {
		*x = NullFloat64{rt, false}
	}
	return nil
}

// satifies the driver.Value interface
func (x NullFloat64) Value() (driver.Value, error) {
	if x.Valid {
		return x.Float64, nil
	}
	return nil, nil
}
