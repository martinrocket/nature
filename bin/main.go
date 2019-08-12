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
		fmt.Print("Enter Text: ")
		// Scans a line from Stdin(Console)
		scanner.Scan()
		// Holds the string that scanned
		text := scanner.Text()
		if len(text) != 0 {
			fmt.Println(text)
		}
		switch text {
		case "x":
			os.Exit(0)
		case "l":
			createCard(GetLion())
		default:
			break
		}

	}

	l := GetLion()
	d := GetAfricanDog()
	tm := GetThomGaz()
	//PAnimal(a)
	createCard(l)
	createCard(d)
	createCard(tm)

}
