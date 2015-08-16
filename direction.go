package main

const (
	Up        = 1
	NotMoving = 0
	Down      = -1
)

// direction is the type for direction, -1 is down and 1 is up
type Direction int

// String is the string representation of a direction
func (d Direction) String() string {
	switch d {
	case Up:
		return "up"
	case Down:
		return "down"
	default:
		return "n/a"
	}
}
