package day4

import (
	"fmt"
	"testing"
)

func TestFindValid(t *testing.T) {
	hashBase := "abcdef"
	valid := findValid(hashBase, 5)
	if valid != fmt.Sprintf("%s%d", hashBase, 609043) {
		t.Errorf("Incorrect valid hash, expected abcdef609043, got %s", valid)
	}

	hashBase = "pqrstuv"
	valid = findValid(hashBase, 5)
	if valid != fmt.Sprintf("%s%d", hashBase, 1048970) {
		t.Errorf("Incorrect valid hash, expected pqrstuv1048970, got %s", valid)
	}
}
