package nulltypes

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

// the degree of precision to REMOVE
// Default time.Microsecond
var TruncateOff = time.Microsecond

// Local the timezone the database is set to
// default UTC
var DatabaseLocation, _ = time.LoadLocation("UTC")

//NullTime is a wrapper around time.Time
type NullTime struct {
	Time  time.Time
	Valid bool `default:"false"`
}

// Time method to get NullTime object from time.Time
func Time(Time time.Time) NullTime {
	return NullTime{Time, true}
}

// satisfy the sql.scanner interface
func (t *NullTime) Scan(value interface{}) error {
	rt, ok := value.(time.Time)
	if ok {
		*t = NullTime{format(rt), true}
	} else {
		*t = NullTime{time.Time{}, false}
	}
	return nil
}

// satifies the driver.Value interface
func (t NullTime) Value() (driver.Value, error) {
	if t.Valid {
		return t.Time, nil
	} else {
		return nil, nil
	}
}

// Now wrapper around the time.Now() function
func Now() NullTime {
	return NullTime{format(time.Now()), true}
}

// Date wrapper around the time.Date() function
func Date(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) NullTime {
	return NullTime{format(time.Date(year, month, day, hour, min, sec, nsec, loc)), true}
}

// insure the correct format
func format(t time.Time) time.Time {
	return t.In(DatabaseLocation).Truncate(TruncateOff)
}

//MarshalJSON method is called by json.Marshal,
//whenever it is of type NullTime
func (x *NullTime) MarshalJSON() ([]byte, error) {
	if !x.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(x.Time)
}

//UnmarshalJSON method is called by json.Unmarshal,
//whenever it is of type NullTime
func (this NullTime) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &this.Time)
	if err != nil {
		return err
	}
	this.Time = format(this.Time)
	this.Valid = true
	return nil
}
