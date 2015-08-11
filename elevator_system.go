package main

import (
	"errors"
)

var (
	// ErrNoElevator will be returned when no elevator could be sent to a
	// pickup request.
	ErrNoElevator = errors.New("error: no elevator could be dispatched")
)

// ElevatorSystem is the collection of all elevators in the building
type ElevatorSystem struct {
	maxFloors int
	elevators []*Elevator
}

// New instantiates a new elevator system with a number of elevators specified
func New(numElevators, maxFloors int) *ElevatorSystem {
	var elevators []*Elevator
	for i := 0; i < numElevators; i++ {
		elevators = append(elevators, &Elevator{ID: i, Floor: 0})
	}

	return &ElevatorSystem{elevators: elevators, maxFloors: maxFloors}
}

// Status returns the status (ID, floor and goal floor) of all elevators in
// the system.
func (es *ElevatorSystem) Status() []*Elevator {
	return es.elevators
}

// Pickup is the interface to request an elevator to the specified floor
func (es *ElevatorSystem) Pickup(floor, direction int) (*Elevator, error) {
	elevator, err := es.nearestElevator(floor, direction)
	if err != nil {
		return nil, err
	}

	elevator.SetGoal(floor)
	return elevator, nil
}

// Step is the ticker for the elevator system to run
func (es *ElevatorSystem) Step() {
	for _, elevator := range es.elevators {
		if elevator.atGoal() {
			continue
		}

		elevator.Move()
	}
}

// nearestElevator gets the nearest elevator to a floor with the correct
// direction
func (es *ElevatorSystem) nearestElevator(floor, direction int) (*Elevator, error) {
	var nearest *Elevator
	score := es.maxFloors

	for _, e := range es.elevators {
		if e.direction() != 0 && e.direction() != direction {
			continue
		}

		nv := 0
		if e.Floor > floor {
			nv = e.Floor - floor
		} else {
			nv = floor - e.Floor
		}

		// Set the lowest score as the nearest elevator
		if nv < score {
			score = nv
			nearest = e
		}
	}

	if nearest == nil {
		return nil, ErrNoElevator
	}

	return nearest, nil
}
