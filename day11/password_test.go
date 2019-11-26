package day11

import (
	"fmt"
	"testing"
)

func TestConditionThree(t *testing.T) {
	password := "aaasdgfdbb"
	if !ConditionThree(password) {
		t.Errorf("`aaasdgfdbb` should be valid for condition 3")
	}

	password = "aab"
	if ConditionThree(password) {
		t.Errorf("`aab` should not be valid for condition 3")
	}

	password = "zzzzzzzz"
	if ConditionThree(password) {
		t.Errorf("`zzzzzzzz` should not be valid for condition 3")
	}

	password = "ccff"
	if !ConditionThree(password) {
		t.Errorf("`ccff` should be valid for condition 3")
	}
}

func TestConditionTwo(t *testing.T) {
	password := "aaasdgfdbb"
	if !ConditionTwo(password) {
		t.Errorf("`aaasdgfdbb` should be valid for condition 2")
	}

	password = "aab"
	if !ConditionTwo(password) {
		t.Errorf("`aab` should be valid for condition 2")
	}

	password = "zzzzzzzz"
	if !ConditionTwo(password) {
		t.Errorf("`zzzzzzzz` should be valid for condition 2")
	}

	password = "ccff"
	if !ConditionTwo(password) {
		t.Errorf("`ccff` should be valid for condition 2")
	}

	password = "ccoff"
	if ConditionTwo(password) {
		t.Errorf("`ccoff` should not be valid for condition 2")
	}
}

func TestConditionOne(t *testing.T) {
	password := "aaasdabcxgfdbb"
	if !ConditionOne(password) {
		t.Errorf("`aaasdabcxgfdbb` should be valid for condition 1")
	}

	password = "aab"
	if ConditionOne(password) {
		t.Errorf("`aab` should not be valid for condition 2")
	}

	password = "zzzzzzzz"
	if ConditionOne(password) {
		t.Errorf("`zzzzzzzz` should not be valid for condition 2")
	}

	password = "ccffefg"
	if !ConditionOne(password) {
		t.Errorf("`ccffefg` should be valid for condition 2")
	}

	// ?VALID?
	password = "ccoffzab"
	if ConditionOne(password) {
		t.Errorf("`ccoffzab` should not be valid for condition 1")
	}
	aa := "abbbz"
	fmt.Println(aa[0], aa[4])

	password = "xyzccoffzab"
	if !ConditionOne(password) {
		t.Errorf("`xyzccoffzab` should be valid for condition 1")
	}
}


func TestIncrementPass(t *testing.T) {
	startPass := "abcdtggwe"
	newPass := IncrementPass(startPass)
	if newPass != "abcdtggwf" {
		t.Errorf("expected `abcdtggwf`, got %s", newPass)
	}

	startPass = "zlsdggkzzz"
	newPass = IncrementPass(startPass)
	if newPass != "zlsdgglaaa" {
		t.Errorf("expected `zlsdgglaaa`, got %s", newPass)
	}
}