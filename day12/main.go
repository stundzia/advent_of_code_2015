package day12

import (
	"aoc2015/helpers"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)



func cheatSilver(txt string) {
	reg := regexp.MustCompile(`(-?\d+)+`)
	matches := reg.FindAllStringSubmatch(txt, -1)
	sum := 0
	for _, m := range matches {
		r, _ := strconv.Atoi(m[0])
		sum += r
	}
	fmt.Println("Solution: ", sum)
}

func CountKV(k string, v interface{}) (sum int, redInside bool) {
	fmt.Println(k)
	sum = 0
	if ww, ok := v.([]interface{});ok {
		for _, n := range ww {
			add, _ := CountKV("", n)
			sum += add
		}
	}
	if ww, ok := v.(map[string]interface{});ok {
		for kk, vv := range ww {
			add, red := CountKV(kk, vv)
			if red {
				return 0, false
			}
			sum += add
		}
	}
	if ww, ok := v.(int);ok {
		return ww, false
	}
	if ww, ok := v.(float64);ok {
		return int(ww), false
	}
	if ww, ok := v.(string);ok {
		if strings.Contains(ww, "red") {
			return 0, true
		}
	}
	fmt.Println(fmt.Sprintf("%T", v))
	return sum, false
}

func DoSilver() {
	input := helpers.LoadInputTxtFromFile(12)
	cheatSilver(input)
}

func DoGold() {
	input := helpers.LoadInputTxtFromFile(12)
	var inputLoaded map[string]interface{}
	err := json.Unmarshal([]byte(input), &inputLoaded)
	if err != nil {
		fmt.Println(err)
	}
	sum := 0
	for k, v := range inputLoaded {
		add, red := CountKV(k, v)
		if !red {
			sum += add
		}
	}
	fmt.Println("Solution: ", sum)  //23400 is too low
}
