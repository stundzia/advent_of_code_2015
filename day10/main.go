package day10

import "fmt"

func splitSame(val []int) (res [][]int) {
	if len(val) < 1 {
		return res
	}
	res = [][]int{{}}
	res[0] = append(res[0], val[0])
	currentResIndex := 0
	for i, v := range val {
		if i == 0 {
			continue
		}
		if v != val[i - 1] {
			res = append(res, []int{v})
			currentResIndex++
		} else {
			res[currentResIndex] = append(res[currentResIndex], v)
		}
	}
	return res
}

func SameIntsCount(input []int) []int {
	return []int{len(input), input[0]}
}

func IntSlice2Dto1D(input [][]int) (res []int) {
	for _, slice := range input {
		for _, num := range slice {
			res = append(res, num)
		}
	}
	return res
}

func DoIteration(input []int) (res []int) {
	splits := splitSame(input)
	newSplits := [][]int{{}}
	for i, split := range splits {
		if i == 0 {
			newSplits[0] = SameIntsCount(split)
		} else {
			newSplits = append(newSplits, SameIntsCount(split))
		}
	}
	return IntSlice2Dto1D(newSplits)
}


func DoSilver() {
	input := []int{3,1,1,3,3,2,2,1,1,3}
	for i := 0; i < 40; i++ {
		input = DoIteration(input)
	}
	fmt.Println("Silver Solution: ", len(input))
}

func DoGold() {
	input := []int{3,1,1,3,3,2,2,1,1,3}
	for i := 0; i < 50; i++ {
		input = DoIteration(input)
	}
	fmt.Println("Gold Solution: ", len(input))
}