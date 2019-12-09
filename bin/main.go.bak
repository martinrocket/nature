package main

import (
	"bufio"
	"fmt"
	"os"
)

//PP is a slice of Players
var PP []player

func main() {

	clearScreen()
	fmt.Println("...running nature.")
	scanner := bufio.NewScanner(os.Stdin) // Scans in Text for user input
	for {

		gameStatus(PP)
		fmt.Println("Selection of Animals:")
		A := readCards()
		l := AnimalList(A)
		for i := range l {
			fmt.Println(l[i])
		}
		fmt.Println("p) Player")
		fmt.Println("x) Exit")
		fmt.Print("Enter Text: ")
		scanner.Scan()
		text := scanner.Text()
		if len(text) != 0 {
			fmt.Println(text)
		}
		for i := range l {
			if text == A[i].Abr {
				createCard(A[i])
				animalMenu(A[i])

			} else if text == "p" {
				p := players()

				PP = append(PP, p)
				break

			} else if text == "x" {
				os.Exit(0)
			}
		}

	}

}
