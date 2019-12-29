package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/martinrocket/nature/ui"
)

type player struct {
	Name      string `json:"name"`
	Continent `json:"continent"`
	Herd      []Animal `json:"herd"`
}

func addPlayer() player {
	p := player{}
	scanner := bufio.NewScanner(os.Stdin) // Scans in Text for user input
	fmt.Print("Enter your name: \n")
	scanner.Scan()

	p.Name = scanner.Text()
	for {

		t := chooseContinent()
		cl := Continents()
		switch t {
		case "af":
			p.Continent = cl[0]
			return p
		case "as":
			p.Continent = cl[1]
			return p
		case "au":
			p.Continent = cl[2]
			return p
		case "eu":
			p.Continent = cl[3]
			return p
		case "na":
			p.Continent = cl[4]
			return p
		case "sa":
			p.Continent = cl[5]
			return p
		}

	}
}

func chooseContinent() string {

	fmt.Println("The following continets are available:")
	cl := Continents()
	for i := range cl {
		fmt.Print(cl[i].Abr + ") " + cl[i].Name + "\n")
	}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Choose your Contient: \n")
	scanner.Scan()
	t := string(scanner.Text())
	return t
}

func addAnimalsPlayer(p []player) {
	playerListChoice(p)
	a := ui.GetInput("Adding Animals to Player, Choose a player by their number")
	for key, value := range p {
		if a == strconv.Itoa(key) {
			logPrint("You chose " + value.Name + "\n")
		}
	}
	fmt.Printf("Choose from the list of animals.  ")
	listCards(false)
	b := ui.GetInput("Choose animal")
	logPrint("Chose: " + b)
	ui.Pause()

}

func playerListChoice(p []player) {
	clearScreen()
	gameTop()
	playerContent(p)
	log.Printf("Call to playerListChoice, Logging game content: %v", p)
	gameBlank()
	gameBottom()
}

func playerContent(p []player) {
	l := len(fmt.Sprintf("%v", "Players:"))

	fmt.Print(vLine + bl + fmt.Sprintf("%v", "Players:") + (strings.Repeat(bl, (gameWidth - l - 1))) + vLine + "\n")
	for i := range p {
		//l2 := len((p[i].Name + " " + p[i].Continent.Name))
		num := strconv.Itoa(i) + ": "
		l1, l2 := ui.LabelLength(1, num, p[i].Name, p[i].Continent.Abr)
		//name := p[i].Name + " " + p[i].Continent.Name
		fmt.Print(vLine + bl + fmt.Sprintf("%v", l2) + (strings.Repeat(bl, (gameWidth - l1 - 1))) + vLine + "\n")

	}
}
