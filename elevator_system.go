package main

import (
	"errors"
)

var (
	// ErrNoElevator will be returned when no elevator could be sent to a
	// pickup request.
	ErrNoElevator = errors.New("error: no elevator could be dispatched")
)

type ElevatorSystem interface {
	Status() []Elevator
	Pickup(int, int) Elevator
	Step()
}

// ElevatorSystem is the collection of all elevators in the building
type elevatorSystem struct {
	maxFloors int
	elevators []Elevator
}

// New instantiates a new elevator system with a number of elevators specified
func New(numElevators, maxFloors int) ElevatorSystem {
	var elevators []Elevator
	for i := 0; i < numElevators; i++ {
		elevators = append(elevators, NewElevator(i))
	}

	return &elevatorSystem{elevators: elevators, maxFloors: maxFloors}
}

// Status returns the status (ID, floor and goal floor) of all elevators in
// the system.
func (es *elevatorSystem) Status() []Elevator {
	return es.elevators
}

// Pickup is the interface to request an elevator to the specified floor, this
// blocks until an elevator has been found.
func (es *elevatorSystem) Pickup(floor, direction int) (e Elevator) {
	err := ErrNoElevator

	// Continue this until we find an elevator
	for err == ErrNoElevator {
		e, err = es.nearestElevator(floor, Direction(direction))
	}

	e.GoToFloor(floor)
	return e
}

// Step is the ticker for the elevator system to run
func (es *elevatorSystem) Step() {
	for _, elevator := range es.elevators {
		elevator.Move()
	}
}

// nearestElevator gets the nearest elevator to a floor with the correct
// direction
func (es *elevatorSystem) nearestElevator(floor int, direction Direction) (Elevator, error) {
	var nearest Elevator
	score := es.maxFloors

	for _, e := range es.elevators {
		if e.Direction() == NotMoving || e.Direction() == direction {
			proximity := e.ProximityTo(floor)
			if proximity < score {
				score = proximity
				nearest = e
			}
		}
	}

	if nearest == nil {
		return nil, ErrNoElevator
	}

	return nearest, nil
}
