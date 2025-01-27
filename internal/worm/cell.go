package worm

import "worm_go/internal/apple"

type wormCell struct {
	height       int
	length       int
	direction    int
	oldDirection int
}

// directions
const NullDirection = 0
const DirectionUp = 1
const DirectionDown = 2
const DirectionLeft = 3
const DirectionRight = 4

func (cell *wormCell) step(headDirection int) {
	switch cell.direction {
	case DirectionUp:
		cell.stepUp()
	case DirectionDown:
		cell.stepDown()
	case DirectionLeft:
		cell.stepLeft()
	case DirectionRight:
		cell.stepRight()
	}

	if headDirection != NullDirection {
		cell.oldDirection = cell.direction
		cell.direction = headDirection
	}
}

func (cell *wormCell) onApple(apple apple.Apple) bool {
	return cell.height == apple.Height && cell.length == apple.Length
}

func (cell *wormCell) stepUp() {
	cell.height = cell.height - 1
}

func (cell *wormCell) stepDown() {
	cell.height = cell.height + 1
}

func (cell *wormCell) stepLeft() {
	cell.length = cell.length - 1
}

func (cell *wormCell) stepRight() {
	cell.length = cell.length + 1
}
