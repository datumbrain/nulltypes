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
func (nf NullFloat64) MarshalJSON() ([]byte, error) {
	if !nf.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(nf.Float64)
}

// UnmarshalJSON method is called by json.Unmarshal,
// whenever it is of type NullFloat64
func (nf *NullFloat64) UnmarshalJSON(b []byte) error {
	var f *float64
	if err := json.Unmarshal(b, &f); err != nil {
		return err
	}
	if f != nil {
		nf.Valid = true
		nf.Float64 = *f
	} else {
		nf.Valid = false
	}
	return nil
}

// Scan satisfies the sql.scanner floaterface
func (nf *NullFloat64) Scan(value interface{}) error {
	rt, ok := value.(float64)
	if ok {
		*nf = NullFloat64{rt, true}
	} else {
		*nf = NullFloat64{rt, false}
	}
	return nil
}

// Value satisfies the driver.Value interface
func (nf NullFloat64) Value() (driver.Value, error) {
	if nf.Valid {
		return nf.Float64, nil
	}
	return nil, nil
}
