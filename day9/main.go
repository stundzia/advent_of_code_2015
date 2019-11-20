package day9

import (
	"aoc2015/helpers"
	"fmt"
	"strconv"
	"strings"
)

func ParseLocationLine(line string) (loc1, loc2 string, distance int) {
	//Tambi to Faerun = 39
	//Tambi to Norrath = 113
	//Tambi to Straylight = 130
	//Tambi to Tristram = 35
	//Tambi to Arbre = 40
	strSplit := strings.Split(line, " to ")
	loc1 = strSplit[0]
	strSplitSplit := strings.Split(strSplit[1], " = ")
	loc2 = strSplitSplit[0]
	distanceStr := strSplitSplit[1]
	distance, _ = strconv.Atoi(distanceStr)
	return loc1, loc2, distance
}

func DoSilverBruteForce() {
	// Left this here just for fun.
	distances := helpers.LoadInputAsStringSlice(9, "\n")
	pool := &LocationPool{}
	for _, d := range distances {
		loc1, loc2, distance := ParseLocationLine(d)
		pool.AddOrUpdateLocations(loc1, loc2, distance)
	}
	shortestDistance := 5000000
	for i := 0; i < 299912; i++ {
		dst, full := pool.DoFullPath()
		if dst < shortestDistance && full {
			shortestDistance = dst
		}
		pool.ResetVisited()
	}
	fmt.Println("Solution: ", shortestDistance)
}

func DoSilver() {
	distances := helpers.LoadInputAsStringSlice(9, "\n")
	pool := &LocationPool{}
	for _, d := range distances {
		loc1, loc2, distance := ParseLocationLine(d)
		pool.AddOrUpdateLocations(loc1, loc2, distance)
	}
	shortestDistance := 5000000
	for _, loc := range pool.Locations {
		dst, full := pool.DoFullPathOSPF(loc)
		if dst < shortestDistance && full {
			shortestDistance = dst
		}
		pool.ResetVisited()
	}
	fmt.Println("Solution: ", shortestDistance)
}

func DoGoldBruteForce()  {
	distances := helpers.LoadInputAsStringSlice(9, "\n")
	pool := &LocationPool{}
	for _, d := range distances {
		loc1, loc2, distance := ParseLocationLine(d)
		pool.AddOrUpdateLocations(loc1, loc2, distance)
	}
	longestDistance := 0
	for i := 0; i < 2999999; i++ {
		dst, full := pool.DoFullPath()
		if dst > longestDistance && full {
			longestDistance = dst
		}
		pool.ResetVisited()
	}
	fmt.Println("Solution: ", longestDistance)
}

func DoGold()  {
	distances := helpers.LoadInputAsStringSlice(9, "\n")
	pool := &LocationPool{}
	for _, d := range distances {
		loc1, loc2, distance := ParseLocationLine(d)
		pool.AddOrUpdateLocations(loc1, loc2, distance)
	}
	longestDistance := 0
	for _, loc := range pool.Locations {
		dst, full := pool.DoFullPathOLPF(loc)
		if dst > longestDistance && full {
			longestDistance = dst
		}
		pool.ResetVisited()
	}
	fmt.Println("Solution: ", longestDistance)
}