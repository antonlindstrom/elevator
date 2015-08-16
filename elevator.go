package main

import (
	"fmt"
)

// Elevator is the interface which to use
type Elevator interface {
	// Return the ID of the elevator
	ID() int

	// String is the representation of the elevator status
	String() string

	// Move makes the elevator move in the direction the elevator is // called to
	Move()

	// GoToFloor sets a floor which the elevator should go to
	GoToFloor(int)

	// AtFloor returns true if the elevator is at the floor specified
	AtFloor(int) bool

	// Direction returns the direction the elevator is going
	Direction() Direction

	// ProximityTo is the proximity to a supplied floor
	ProximityTo(int) int
}

// Elevator is the state of an elevator at any given moment
type elevator struct {
	id        int
	floor     int
	goalFloor int
	floors    Floors
}

// NewElevator returns a new elevator instance
func NewElevator(id int) Elevator {
	return &elevator{
		id: id,
	}
}

// ID returns the ID of the elevator
func (e *elevator) ID() int {
	return e.id
}

// String is the string representation of an elevator
func (e *elevator) String() string {
	return fmt.Sprintf("elevator id=%d, floor=%d, goal=%d, direction=%s", e.id, e.floor, e.goalFloor, e.Direction())
}

// Move makes the elevator move one step at the direction it's going
func (e *elevator) Move() {
	switch e.Direction() {
	case Up:
		e.floor++
	case Down:
		e.floor--
	}

	e.removeVisited()
	e.updateGoal()
}

// GoToFloor is the interface to order an elevator to go to a specific floor
func (e *elevator) GoToFloor(floor int) {
	if e.floors.Len() == 0 {
		e.goalFloor = floor
	}
	e.floors.Add(floor)
}

// AtFloor checks if the elevator is at a specific floor
func (e *elevator) AtFloor(floor int) bool {
	return e.floor == floor
}

// Direction is the direction the elevator is going
func (e *elevator) Direction() Direction {
	if e.floor > e.goalFloor {
		return Down
	} else if e.floor < e.goalFloor {
		return Up
	}
	return 0
}

// ProximityTo returns the proximity to the floor supplied
func (e *elevator) ProximityTo(floor int) int {
	if e.floor > floor {
		return e.floor - floor
	}
	return floor - e.floor
}

// updateGoal updates the goal to the highest/lowest
func (e *elevator) updateGoal() {
	if e.floors.Len() <= 0 {
		return
	}

	if e.Direction() == Up {
		e.goalFloor = e.floors.PeekHighest()
	} else {
		e.goalFloor = e.floors.PeekLowest()
	}
}

// removeVistied removes visited floors
func (e *elevator) removeVisited() {
	if e.floors.Len() <= 0 {
		return
	}

	switch e.floor {
	case e.floors.PeekLowest():
		e.floors.Lowest()
	case e.floors.PeekHighest():
		e.floors.Highest()
	}
}
