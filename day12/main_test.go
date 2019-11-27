package day12

import (
	"encoding/json"
	"fmt"
	"testing"
)

func loadString(s string) (inputLoaded map[string]interface{}) {
	err := json.Unmarshal([]byte(s), &inputLoaded)
	if err != nil {
		fmt.Println(err)
	}
	return inputLoaded
}

func TestCountKV(t *testing.T) {
	tst := loadString(`{"a": 55, "b": 23}`)
	num, red := CountKV("", tst)
	if num != 78 || red != false {
		t.Errorf("expected 78, false but got %d %v", num, red)
	}

	tst = loadString(`{"a": 55, "b": 23, "c": {"a": 12, "t": "red"}}`)
	num, red = CountKV("", tst)
	if num != 78 || red != false {
		t.Errorf("expected 78, false but got %d %v", num, red)
	}

	tst = loadString(`{"a": "red", "b": 23, "c": {"a": 12, "t": "red"}}`)
	num, red = CountKV("", tst)
	if num != 78 || red != false {
		t.Errorf("expected 78, false but got %d %v", num, red)
	}
}