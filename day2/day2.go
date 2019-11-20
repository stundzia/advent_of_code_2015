package day2

import (
	"aoc2015/helpers"
	"fmt"
	"sort"
	"strings"
)

func calcPresentWrapSize(dims []int) int {
	bonus := dims[0] * dims[1]
	size := (dims[0] * dims[1] * 2) + (dims[1] * dims[2] * 2) + (dims[2] * dims[0] * 2)
	return size + bonus
}

func parsePresentDimensions(dimensionsStr string) []int  {
	dimsSlice := strings.Split(dimensionsStr, "x")
	dimsInt := helpers.SliceStrToInt(dimsSlice)
	sort.Ints(dimsInt)
	return dimsInt
}

func getPresentWrapSize(dimsStr string) int {
	dimsInt := parsePresentDimensions(dimsStr)
	return calcPresentWrapSize(dimsInt)
}

func getPresentRibbonLen(dimsStr string) int {
	dimsInt := parsePresentDimensions(dimsStr)
	ribbonTie := dimsInt[0] * 2 + dimsInt[1] * 2
	ribbonBow := dimsInt[0] * dimsInt[1] * dimsInt[2]
	return ribbonTie + ribbonBow
}

func DoSilver()  {
	sum := 0
	presents := helpers.LoadInputAsStringSlice(2, "\n")
	for _, present := range presents {
		sum += getPresentWrapSize(present)
	}
	fmt.Println(sum)
}

func DoGold()  {
	sum := 0
	presents := helpers.LoadInputAsStringSlice(2, "\n")
	for _, present := range presents {
		sum += getPresentRibbonLen(present)
	}
	fmt.Println(sum)
}