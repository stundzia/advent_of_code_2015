package day7

import (
	"aoc2015/helpers"
	"fmt"
	"strings"
)

func createWireFromLogicLine(board *Board, line string)  {
	lineSlice := strings.Split(line," -> ")
	if len(lineSlice) != 2 {
		return
	}
	wireId := lineSlice[1]
	signalString := lineSlice[0]
	board.AddWire(wireId, signalString)
}

func DoSilver()  {
	lines := helpers.LoadInputAsStringSlice(7, "\n")
	board := &Board{}
	for _, line := range lines {
		createWireFromLogicLine(board, line)
	}
	for _, wire := range board.Wires {
		wire.ParseSignalString()
	}
	for _, wire := range board.Wires {
		wire.EvaluateSignal()
	}
	resWire := board.GetWireById("a")
	fmt.Println(resWire.Signal)
}

func DoGold()  {
	lines := helpers.LoadInputAsStringSlice(7, "\n")
	board := &Board{}
	for _, line := range lines {
		createWireFromLogicLine(board, line)
	}
	b := board.GetWireById("b")
	b.Signal = 16076
	b.SignalParsed = true
	for _, wire := range board.Wires {
		wire.ParseSignalString()
	}
	for _, wire := range board.Wires {
		wire.EvaluateSignal()
	}
	resWire := board.GetWireById("a")
	fmt.Println(resWire.Signal)
}