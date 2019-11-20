package day3


type Santa struct {
	CurrentlyAt *House
	HousesVisited *HousePool
	VisitCount int
}

func (santa *Santa) Move(direction rune) {
	coords := make([]int, 2, 2)
	coords = append([]int(nil), santa.CurrentlyAt.Coords...)
	santa.CurrentlyAt.HasSanta = false
	santa.VisitCount += 1
	switch direction {
	case '^':
		coords[1] += 1
		if santa.CurrentlyAt.North == nil {
			if targetHouse := santa.HousesVisited.getHouseByCoords(coords); targetHouse != nil {
				santa.CurrentlyAt.North = targetHouse
			} else {
				santa.CurrentlyAt.North = &House{
					South: santa.CurrentlyAt,
					HasSanta: true,
					VisitCount: 1,
					Coords: coords,
				}
				santa.HousesVisited.Count += 1
				santa.HousesVisited.Houses = append(santa.HousesVisited.Houses, santa.CurrentlyAt.North)
			}
		}
		santa.CurrentlyAt = santa.CurrentlyAt.North
		santa.CurrentlyAt.VisitCount += 1
		santa.CurrentlyAt.HasSanta = true
	case '<':
		coords[0] -= 1
		if santa.CurrentlyAt.West == nil {
			if targetHouse := santa.HousesVisited.getHouseByCoords(coords); targetHouse != nil {
				santa.CurrentlyAt.West = targetHouse
			} else {
				santa.CurrentlyAt.West = &House{
					East: santa.CurrentlyAt,
					HasSanta: true,
					VisitCount: 1,
					Coords: coords,
				}
				santa.HousesVisited.Count += 1
				santa.HousesVisited.Houses = append(santa.HousesVisited.Houses, santa.CurrentlyAt.West)
			}
		}
		santa.CurrentlyAt = santa.CurrentlyAt.West
		santa.CurrentlyAt.VisitCount += 1
		santa.CurrentlyAt.HasSanta = true
	case '>':
		coords[0] += 1
		if santa.CurrentlyAt.East == nil {
			if targetHouse := santa.HousesVisited.getHouseByCoords(coords); targetHouse != nil {
				santa.CurrentlyAt.East = targetHouse
			} else {
				santa.CurrentlyAt.East = &House{
					West: santa.CurrentlyAt,
					HasSanta:true,
					VisitCount:1,
					Coords: coords,
				}
				santa.HousesVisited.Count += 1
				santa.HousesVisited.Houses = append(santa.HousesVisited.Houses, santa.CurrentlyAt.East)
			}
		}
		santa.CurrentlyAt = santa.CurrentlyAt.East
		santa.CurrentlyAt.VisitCount += 1
		santa.CurrentlyAt.HasSanta = true
	case 'v':
		coords[1] -= 1
		if santa.CurrentlyAt.South == nil {
			if targetHouse := santa.HousesVisited.getHouseByCoords(coords); targetHouse != nil {
				santa.CurrentlyAt.South = targetHouse
			} else {
				santa.CurrentlyAt.South = &House{
					North: santa.CurrentlyAt,
					HasSanta:true,
					VisitCount:1,
					Coords: coords,
				}
				santa.HousesVisited.Count += 1
				santa.HousesVisited.Houses = append(santa.HousesVisited.Houses, santa.CurrentlyAt.South)
			}
		}
		santa.CurrentlyAt = santa.CurrentlyAt.South
		santa.CurrentlyAt.VisitCount += 1
		santa.CurrentlyAt.HasSanta = true
	}
}