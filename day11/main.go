package day11

import "fmt"

func DoSilver() {
	password := "cqjxjnds"
	for ;; {
		password = IncrementPass(password)
		if IsValid(password) {
			fmt.Println("Solution: ", password)
			break
		}
	}
}

func DoGold() {
	password := "cqjxxyzz"
	for ;; {
		password = IncrementPass(password)
		if IsValid(password) {
			fmt.Println("Solution: ", password)
			break
		}
	}
}