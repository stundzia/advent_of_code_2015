package day6

import "testing"

func TestGrid(t *testing.T) {
	grid := &Grid{}
	grid.InitiateLightGrid(100, 100)
	on, off, _ := grid.GetLightStatus()
	if on != 0 {
		t.Errorf("Invalid on light count, got %d, should be 0", on)
	}
	if off != 10000 {
		t.Errorf("Invalid off light count, got %d, should be 10000", off)
	}

	grid.TurnOn(0, 10, 0, 10)
	on, off, _ = grid.GetLightStatus()
	if on != 121 {
		t.Errorf("Invalid on light count, got %d, should be 121", on)
	}
	if off != 9879 {
		t.Errorf("Invalid off light count, got %d, should be 9879", off)
	}

	grid.Toggle(5, 12, 5, 12)
	// 36 should turn off, 28 should turn on
	on, off, _ = grid.GetLightStatus()
	if on != 113 {
		t.Errorf("Invalid on light count, got %d, should be 113", on)
	}
	if off != 9887 {
		t.Errorf("Invalid off light count, got %d, should be 9887", off)
	}
}