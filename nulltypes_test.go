package nulltypes

import (
	"encoding/json"
	"testing"
	"time"
)

type TestStruct struct {
	Name          NullString `json:"Name,omitempty"`
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
var tStrWOmitEmpty = `{"Name":null,"IsMarried":true,"Height":170.4,"Age":40,"Income":4000000,"BornAt":"2020-10-16T01:04:30Z","NameNull":null,"IsMarriedNull":null,"HeightNull":null,"AgeNull":null,"IncomeNull":null,"BornAtNull":null}`

var tObj = TestStruct{
	Name:      NullString{String: "John Doe", Valid: true},
	IsMarried: NullBool{Bool: true, Valid: true},
	Height:    NullFloat64{Float64: 170.4, Valid: true},
	Age:       NullInt32{Int32: 40, Valid: true},
	Income:    NullInt64{Int64: 4000000, Valid: true},
	BornAt:    NullTime{Time: timeToday, Valid: true},
}

var tObjWithEmptyString = TestStruct{
	IsMarried: NullBool{Bool: true, Valid: true},
	Height:    NullFloat64{Float64: 170.4, Valid: true},
	Age:       NullInt32{Int32: 40, Valid: true},
	Income:    NullInt64{Int64: 4000000, Valid: true},
	BornAt:    NullTime{Time: timeToday, Valid: true},
}

func TestUnmarshalJSON(t *testing.T) {
	// Test with all fields present
	var ts1 TestStruct
	err := json.Unmarshal([]byte(tStr), &ts1)
	if err != nil {
		t.Errorf("Error unmarshaling JSON: %v", err)
	}
	if ts1 != tObj {
		t.Errorf("Unmarshaling didn't work as expected:\nGot: %+v\nWant: %+v", ts1, tObj)
	}

	// Test with Name field omitted (Name.Valid should be false)
	var ts2 TestStruct
	err = json.Unmarshal([]byte(tStrWOmitEmpty), &ts2)
	if err != nil {
		t.Errorf("Error unmarshaling JSON: %v", err)
	}
	if ts2 != tObjWithEmptyString {
		t.Errorf("Unmarshaling didn't work as expected:\nGot: %+v\nWant: %+v", ts2, tObjWithEmptyString)
	}
}

func TestMarshalJSON(t *testing.T) {
	// Test marshaling with all fields present
	b, err := json.Marshal(tObj)
	if err != nil {
		t.Errorf("Error marshaling object: %v", err)
	}
	if string(b) != tStr {
		t.Errorf("Marshaling didn't work as expected:\nGot: %v\nWant: %v", string(b), tStr)
	}

	// Test marshaling with Name field omitted (Name.Valid is false)
	b, err = json.Marshal(tObjWithEmptyString)
	if err != nil {
		t.Errorf("Error marshaling object: %v", err)
	}
	if string(b) != tStrWOmitEmpty {
		t.Errorf("Marshaling didn't work as expected:\nGot: %v\nWant: %v", string(b), tStrWOmitEmpty)
	}
}
