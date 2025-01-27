package worm

import "worm_go/internal/apple"

type Worm struct {
	cells []wormCell
}

func Make(length int, height int) Worm {
	return Worm{
		cells: []wormCell{
			{
				height:       height / 2,
				length:       length / 2,
				direction:    DirectionLeft,
				oldDirection: DirectionLeft,
			},
		},
	}
}

func (w *Worm) SetDirection(direction int) {
	w.cells[0].direction = direction
	w.cells[0].oldDirection = direction
}

func (w *Worm) Step(a *apple.Apple) {
	for i := 0; i < len(w.cells); i++ {
		if i > 0 {
			w.cells[i].step(w.cells[i-1].oldDirection)
		} else {
			w.cells[i].step(NullDirection)
		}
	}

	if w.cells[0].onApple(*a) {
		w.prolongation()
		a.Eat()
	}
}

func (w *Worm) prolongation() {
	newHead := w.cells[0]
	newHead.step(NullDirection)

	w.cells = append([]wormCell{newHead}, w.cells...)
}

func (w *Worm) IsWormCell(length int, height int) bool {
	for _, c := range w.cells {
		if c.length == length && c.height == height {
			return true
		}
	}
	return false
}
