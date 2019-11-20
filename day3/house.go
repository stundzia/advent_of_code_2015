package day3


type House struct {
	East *House
	West *House
	North *House
	South *House
	Coords []int
	HasSanta bool
	VisitCount int
}

type HousePool struct {
	Houses []*House
	Count int
}

func (pool *HousePool) getHouseByCoords(coords []int) *House {
	for _, house := range pool.Houses {
		if house.Coords[0] == coords[0] && house.Coords[1] == coords[1] {
			return house
		}
	}
	return nil
}
