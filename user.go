package main

import (
	"fmt"
)

type User struct {
	atFloor    int
	toFloor    int
	inElevator bool

	elevators        ElevatorSystem
	assignedElevator Elevator

	blockChain chan struct{}
}

// NewUser creates a new User instance
func NewUser(es ElevatorSystem, atFloor int) *User {
	return &User{
		elevators: es,
		atFloor:   atFloor,
		blockChain: make(chan struct{}),
	}
}

// String will return a string representation of the user
func (u *User) String() string {
	if u.assignedElevator != nil {
		return fmt.Sprintf(
			"user id=%p, start-floor=%d, end-floor=%d, direction=%s, in-elevator=%t, elevator-id=%d",
			u,
			u.atFloor,
			u.toFloor,
			u.direction(),
			u.inElevator,
			u.assignedElevator.ID(),
		)
	}
	return fmt.Sprintf("user id=%p, start-floor=%d, end-floor=%d, direction=%s", u, u.atFloor, u.toFloor, u.direction())
}

// RequestElevator is a request for elevator pickup for a user
func (u *User) RequestElevator(toFloor int) {
	elevator := u.elevators.Pickup(u.atFloor, int(u.direction()))
	u.toFloor = toFloor
	u.assignedElevator = elevator
}

// Update is what is triggered by the notifier and will check elevator floors
func (u *User) Update() {
	fmt.Println(u)

	if u.assignedElevator == nil {
		return
	}

	// handle go to floor
	if u.assignedElevator.AtFloor(u.atFloor) && !u.inElevator {
		u.inElevator = true
		u.assignedElevator.GoToFloor(u.toFloor)
	}

	// handle going out of elevator
	if u.assignedElevator.AtFloor(u.toFloor) && u.inElevator {
		u.inElevator = false
		u.blockChain <- struct{}{}
	}
}

// Block will block until the update has triggered and the user is at the
// correct floor
func (u *User) Block() {
	for _ = range u.blockChain {
		close(u.blockChain)
	}
}

// direction  is a wrapper to check which direction the user is going
func (u *User) direction() Direction {
	if u.atFloor > u.toFloor {
		return Direction(-1)
	}
	return Direction(1)
}
