package day7

import "strconv"

func stringToUint16(str string) uint16 {
	i, _ := strconv.Atoi(str)
	return uint16(i)
}

func isInt(str string) bool {
	if _, err := strconv.Atoi(str); err != nil {
		return false
	}
	return true
}