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
		case "d":
			createCard(GetAfricanDog())
		case "tg":
			createCard(GetThomGaz())
		default:
			break
		}

	}

}
