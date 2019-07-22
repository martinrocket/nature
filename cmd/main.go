package main

import (
	"fmt"
)

func main() {
	clearScreen()
	fmt.Println("...running nature.")
	a := GetLion()
	PAnimal(a)
	buildBox(40, 40, a)

}
