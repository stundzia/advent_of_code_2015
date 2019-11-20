package day3

import (
	"aoc2015/helpers"
	"fmt"
)



func DoSilver()  {
	moves := helpers.LoadInputTxtFromFile(3)
	housePool := &HousePool{
		Houses: []*House{},
		Count:  1,
	}
	startingHouse := &House{
		East:       nil,
		West:       nil,
		North:      nil,
		South:      nil,
		HasSanta:   true,
		VisitCount: 1,
		Coords: []int{0, 0},
	}
	housePool.Houses = append(housePool.Houses, startingHouse)
	santa := Santa{
		CurrentlyAt: startingHouse,
		HousesVisited: housePool,
		VisitCount: 1,
	}
	for _, l := range moves {
		santa.Move(l)
	}
	fmt.Println(santa.HousesVisited.Count)  // Should match with below.
	fmt.Println(len(santa.HousesVisited.Houses))  // Should match with above.
}

func DoGold()  {
	moves := helpers.LoadInputTxtFromFile(3)
	housePool := &HousePool{
		Houses: []*House{},
		Count:  1,
	}
	startingHouse := &House{
		East:       nil,
		West:       nil,
		North:      nil,
		South:      nil,
		HasSanta:   true,
		VisitCount: 1,
		Coords: []int{0, 0},
	}
	housePool.Houses = append(housePool.Houses, startingHouse)
	santa := Santa{
		CurrentlyAt: startingHouse,
		HousesVisited: housePool,
		VisitCount: 1,
	}
	// Create robo-santa that shares the house pool and starts at the same house.
	roboSanta := Santa{
		CurrentlyAt:   startingHouse,
		HousesVisited: housePool,
		VisitCount:    1,
	}
	for i, l := range moves {
		if i % 2 == 0 {
			santa.Move(l)
		} else {
			roboSanta.Move(l)
		}
	}
	fmt.Println(santa.HousesVisited.Count)  // Should match with below.
	fmt.Println(len(santa.HousesVisited.Houses))  // Should match with above.
}
