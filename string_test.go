package nulltypes

import (
	"encoding/json"
	"testing"
)

type TestStruct struct {
	Name NullString
}

var tStr = `{"Name":"John Doe"}`
var tObj = TestStruct{
	Name: NullString{
		String: "John Doe",
		Valid:  true,
	},
}

func TestMarshalJSON(t *testing.T) {
	var ts1 TestStruct
	b := []byte(tStr)
	json.Unmarshal(b, &ts1)
	t.Logf("%+v", ts1)

	if ts1 != tObj {
		t.Errorf("Unmarshalling string didn't work as expected, %+v != %+v, somehow!", ts1, tObj)
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
