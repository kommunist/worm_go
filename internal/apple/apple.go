package apple

import "math/rand"

type Apple struct {
	Height int
	Length int
	Status int
}

func (a *Apple) Eat() {
	a.Status = 1
}

func (a *Apple) IsEat() bool {
	return a.Status == 1
}

func (a *Apple) IsAppleCell(length int, height int) bool {
	return a.Length == length && a.Height == height
}

func (a *Apple) Replace(length int, height int) {
	a.Status = 0
	a.Height = rand.Intn(height-2) + 1
	a.Length = rand.Intn(length-2) + 1
}

func Make(length int, height int) Apple {
	a := Apple{}
	a.Replace(length, height)
	return a
}
