package nulltypes

import (
	"encoding/json"
	"testing"
	"time"
)

type TestStruct struct {
	Name          NullString
	IsMarried     NullBool
	Height        NullFloat64
	Age           NullInt32
	Income        NullInt64
	BornAt        NullTime
	NameNull      NullString
	IsMarriedNull NullBool
	HeightNull    NullFloat64
	AgeNull       NullInt32
	IncomeNull    NullInt64
	BornAtNull    NullTime
}

var tz, _ = time.LoadLocation("UTC")
var timeToday time.Time = time.Date(2020, 10, 16, 01, 04, 30, 0, tz)
var tStr = `{"Name":"John Doe","IsMarried":true,"Height":170.4,"Age":40,"Income":4000000,"BornAt":"2020-10-16T01:04:30Z","NameNull":null,"IsMarriedNull":null,"HeightNull":null,"AgeNull":null,"IncomeNull":null,"BornAtNull":null}`
var tObj = TestStruct{
	Name:      NullString{String: "John Doe", Valid: true},
	IsMarried: NullBool{Bool: true, Valid: true},
	Height:    NullFloat64{Float64: 170.4, Valid: true},
	Age:       NullInt32{Int32: 40, Valid: true},
	Income:    NullInt64{Int64: 4000000, Valid: true},
	BornAt:    NullTime{Time: timeToday, Valid: true},
}

func TestMarshalJSON(t *testing.T) {
	var ts1 TestStruct
	b := []byte(tStr)
	err := json.Unmarshal(b, &ts1)
	if err != nil {
		t.Errorf("Something went bad while unmarshaling %+v obj...", tObj)
	}

	t.Logf("%+v", ts1)

	if ts1 != tObj {
		t.Errorf("Unmarshalling string didn't work as expected, \n%+v\n!=\n%+v\n, somehow!", ts1, tObj)
	}
}

func TestUnmarshalJSON(t *testing.T) {
	b, err := json.Marshal(tObj)
	if err != nil {
		t.Errorf("Something went bad while marshaling %+v obj...", tObj)
	}

	if string(b) != tStr {
		t.Errorf("Marshalling string didn't work as expected, %v != %v, somehow!", string(b), tStr)
	}
}
