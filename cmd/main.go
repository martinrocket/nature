package main

import (
	"fmt"
)

func main() {
	clearScreen()
	fmt.Println("...running nature.")
	l := GetLion()
	d := GetAfricanDog()
	tm := GetThomGaz()
	//PAnimal(a)
	createCard(l)
	createCard(d)
	createCard(tm)

}
