package day1

import (
	"aoc2015/helpers"
	"fmt"
)

func DoSilver()  {
	floor := 0
	actionMap := map[rune]int{
		'(': 1,
		')': -1,
	}
	txt := helpers.LoadInputTxtFromFile(1)
	actions := []rune(txt)
	for _, r := range actions {
		floor += actionMap[r]
	}
	fmt.Println(floor)
}

func DoGold()  {
	floor := 0
	actionMap := map[rune]int{
		'(': 1,
		')': -1,
	}
	txt := helpers.LoadInputTxtFromFile(1)
	actions := []rune(txt)
	fmt.Println("Number of actions: ", len(actions))
	for action_num, r := range actions {
		floor += actionMap[r]
		if floor == -1 {
			fmt.Println(action_num + 1)
			break
		}
	}
	fmt.Println(floor)
}