package main

import (
	"bufio"
	"fmt"
	"os"
)

func animalMenu(a Animal) {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Animal Menu, please choose.")
		fmt.Println("a) Add Animal to my herd.")
		fmt.Println("b) Write Animal to file.")
		fmt.Println("x) Exit to previous.")
		fmt.Print("Enter Text: ")
		scanner.Scan()
		text := scanner.Text()
		if len(text) != 0 {
			fmt.Println(text)
		}
		switch text {
		case "x":
			return
		case "a":
			createCard(a)
		case "b":
			writeJson(a)
		default:
			break
		}

	}
}
