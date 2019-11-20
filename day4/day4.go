package day4

import (
	"crypto/md5"
	"fmt"
)

func testValid(hashStr string, zeroCount int) bool  {
	valid := true
	for i := 0; i < zeroCount; i++ {
		if hashStr[i] == '0' {
			continue
		} else {
			valid = false
			break
		}
	}
	return valid
}

func findValid(hashStart string, zeroCount int) string {
	var hashStr string
	for i := 0; i < 500000000; i++ {
		hashStr = fmt.Sprintf("%s%d", hashStart, i)
		hashRes := md5.Sum([]byte(hashStr))
		if testValid(fmt.Sprintf("%x", hashRes), zeroCount) {
			return hashStr
		}
	}
	return ""
}


func DoSilver() {
	fmt.Println(findValid("bgvyzdsv", 5))
}

func DoGold() {
	fmt.Println(findValid("bgvyzdsv", 6))
}