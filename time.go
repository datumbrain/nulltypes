package nulltypes

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

// TruncateOff the degree of precision to REMOVE
// Default time.Microsecond
var TruncateOff = time.Microsecond

// DatabaseLocation is local the timezone
// the database is set to default UTC
var DatabaseLocation, _ = time.LoadLocation("UTC")

// NullTime is a wrapper around time.Time
type NullTime struct {
	Time  time.Time
	Valid bool `default:"false"`
}

// Time method to get NullTime object from time.Time
func Time(Time time.Time) NullTime {
	return NullTime{Time, true}
}

// MarshalJSON method is called by json.Marshal,
// whenever it is of type NullTime
func (nt NullTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(nt.Time)
}

// UnmarshalJSON method is called by json.Unmarshal,
// whenever it is of type NullTime
func (nt *NullTime) UnmarshalJSON(b []byte) error {
	var t *time.Time
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}
	if t != nil {
		nt.Valid = true
		nt.Time = *t
	} else {
		nt.Valid = false
	}
	return nil
}

// Scan satisfies the sql.scanner interface
func (nt *NullTime) Scan(value interface{}) error {
	rt, ok := value.(time.Time)
	if ok {
		*nt = NullTime{format(rt), true}
	} else {
		*nt = NullTime{time.Time{}, false}
	}
	return nil
}

// Value satisfies the driver.Value interface
func (nt NullTime) Value() (driver.Value, error) {
	if nt.Valid {
		return nt.Time, nil
	}

	return nil, nil
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
