package day7

import (
	"fmt"
	"regexp"
	"strconv"
)

type Wire struct {
	board *Board
	Signal uint16
	SignalString string
	SignalParsed bool
	SignalOperator string
	SignalOperands []string
	Id string
}

var operatorRegex = regexp.MustCompile(`[^A-Z]*([A-Z]{2,6})[^A-Z]*`)
var twoOperandRegex = regexp.MustCompile(`([a-z]{1,2}|\d+) [A-Z]{2,6} ([a-z]{1,2}|\d+)`)
var singleOperandRegex = regexp.MustCompile(`^[A-Z]{2,6} ([a-z]{1,2}|\d+)`)

type Board struct {
	Wires []*Wire
}

func (board *Board) GetWireById(id string) *Wire {
	for _, wire := range board.Wires {
		if wire.Id == id {
			return wire
		}
	}
	return nil
}

func (board *Board) AddWire(id string, signalString string) {
	if wireFound := board.GetWireById(id); wireFound == nil {
		board.Wires = append(board.Wires, &Wire{
			board: 	board,
			Signal:       0,
			SignalParsed: false,
			SignalString: signalString,
			Id:           id,
		})
	}
}

func (wire *Wire) ParseSignalString() {
	operMatch := operatorRegex.FindAllStringSubmatch(wire.SignalString, -1)
	if len(operMatch) > 0 && len(operMatch[0]) > 1 {
		wire.SignalOperator = operMatch[0][1]
	} else {
		wire.SignalOperator = ""
		wire.SignalOperands = []string{wire.SignalString}
	}
	operandMatch := twoOperandRegex.FindAllStringSubmatch(wire.SignalString, -1)
	if len(operandMatch) > 0 && len(operandMatch[0]) == 3 {
		wire.SignalOperands = []string{operandMatch[0][1], operandMatch[0][2]}
	} else {
		operandMatch = singleOperandRegex.FindAllStringSubmatch(wire.SignalString, -1)
		if len(operandMatch) > 0 && len(operandMatch[0]) == 2 {
			wire.SignalOperands = []string{operandMatch[0][1]}
		}
	}
}

func (wire *Wire) EvaluateSignal() {
	if wire.SignalParsed {
		return
	}
	for i, operand := range wire.SignalOperands {
		if !isInt(operand) {
			ww := wire.board.GetWireById(operand)
			if ww != nil {
				ww.EvaluateSignal()
				wire.SignalOperands[i] = strconv.Itoa(int(ww.Signal))
			} else {
				panic(fmt.Errorf("this should not be, wire not found %s", operand))
			}
		}
	}
	if len(wire.SignalOperator) == 0 {
		wire.Signal = stringToUint16(wire.SignalOperands[0])
	} else {
		op1 := stringToUint16(wire.SignalOperands[0])
		var op2 uint16
		if len(wire.SignalOperands) == 2 {
			op2 = stringToUint16(wire.SignalOperands[1])
			switch wire.SignalOperator {
			case "AND":
				wire.Signal = op1 & op2
			case "LSHIFT":
				wire.Signal = op1 << op2
			case "RSHIFT":
				wire.Signal = op1 >> op2
			case "OR":
				wire.Signal = op1 | op2
			}
		} else {
			switch wire.SignalOperator {
			case "NOT":
				wire.Signal = ^op1
			}
		}
	}
	wire.SignalParsed = true
}