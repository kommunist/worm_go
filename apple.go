package main

import "math/rand"

type appleCell struct {
	height int
	length int
}

func (apple *appleCell) init() {
	apple.height = rand.Intn(height-2) + 1
	apple.length = rand.Intn(length-2) + 1
}
