package position

import "errors"

type Position struct {
	X int
	Y int
}

func (p Position) Move(d Direction) Position {
	return Position{
		X: p.X + d.getXDelta(),
		Y: p.Y + d.getYDelta(),
	}
}

type Direction int64

const (
	DirectionNorth Direction = iota
	DirectionSouth
	DirectionEast
	DirectionWest
)

func ParseDirection(str string) (Direction, error) {
	switch str {
	case "n":
		fallthrough
	case "north":
		return DirectionNorth, nil

	case "s":
		fallthrough
	case "south":
		return DirectionSouth, nil

	case "e":
		fallthrough
	case "east":
		return DirectionEast, nil

	case "w":
		fallthrough
	case "west":
		return DirectionWest, nil
	}
	return DirectionNorth, errors.New("Unknown direction")
}

func (d Direction) ToString() string {
	switch d {
	case DirectionEast:
		return "East"
	case DirectionWest:
		return "West"
	case DirectionSouth:
		return "South"
	case DirectionNorth:
		return "North"
	default:
		panic("Unknown direction!")
	}
}

func (d Direction) getXDelta() int {
	switch d {
	case DirectionEast:
		return 1
	case DirectionWest:
		return -1
	default:
		return 0
	}
}

func (d Direction) getYDelta() int {
	switch d {
	case DirectionSouth:
		return 1
	case DirectionNorth:
		return -1
	default:
		return 0
	}
}
