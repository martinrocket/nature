package main

import (
	"fmt"
)

func main() {
	clearScreen()
	fmt.Println("...running nature.")
	a := GetLion()
	PAnimal(a)
	createCard(a)
	//buildCard(30, 40, a)

}
