package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

//PP is a slice of Players
var PP []player

//Scanner is used for user input
var Scanner = bufio.NewScanner(os.Stdin)

func main() {
	creatLog()
	startApp()
}

func logPrint(t string) {
	fmt.Println(t)
	log.Println(t)
}

func startApp() {
	logPrint("Starting nature...")
	for {
		gameStatus(PP)
		m := getMenu()
		printMenu(m)
		fmt.Print("Enter Text: ")
		Scanner.Scan()
		text := Scanner.Text()
		log.Println("call to Scanner.Text, User selects: ", text)
		if len(text) != 0 {
			fmt.Println(text)
		}
		choice(strings.ToUpper(text))

	}

}

func choice(t string) {

	switch t {
	case "DC":
		listCards(true)
	case "AP":
		PP = append(PP, addPlayer())
		log.Printf("Added Player, player list: %v\n", PP)
	case "AA":
		addAnimalsPlayer(PP)
	case "X":
		os.Exit(0)
	}

}
