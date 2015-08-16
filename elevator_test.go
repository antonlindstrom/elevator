package main

import (
	"testing"
)

// TestDirectionString tests that the direction.String() returns the correct
// values
func TestDirectionString(t *testing.T) {
	tests := map[Direction]string{
		1:  "up",
		-1: "down",
	}

	for k, v := range tests {
		if k.String() != v {
			t.Errorf("Expected direction %d = %v, got %v\n", k, v, k.String())
		}
	}
}

// TestGoToFloor tests that the proper floor is set as goalFloor
func TestGoToFloor(t *testing.T) {
	e := &elevator{id: 1}
	e.GoToFloor(1)

	if e.goalFloor != 1 {
		t.Errorf("Expected goalFloor = 1, got %v\n", e.goalFloor)
	}

	e.GoToFloor(3)

	// update goal is required but done in e.Move()
	e.updateGoal()

	if e.goalFloor != 3 {
		t.Errorf("Expected goalFloor = 3, got %v\n", e.goalFloor)
	}
}

// TestMoveUp tests that the floor is being decremented properly
func TestMoveUp(t *testing.T) {
	e := &elevator{id: 1}
	e.GoToFloor(1)

	if e.goalFloor != 1 {
		t.Errorf("Expected goalFloor = 1, got %v\n", e.goalFloor)
	}

	e.Move()

	// Should move upwards
	if !e.AtFloor(1) {
		t.Errorf("Expected AtFloor(1) = true, got false\n")
	}
}

// TestMoveDown tests that the floor is being decremented properly
func TestMoveDown(t *testing.T) {
	e := &elevator{id: 1, floor: 1}
	e.GoToFloor(0)

	if e.goalFloor != 0 {
		t.Errorf("Expected goalFloor = 0, got %v\n", e.goalFloor)
	}

	e.Move()

	// Should move downwards one step
	if !e.AtFloor(0) {
		t.Errorf("Expected AtFloor(0) = true, got false\n")
	}
}

// TestDirection tests that the direction is set properly
func TestDirection(t *testing.T) {
	eUp := &elevator{id: 1, floor: 0, goalFloor: 10}

	if eUp.Direction() != 1 {
		t.Errorf("Expected Direction = up got %v\n", eUp.Direction())
	}

	eDown := &elevator{id: 1, floor: 10, goalFloor: 0}

	if eDown.Direction() != -1 {
		t.Errorf("Expected Direction = down got %v\n", eDown.Direction())
	}
}

// TestProximityTo tests that the proximity to a certain floor is reached
func TestProximityTo(t *testing.T) {
	tests := []struct {
		ElevatorFloor int
		UserFloor     int

		Expected int
	}{
		{
			ElevatorFloor: 0,
			UserFloor:     1,
			Expected:      1,
		},
		{
			ElevatorFloor: 0,
			UserFloor:     2,
			Expected:      2,
		},
		{
			ElevatorFloor: 2,
			UserFloor:     0,
			Expected:      2,
		},
		{
			ElevatorFloor: 5,
			UserFloor:     0,
			Expected:      5,
		},
	}

	for _, test := range tests {
		e := &elevator{id: 1, floor: test.ElevatorFloor}

		if e.ProximityTo(test.UserFloor) != test.Expected {
			t.Errorf("Expected ProximityTo = %v got %v\n", test.Expected, e.ProximityTo(test.UserFloor))
		}
	}
}

func TestRemoveVisited(t *testing.T) {
	e := &elevator{id: 1, floor: 0, goalFloor: 10}

	e.floors.Add(0)
	e.removeVisited()

	if e.floors.Len() != 0 {
		t.Errorf("Expected floors.Len() = 0 got %v\n", e.floors.Len())
	}

	e.floors.Add(0)
	e.floors.Add(1)
	e.removeVisited()

	if e.floors.Len() != 1 && e.floors.data[0] == 1 {
		t.Errorf("Expected floors = [1] got %+v\n", e.floors.data)
	}
}
