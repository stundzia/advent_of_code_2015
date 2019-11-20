package day3

import (
	"testing"
)

func TestMove1(t *testing.T) {
	santa := Santa{
		CurrentlyAt: &House{
			East:       nil,
			West:       nil,
			North:      nil,
			South:      nil,
			HasSanta:   true,
			VisitCount: 1,
			Coords: []int{0,0},
		},
		HousesVisited: &HousePool{
			Houses: []*House{},
			Count:  1,
		},
		VisitCount: 1,
	}
	santa.HousesVisited.Houses = append(santa.HousesVisited.Houses, santa.CurrentlyAt)
	moveStr := "^>v<"
	for _, l := range moveStr {
		santa.Move(l)
	}

	if santa.HousesVisited.Count != 4 {
		t.Errorf("Visited houses count incorrect, got: %d, want: 4.", santa.HousesVisited.Count)
	}
	if len(santa.HousesVisited.Houses) != 4 {
		t.Errorf("House pool size incorrect, got: %d, want: 4.", len(santa.HousesVisited.Houses))
	}
	if santa.VisitCount != 5 {
		t.Errorf("Total visit count incorrect, got: %d, want: 5.", santa.VisitCount)
	}
}

func TestMove2(t *testing.T) {
	santa := Santa{
		CurrentlyAt: &House{
			East:       nil,
			West:       nil,
			North:      nil,
			South:      nil,
			HasSanta:   true,
			VisitCount: 1,
			Coords: []int{0,0},
		},
		HousesVisited: &HousePool{
			Houses: []*House{},
			Count:  1,
		},
		VisitCount: 1,
	}
	santa.HousesVisited.Houses = append(santa.HousesVisited.Houses, santa.CurrentlyAt)
	moveStr := "^v^v^v^v^vvvv^^"
	for _, l := range moveStr {
		santa.Move(l)
	}

	if santa.HousesVisited.Count != 5 {
		t.Errorf("Visited houses count incorrect, got: %d, want: 5.", santa.HousesVisited.Count)
	}
	if len(santa.HousesVisited.Houses) != 5 {
		t.Errorf("House pool size incorrect, got: %d, want: 5.", len(santa.HousesVisited.Houses))
	}
	if santa.VisitCount != 16 {
		t.Errorf("Total visit count incorrect, got: %d, want: 16.", santa.VisitCount)
	}
}