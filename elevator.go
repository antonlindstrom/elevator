package main

import (
	"fmt"
)

const (
	Up   = 1
	Down = -1
)

// Elevator is the state of an elevator at any given moment
type Elevator struct {
	ID        int
	Floor     int
	GoalFloor int
	Floors    []int
}

// String is the string representation of an elevator
func (e *Elevator) String() string {
	return fmt.Sprintf("elevator %d is currently on floor %d with goal floor %d", e.ID, e.Floor, e.GoalFloor)
}

// Move makes the elevator move one step at the direction it's going
func (e *Elevator) Move() {
	if e.GoalFloor == 0 {
		return
	}

	if e.direction() == Down {
		e.Floor--
	}

	if e.direction() == Up {
		e.Floor++
	}
}

// SetGoal is the interface to order an elevator to go to a specific floor
func (e *Elevator) SetGoal(goalFloor int) {
	if len(e.Floors) == 0 {
		e.GoalFloor = goalFloor
	}

	e.Floors = append(e.Floors, goalFloor)

	// checkGoal is done here because we want to see instant changes in
	// the GoalFloor when this is updated
	e.checkGoal()
}

// direction is the direction the elevator is going
func (e *Elevator) direction() int {
	if e.Floor > e.GoalFloor {
		return Down
	} else if e.Floor < e.GoalFloor {
		return Up
	}

	return 0
}

// atGoal checks if the elevator is at the current goal
func (e *Elevator) atGoal() bool {
	// Check goal to mark visited floors as visited
	e.checkGoal()

	if e.Floor == e.GoalFloor {
		return true
	}

	return false
}

// checkGoal checks all the goals (all floors that has been pushed and marks
// them as visited or sets the new goal.
func (e *Elevator) checkGoal() {
	var activeGoals []int
	max := -1
	min := -1

	for _, floor := range e.Floors {
		// The floor has been visited
		if e.Floor == floor {
			continue
		}

		if floor > max || max == -1 {
			max = floor
		}
		if floor < min || min == -1 {
			min = floor
		}

		activeGoals = append(activeGoals, floor)
	}

	if len(e.Floors) > 0 {
		if e.direction() == Up {
			e.GoalFloor = max
		} else if e.direction() == Down {
			e.GoalFloor = min
		}
	}

	e.Floors = activeGoals
}
