package main

type wormStruct struct {
	cells []wormCell
}

func (worm *wormStruct) step(apple *appleCell) {
	for i := 0; i < len(worm.cells); i++ {
		if i > 0 {
			worm.cells[i].step(worm.cells[i-1].oldDirection)
		} else {
			worm.cells[i].step(nullDirection)
		}
	}

	if worm.cells[0].onApple(*apple) {
		worm.prolongation()
		apple.init()
	}
}

func (worm *wormStruct) init(length int, height int) {
	worm.cells = []wormCell{
		wormCell{height: height / 2, length: length / 2, direction: directionLeft, oldDirection: directionLeft},
	}
}

func (worm *wormStruct) prolongation() {
	newHead := worm.cells[0]
	newHead.step(nullDirection)

	worm.cells = append([]wormCell{newHead}, worm.cells...)
}
