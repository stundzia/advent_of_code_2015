package day11

import (
	"strings"
)

var invalidCharacters = []string{"i", "o", "l"}

func ConditionOne(password string) bool {
	incrementCount := 0
	for i := 1; i < len(password); i++ {
		if password[i] - password[i - 1] == 1 {
			incrementCount++
			if incrementCount >= 2 {
				return true
			}
		} else {
			incrementCount = 0
		}
	}
	return false
}

func ConditionTwo(password string) bool {
	for _, c := range invalidCharacters {
		if strings.Contains(password, c) {
			return false
		}
	}
	return true
}

func ConditionThree(password string) bool {
	var pairOne []uint8
	for i := 1; i < len(password); i++ {
		if password[i-1] == password[i] {
			pairOne = []uint8{password[i-1], password[i]}
			if i < len(password) - 1 {
				for ; i < len(password); i++ {
					if password[i-1] == password[i] && password[i] != pairOne[0] {
						return true
					}
				}
			}
		}
	}
	return false
}

func IsValid(password string) bool {
	return ConditionOne(password) && ConditionTwo(password) && ConditionThree(password)
}

func IncrementRune(letter uint8) (res uint8, isLast bool) {
	if letter == 122 {
		return 97, true
	} else {
		return letter + 1, false
	}
}

func IncrementPass(password string) (res string) {
	passRunes := []uint8(password)
	var isLast bool
	for i := len(password) - 1; i > 0; i-- {
		passRunes[i], isLast = IncrementRune(password[i])
		if !isLast {
			return string(passRunes)
		}
	}
	return string(passRunes)
}