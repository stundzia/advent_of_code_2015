package day10

import (
	"fmt"
	"testing"
)

func test2DsliceEqual(slice1 [][]int, slice2 [][]int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i, _ := range slice1 {
		if len(slice1[i]) != len(slice2[i]) {
			return false
		}
		for t, _ := range slice1[i] {
			if slice1[i][t] != slice2[i][t] {
				return false
			}
		}
	}
	return true
}

func TestSplitSame(t *testing.T) {
	val := []int{1}
	res := splitSame(val)
	expectedRes := [][]int{[]int{1}}
	if !test2DsliceEqual(res, expectedRes) {
		t.Errorf("Invalid result")
	}

	val = []int{1, 1, 2, 3, 3}
	res = splitSame(val)
	expectedRes = [][]int{{1, 1}, {2}, {3, 3}}
	if !test2DsliceEqual(res, expectedRes) {
		fmt.Println(res, " != ", expectedRes)
		t.Errorf("Invalid result")
	}
}

func TestDoIteration(t *testing.T) {
	input := []int{1,1,2,3,3}
	input = DoIteration(input)
	expect := []int{2,1,1,2,2,3}
	for i, v := range input {
		if v != expect[i] {
			t.Errorf("Incorrect val at index %d. Expected %d, got %d", i, expect[i], v)
		}
	}
	input = DoIteration(input)
	expect = []int{1,2,2,1,2,2,1,3}
	for i, v := range input {
		if v != expect[i] {
			t.Errorf("Incorrect val at index %d. Expected %d, got %d", i, expect[i], v)
		}
	}
	input = DoIteration(input)
	expect = []int{1,1,2,2,1,1,2,2,1,1,1,3}
	for i, v := range input {
		if v != expect[i] {
			t.Errorf("Incorrect val at index %d. Expected %d, got %d", i, expect[i], v)
		}
	}
}