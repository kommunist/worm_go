package main

type wormCell struct {
  height int
  length int
  direction int
  oldDirection int
}


// directions
const nullDirection = 0
const directionUp = 1
const directionDown = 2
const directionLeft = 3
const directionRight = 4


func (cell *wormCell) step(headDirection int) {
  switch cell.direction {
  case directionUp:
    cell.stepUp()
  case directionDown:
    cell.stepDown()
  case directionLeft:
    cell.stepLeft()
  case directionRight:
    cell.stepRight()
  }

  if headDirection != nullDirection {
    cell.oldDirection = cell.direction
    cell.direction = headDirection
  }
}

func (cell *wormCell) onApple (apple appleCell) bool {
  return cell.height == apple.height && cell.length == apple.length
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
