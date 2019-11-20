package helpers

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type solutionGetter func()

func EvaluateRuntime(function solutionGetter) {
	start := time.Now()
	function()
	fmt.Println("Solution evaluation took: ", time.Now().Sub(start))
}

func FailOnError(err error, msg string) {
	if err != nil {
		fmt.Println("Err msg: ", msg, " Err: ", err)
		log.Fatalf("Fatal Exception `%s` : %s", msg, err)
	}
}

func IntInBetween(i, min, max int) bool {
	if (i >= min) && (i <= max) {
		return true
	} else {
		return false
	}
}

func LoadInputTxtFromFile(dayNum int) string {
	fname := fmt.Sprintf("./aoc2015/day%d/input.txt", dayNum)
	absPath, _ := filepath.Abs(fname)
	dat, err := ioutil.ReadFile(absPath)
	if err != nil {
		FailOnError(err, "Could not load file")
	}
	return string(dat)
}


func LoadInputAsStringSlice(day int, sep string) []string {
	input := LoadInputTxtFromFile(day)
	return strings.Split(input, "\n")
}

func SliceStrToInt(slice []string) []int {
	res := []int{}
	for _, item := range slice {
		i, err := strconv.ParseInt(item, 10, 32)
		FailOnError(err, "Could not parse int")
		res = append(res, int(i))
	}
	return res
}

func StringContainsSubstringCount(str string, substring string) []int {
	substringStarts := []int{}
	MainLoop:
	for i, _ := range str {
		if str[i] == substring[0] {
			for t:=1;t < len(substring); t++ {
				if len(str) < i + t +1 || str[i+t] != substring[t] {
					continue MainLoop
				}
			}
			substringStarts = append(substringStarts, i)
		}
	}
	return substringStarts
}

func StringContainsCharCount(str string, letter rune) int {
	count := 0
	for _, char := range str {
		if letter == char {
			count++
		}
	}
	return count
}