package main

import (
	"testing"
)

func TestSetGoal(t *testing.T) {
	e := &Elevator{}
	e.SetGoal(1)

	if e.GoalFloor != 1 {
		t.Errorf("Expected GoalFloor = 1, got %v\n", e.GoalFloor)
	}

	if len(e.Floors) != 1 {
		t.Errorf("Expected Floors length = 1, got %v\n", len(e.Floors))
	}

	e.SetGoal(3)

	if e.GoalFloor != 3 {
		t.Errorf("Expected GoalFloor = 3, got %v\n", e.GoalFloor)
	}

	if len(e.Floors) != 2 {
		t.Errorf("Expected Floors length = 2, got %v\n", len(e.Floors))
	}
}

func TestMove(t *testing.T) {
	e := &Elevator{}
	e.SetGoal(1)

	if e.Floor != 0 {
		t.Errorf("Expected Floor = 0, got %v\n", e.Floor)
	}

	e.Move()

	if e.Floor != 1 {
		t.Errorf("Expected Floor = 1, got %v\n", e.Floor)
	}
}
