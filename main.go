package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"log"
	"os"
	"os/exec"
	"time"
)

const length = 30
const height = 20

const cellEmpty = 0
const cellOfWorm = 1
const cellOfApple = 2

func clearOutput() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}

func printField(worm wormStruct, apple appleCell) {
	var field [height][length]int

	for _, v := range worm.cells {
		field[v.height][v.length] = cellOfWorm
	}

	field[apple.height][apple.length] = cellOfApple

	for _, line := range field {
		for _, column := range line {
			switch column {
			case cellEmpty:
				fmt.Print(".")
			case cellOfWorm:
				fmt.Print("*")
			case cellOfApple:
				fmt.Print("&")
			}
		}
		fmt.Println("")
	}
}

func main() {

	var worm wormStruct
	worm.init(length, height)
	messages := make(chan int)

	var apple appleCell
	apple.init()

	go func(directionChan chan int) {
		for {
			_, key, err := keyboard.GetSingleKey()
			if err != nil {
				panic(err)
			}
			switch key {
			case keyboard.KeyArrowUp:
				directionChan <- directionUp
			case keyboard.KeyArrowDown:
				directionChan <- directionDown
			case keyboard.KeyArrowLeft:
				directionChan <- directionLeft
			case keyboard.KeyArrowRight:
				directionChan <- directionRight
			}
		}
	}(messages)

	for true {
		clearOutput()

		select {
		case msg1 := <-messages:
			worm.cells[0].direction = msg1
			worm.cells[0].oldDirection = msg1
		default:
			worm.step(&apple)
		}

		printField(worm, apple)

		time.Sleep(70 * time.Millisecond)
	}
}
