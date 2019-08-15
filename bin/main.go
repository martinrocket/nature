package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	clearScreen()
	fmt.Println("...running nature.")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Selection of Animals:")
		l := AnimalList()
		for i := range l {
			fmt.Println(l[i])
		}
		fmt.Println("c) Read in Animal Cards")
		fmt.Print("Enter Text: ")
		scanner.Scan()
		text := scanner.Text()
		if len(text) != 0 {
			fmt.Println(text)
		}
		switch text {
		case "x":
			os.Exit(0)
		case "l":
			createCard(GetLion())
			animalMenu(GetLion())
		case "d":
			createCard(GetAfricanDog())
			animalMenu(GetAfricanDog())
		case "tg":
			createCard(GetThomGaz())
			animalMenu(GetThomGaz())
		case "c":
			readCards()
		default:
			break
		}

	}

}
