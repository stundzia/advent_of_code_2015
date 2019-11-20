package day9

import (
	"math/rand"
	"time"
)

type LocationPool struct {
	Locations []*Location
}

type Location struct {
	Name string
	DistanceMap map[*Location]int
	Visited bool
}

func (pool *LocationPool) GetLocationByName(name string) (location *Location) {
	for _, location := range pool.Locations {
		if location.Name == name {
			return location
		}
	}
	return nil
}

func (pool *LocationPool) GetOrCreateLocation(name string) (location *Location) {
	location = pool.GetLocationByName(name)
	if location == nil {
		location = &Location{
			Name:        name,
			DistanceMap: make(map[*Location]int),
			Visited:     false,
		}
		pool.Locations = append(pool.Locations, location)
	}
	return location
}

func (pool *LocationPool) AddOrUpdateLocations(loc1, loc2 string, distance int) {
	location1 := pool.GetOrCreateLocation(loc1)
	location2 := pool.GetOrCreateLocation(loc2)
	location1.UpdateDistance(location2, distance)
	location2.UpdateDistance(location1, distance)
}

func (loc *Location) UpdateDistance(target *Location, distance int) {
	loc.DistanceMap[target] = distance
}

func (pool *LocationPool) GetUnvisitedLocations() (res []*Location) {
	for _, loc := range pool.Locations {
		if loc.Visited == false {
			res = append(res, loc)
		}
	}
	return res
}


func (pool *LocationPool) GetRandomUnvisitedLocation() (loc *Location) {
	unvisitedLocations := pool.GetUnvisitedLocations()
	if len(unvisitedLocations) == 0 {
		return nil
	}
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	loc = unvisitedLocations[r.Intn(len(unvisitedLocations))]
	return loc
}

func (location *Location) GetUnvisitedDestinations() (destinations []*Location) {
	for loc, _ := range location.DistanceMap {
		if loc.Visited == false {
			destinations = append(destinations, loc)
		}
	}
	return destinations
}


func (pool *LocationPool) DoFullPath() (distance int, full bool) {
	full = false
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	traveled := 0
	current := pool.GetRandomUnvisitedLocation()
	current.Visited = true
	unvisited := current.GetUnvisitedDestinations()
	for ; len(unvisited) > 0; {
		target := unvisited[r.Intn(len(unvisited))]
		traveled += current.DistanceMap[target]
		current = target
		current.Visited = true
		unvisited = current.GetUnvisitedDestinations()
	}
	if len(pool.GetUnvisitedLocations()) == 0 {
		full = true
	}
	return traveled, full
}

func (location *Location) GetNextLocationOSPF() (loc *Location) {
	shortestDist := 99999999
	for k, d := range location.DistanceMap {
		if k.Visited == false {
			if loc == nil || d < shortestDist {
				loc = k
				shortestDist = d
			}
		}
	}
	return loc
}

func (location *Location) GetNextLocationOLPF() (loc *Location) {
	longestDist := 0
	for k, d := range location.DistanceMap {
		if k.Visited == false {
			if loc == nil || d > longestDist {
				loc = k
				longestDist = d
			}
		}
	}
	return loc
}

func (pool *LocationPool) DoFullPathOSPF(start *Location) (distance int, full bool) {
	full = false
	traveled := 0
	current := start
	current.Visited = true
	unvisited := current.GetUnvisitedDestinations()
	for ; len(unvisited) > 0; {
		target := current.GetNextLocationOSPF()
		traveled += current.DistanceMap[target]
		current = target
		current.Visited = true
		unvisited = current.GetUnvisitedDestinations()
	}
	if len(pool.GetUnvisitedLocations()) == 0 {
		full = true
	}
	return traveled, full
}

func (pool *LocationPool) DoFullPathOLPF(start *Location) (distance int, full bool) {
	full = false
	traveled := 0
	current := start
	current.Visited = true
	unvisited := current.GetUnvisitedDestinations()
	for ; len(unvisited) > 0; {
		target := current.GetNextLocationOLPF()
		traveled += current.DistanceMap[target]
		current = target
		current.Visited = true
		unvisited = current.GetUnvisitedDestinations()
	}
	if len(pool.GetUnvisitedLocations()) == 0 {
		full = true
	}
	return traveled, full
}

func (pool *LocationPool) ResetVisited() {
	for _, loc := range pool.Locations {
		loc.Visited = false
	}
}