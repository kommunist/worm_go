package field

import "fmt"

type Cell int

func (c *Cell) Print() {
	val := *c
	switch val {
	case 0:
		fmt.Print(".")
	case 1:
		fmt.Print("*")
	case 2:
		fmt.Print("&")
	}
}
