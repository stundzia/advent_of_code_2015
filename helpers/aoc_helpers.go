package helpers

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
	"strings"
)

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