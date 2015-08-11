package main

import (
	"fmt"
)

func main() {

	// Create a new elevator system
	e := New(2, 20)

	// Pick up a person at 2, going up
	_, err := e.Pickup(2, 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Tick the world
	e.Step()

	// Print the status
	TickStatus(e, 1)

	// Pick up another user at floor 3, going down
	p, err := e.Pickup(3, -1)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Tick the world
	e.Step()

	// Set the goal to be 1 for the elevator that picked up user at floor
	// 3
	p.SetGoal(1)

	// Print status
	TickStatus(e, 2)

	// Tick the st
	_, err = e.Pickup(3, 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Tick and print status
	e.Step()
	TickStatus(e, 3)

	// Tick and print status
	e.Step()
	TickStatus(e, 4)

}

// TickStatus prints tick and elevator statuses.
func TickStatus(es *ElevatorSystem, tickId int) {
	fmt.Println("Tick", tickId)
	for _, elevator := range es.Status() {
		fmt.Printf("%s\n", elevator)
	}
}
