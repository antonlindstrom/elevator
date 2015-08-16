package main

import (
	"testing"
)

// nullElevator is made for testing
type nullElevator struct {
	id        int
	direction Direction
	proximity int
}

func (e *nullElevator) ID() int {
	return e.id
}

func (e *nullElevator) String() string {
	return string(e.id)
}

func (e *nullElevator) Move() {
}

func (e *nullElevator) GoToFloor(int) {
}

func (e *nullElevator) AtFloor(int) bool {
	return false
}

func (e *nullElevator) Direction() Direction {
	return e.direction
}

func (e *nullElevator) ProximityTo(int) int {
	return e.proximity
}

// TestNearestElevator tests that the nearest elevator is in fact returned
func TestNearestElevator(t *testing.T) {
	ne1 := &nullElevator{
		id:        0,
		proximity: 1,
		direction: -1,
	}
	ne2 := &nullElevator{
		id:        1,
		proximity: 2,
		direction: 1,
	}
	ne3 := &nullElevator{
		id:        2,
		proximity: 1,
		direction: 1,
	}

	es := &elevatorSystem{elevators: []Elevator{ne1, ne2, ne3}, maxFloors: 5}
	nearest, err := es.nearestElevator(0, 1)
	if err != nil {
		t.Fatalf("Got unexpected error: %s\n", err)
	}

	if nearest != ne3 {
		t.Errorf("Expected nearestElevator = 2, got %s\n", nearest)
	}
}

func TestSimulateCall(t *testing.T) {
	ne1 := &elevator{id: 0}
	ne2 := &elevator{id: 1}
	es := &elevatorSystem{elevators: []Elevator{ne1, ne2}, maxFloors: 5}

	elevator1 := es.Pickup(1, 1)
	elevator2 := es.Pickup(1, -1)

	if elevator1 == elevator2 {
		t.Errorf("Expected pickup with direction up to not equal pickup with direction down\n")
	}

	es.Step()

	if ne1.floor != 1 {
		t.Errorf("Expected ne1.floor == 1, got %d\n", ne1.floor)
	}

	if ne2.floor != 1 {
		t.Errorf("Expected ne2.floor == 1, got %d\n", ne2.floor)
	}
}
