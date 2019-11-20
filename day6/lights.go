package day6

import (
	"aoc2015/helpers"
	"fmt"
	"strings"
)

type Grid struct {
	Lights []*Light
}

type Light struct {
	StatusOn bool
	Brightness int
	X int
	Y int
}

func (light *Light) TurnOn()  {
	light.StatusOn = true
}

func (light *Light) TurnOff()  {
	light.StatusOn = false
}

func (light *Light) Toggle()  {
	light.StatusOn = !light.StatusOn
}

func (light *Light) AdjustBrightness(delta int) {
	light.Brightness += delta
	if light.Brightness < 0 {
		light.Brightness = 0
	}
}

func (grid *Grid) InitiateLightGrid(xsize int, ysize int) {
	for x := 0; x < xsize; x++ {
		for y := 0; y < ysize; y++ {
			grid.Lights = append(
				grid.Lights,
				&Light{
					StatusOn: false,
					Brightness: 0,
					X:        x,
					Y:        y,
				},
			)
		}
	}
}

func (grid *Grid) GetLightStatus() (on int, off int, totalBrightness int) {
	for _, light := range grid.Lights {
		totalBrightness += light.Brightness
		if light.StatusOn {
			on++
		} else {
			off++
		}
	}
	return on, off, totalBrightness
}

func (grid *Grid) GetLights(fromX int, toX int, fromY int, toY int) (lights []*Light) {
	for _, light := range grid.Lights {
		if helpers.IntInBetween(light.X, fromX, toX) && helpers.IntInBetween(light.Y, fromY, toY) {
			lights = append(lights, light)
		}
	}
	return lights
}

func (grid *Grid) AdjustBrightness(fromX int, toX int, fromY int, toY int, delta int) {
	lights := grid.GetLights(fromX, toX, fromY, toY)
	for _, light := range lights {
		light.AdjustBrightness(delta)
	}
}

func (grid *Grid) TurnOff(fromX int, toX int, fromY int, toY int) {
	lights := grid.GetLights(fromX, toX, fromY, toY)
	for _, light := range lights {
		light.TurnOff()
	}
}

func (grid *Grid) TurnOn(fromX int, toX int, fromY int, toY int) {
	lights := grid.GetLights(fromX, toX, fromY, toY)
	for _, light := range lights {
		light.TurnOn()
	}
}

func (grid *Grid) Toggle(fromX int, toX int, fromY int, toY int) {
	lights := grid.GetLights(fromX, toX, fromY, toY)
	for _, light := range lights {
		light.Toggle()
	}
}

// This is a method used for debuging. Has a hardcoded 200x200 grid limit.
func (grid *Grid) PrintGrid() {
	lightArray := [200][200]string{}
	for _, light := range grid.Lights {
		if light.StatusOn {
			lightArray[light.Y][light.X] = "*"
		} else {
			lightArray[light.Y][light.X] = "#"
		}
	}
	for _, line := range lightArray {
		fmt.Println(strings.TrimSpace(strings.Join(line[:], "")))
	}
}