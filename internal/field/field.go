package field

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"worm_go/internal/apple"
	"worm_go/internal/worm"
)

type Field struct {
	length int
	height int
}

func Make(length int, height int) *Field {
	item := Field{length: length, height: height}
	return &item
}

func (f *Field) Print(w worm.Worm, a apple.Apple) {
	for i := 0; i < f.height; i++ {
		for j := 0; j < f.length; j++ {
			if w.IsWormCell(j, i) {
				fmt.Print("*")
			} else if a.IsAppleCell(j, i) {
				fmt.Print("@")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}

func Clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}
