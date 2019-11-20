package day6

import "testing"

func TestHandleCommand(t *testing.T) {
	grid := &Grid{}
	grid.InitiateLightGrid(100, 100)
	commands := []string{
		"turn on 10,10 through 20,20",
		"toggle 10,10 through 20,20",
		"toggle 15,15 through 25,25",
		"turn off 10,10 through 20,20",
	}
	on, off, _ := grid.GetLightStatus()
	if on != 0 {
		t.Errorf("Invalid on light count, got %d, should be 0", on)
	}
	if off != 10000 {
		t.Errorf("Invalid off light count, got %d, should be 10000", off)
	}

	HandleCommand(grid, commands[0])
	on, off, _ = grid.GetLightStatus()
	if on != 121 {
		t.Errorf("Invalid on light count, got %d, should be 121", on)
	}
	if off != 9879 {
		t.Errorf("Invalid off light count, got %d, should be 9879", off)
	}

	HandleCommand(grid, commands[1])
	on, off, _ = grid.GetLightStatus()
	if on != 0 {
		t.Errorf("Invalid on light count, got %d, should be 0", on)
	}
	if off != 10000 {
		t.Errorf("Invalid off light count, got %d, should be 10000", off)
	}

	HandleCommand(grid, commands[2])
	on, off, _ = grid.GetLightStatus()
	if on != 121 {
		t.Errorf("Invalid on light count, got %d, should be 121", on)
	}
	if off != 9879 {
		t.Errorf("Invalid off light count, got %d, should be 9879", off)
	}

	HandleCommand(grid, commands[3])
	on, off, _ = grid.GetLightStatus()
	if on != 85 {
		t.Errorf("Invalid on light count, got %d, should be 85", on)
	}
	if off != 9915 {
		t.Errorf("Invalid off light count, got %d, should be 9915", off)
	}
}

func TestHandleBrightnessCommand(t *testing.T) {
	grid := &Grid{}
	grid.InitiateLightGrid(100, 100)
	commands := []string{
		"turn on 10,10 through 20,20",
		"toggle 10,10 through 20,20",
		"toggle 15,15 through 25,25",
		"turn off 10,10 through 20,20",
	}
	_, _, brightness := grid.GetLightStatus()
	if brightness != 0 {
		t.Errorf("Invalid brightness level, got %d, should be 0", brightness)
	}

	HandleBrightnessCommand(grid, commands[0])
	_, _, brightness = grid.GetLightStatus()
	if brightness != 121 {
		t.Errorf("Invalid brightness level, got %d, should be 121", brightness)
	}

	HandleBrightnessCommand(grid, commands[1])
	_, _, brightness = grid.GetLightStatus()
	if brightness != 363 {
		t.Errorf("Invalid brightness level, got %d, should be 363", brightness)
	}
}