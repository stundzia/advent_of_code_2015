package day6

import (
	"aoc2015/helpers"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var coordsRegex = regexp.MustCompile(`[^\d]*(\d+),(\d+)[^\d]*(\d+),(\d+).*`)

func ParseCoords(fromX, fromY, toX, toY string) (fromXint, fromYint, toXint, toYint int) {
	fromXint64, _ := strconv.ParseInt(fromX, 10, 32)
	fromXint = int(fromXint64)
	fromYint64, _ := strconv.ParseInt(fromY, 10, 32)
	fromYint = int(fromYint64)
	toXint64, _ := strconv.ParseInt(toX, 10, 32)
	toXint = int(toXint64)
	toYint64, _ := strconv.ParseInt(toY, 10, 32)
	toYint = int(toYint64)
	return fromXint, fromYint, toXint, toYint
}

func HandleCommand(grid *Grid, command string) {
	match := coordsRegex.FindAllStringSubmatch(command, -1)
	if len(match) > 0 && len(match[0]) >= 5 {
		fromX := match[0][1]
		fromY := match[0][2]
		toX := match[0][3]
		toY := match[0][4]
		fromXint, fromYint, toXint, toYint := ParseCoords(fromX, fromY, toX, toY)
		switch true {
		case strings.HasPrefix(command, "turn off"):
			grid.TurnOff(fromXint, toXint, fromYint, toYint)
		case strings.HasPrefix(command, "toggle"):
			grid.Toggle(fromXint, toXint, fromYint, toYint)
		case strings.HasPrefix(command, "turn on"):
			grid.TurnOn(fromXint, toXint, fromYint, toYint)
		}
	}
}

func HandleBrightnessCommand(grid *Grid, command string) {
	match := coordsRegex.FindAllStringSubmatch(command, -1)
	if len(match) > 0 && len(match[0]) >= 5 {
		fromX := match[0][1]
		fromY := match[0][2]
		toX := match[0][3]
		toY := match[0][4]
		fromXint, fromYint, toXint, toYint := ParseCoords(fromX, fromY, toX, toY)
		switch true {
		case strings.HasPrefix(command, "turn off"):
			grid.AdjustBrightness(fromXint, toXint, fromYint, toYint, -1)
		case strings.HasPrefix(command, "toggle"):
			grid.AdjustBrightness(fromXint, toXint, fromYint, toYint, 2)
		case strings.HasPrefix(command, "turn on"):
			grid.AdjustBrightness(fromXint, toXint, fromYint, toYint, 1)
		}
	}
}

func DoSilver() {
	commands := helpers.LoadInputAsStringSlice(6, "\n")
	grid := &Grid{}
	grid.InitiateLightGrid(999, 999)
	for _, c := range commands {
		HandleCommand(grid, c)
	}
	fmt.Println("Part 1 Lights on, off, brightness: ")
	fmt.Println(grid.GetLightStatus())
}

func DoGold() {
	commands := helpers.LoadInputAsStringSlice(6, "\n")
	grid := &Grid{}
	grid.InitiateLightGrid(999, 999)
	for _, c := range commands {
		HandleBrightnessCommand(grid, c)
	}
	_, _, brightness := grid.GetLightStatus()
	fmt.Println("Part 2 Total brightness: ", brightness)
}