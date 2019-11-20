package day5

import (
	"aoc2015/helpers"
	aoc2018helpers "aoc2018/helpers"
	"fmt"
	"strings"
)

var vowels = "aeiou"

func conditionOne(str string) bool {
	vowelCount := 0
	for _, vowel := range vowels {
		vowelCount += aoc2018helpers.StringContainsCharCount(str, vowel)
		if vowelCount >= 3 {
			return true
		}
	}
	return false
}

func conditionTwo(str string) bool {
	for i := 0; i < len(str) - 1; i++ {
		if str[i] == str[i+1] {
			return true
		}
	}
	return false
}

func conditionThree(str string) bool {
	substrings := []string{
		"ab", "cd", "pq", "xy",
	}
	for _, substring := range substrings {
		if strings.Contains(str, substring) {
			return false
		}
	}
	return true
}

// isNice (TM) checks if a string is considered nice by the conditions
// provided in part 1 (funcs conditionOne, conditionTwo and conditionThree)
func isNice(str string) bool {
	if conditionOne(str) && conditionTwo(str) && conditionThree(str) {
		return true
	}
	return false
}

func conditionFour(str string) bool {
	for i := 0; i < len(str) - 1; i++ {
		counts := aoc2018helpers.StringContainsSubstringCount(str, fmt.Sprintf("%c%c", str[i], str[i+1]))
		fmt.Println(counts)
		if len(counts) == 2 {
			if counts[1] - counts[0] > 1 {
				return true
			}
		} else if len(counts) > 2 {
			return true
		}
	}

	return false
}

func conditionFive(str string) bool {
	for i := 0; i < len(str) - 2; i++ {
		if str[i] == str[i+2] {
			return true
		}
	}
	return false
}

// isNice2 (TM) checks if a string is considered nice by the conditions
// provided in part 2 (funcs conditionFour and conditionFive)
func isNice2(str string) bool {
	if conditionFour(str) && conditionFive(str) {
		return true
	}
	return false
}

func DoSilver() {
	input := helpers.LoadInputAsStringSlice(5, "\n")
	nice := 0
	naughty := 0
	for _, str := range input {
		if isNice(str) {
			nice++
		} else {
			naughty++
		}
	}
	fmt.Println("Nice: ", nice, " Naughty: ", naughty)
}

func DoGold() {
	input := helpers.LoadInputAsStringSlice(5, "\n")
	nice := 0
	naughty := 0
	for _, str := range input {
		if isNice2(str) {
			nice++
		} else {
			naughty++
		}
	}
	fmt.Println("Nice: ", nice, " Naughty: ", naughty)
}