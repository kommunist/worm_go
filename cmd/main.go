package main

import (
	"time"
	"worm_go/internal/apple"
	"worm_go/internal/field"
	"worm_go/internal/worm"

	"github.com/eiannone/keyboard"
)

const length = 30
const height = 20

func main() {
	f := field.Make(length, height)
	a := apple.Make(length, height)
	w := worm.Make(length, height)

	messages := make(chan int)
	go listenButtons(messages)

	for {
		field.Clear()

		select {
		case msg1 := <-messages:
			w.SetDirection(msg1)
		default:
			w.Step(&a)
		}

		f.Print(w, a)
		if a.IsEat() {
			a.Replace(length, height)
		}
		time.Sleep(70 * time.Millisecond)
	}
}

func listenButtons(directionChan chan int) {
	for {
		_, key, err := keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}
		switch key {
		case keyboard.KeyArrowUp:
			directionChan <- worm.DirectionUp
		case keyboard.KeyArrowDown:
			directionChan <- worm.DirectionDown
		case keyboard.KeyArrowLeft:
			directionChan <- worm.DirectionLeft
		case keyboard.KeyArrowRight:
			directionChan <- worm.DirectionRight
		}
	}
}
