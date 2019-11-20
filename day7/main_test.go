package day7

import "testing"

func TestCreateWireFromLogicLineAndParseSignalString(t *testing.T) {
	logicLines := []string{
		"hz AND ik -> im",
		"z OR aa -> ab",
		"NOT bf -> bg",
		"la LSHIFT 15 -> le",
	}
	board := &Board{}
	createWireFromLogicLine(board, logicLines[0])
	wire := board.Wires[0]
	if wire.Id != "im" {
		t.Errorf("Wire id should be `im`, got %s", wire.Id)
	}
	wire.ParseSignalString()
	if wire.SignalOperands[0] != "hz" {
		t.Errorf("Wire first signal operand should be `hz`, got %s", wire.SignalOperands[0])
	}
	if wire.SignalOperands[1] != "ik" {
		t.Errorf("Wire second signal operand should be `ik`, got %s", wire.SignalOperands[1])
	}
	if wire.SignalOperator != "AND" {
		t.Errorf("Wire signal operator should be `AND`, got %s", wire.SignalOperator)
	}

	createWireFromLogicLine(board, logicLines[1])
	wire1 := board.Wires[1]
	if wire1.Id != "ab" {
		t.Errorf("Wire id should be `ab`, got %s", wire.Id)
	}
	wire1.ParseSignalString()
	if wire1.SignalOperands[0] != "z" {
		t.Errorf("Wire first signal operand should be `z`, got %s", wire1.SignalOperands[0])
	}
	if wire1.SignalOperands[1] != "aa" {
		t.Errorf("Wire second signal operand should be `aa`, got %s", wire1.SignalOperands[1])
	}
	if wire1.SignalOperator != "OR" {
		t.Errorf("Wire signal operator should be `OR`, got %s", wire1.SignalOperator)
	}
}