package main

import (
	"fmt"
	"log"
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
	p.Name = ui.GetInput("Enter your name")
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
	t := ui.GetInput("Choose your Continent")
	return t
}

func addAnimalsPlayer(p []player) {
	playerListChoice(p)
	a := ui.GetInput("Adding Animals to Player, Choose a player by their number")
	for key, value := range p {
		if a == strconv.Itoa(key) {
			logPrint("You chose " + value.Name + "\n")
			playerInfo(value.Name)
		}
	}
	fmt.Printf("Choose from the list of animals.  ")
	listCards(false)
	b := ui.GetInput("Choose animal")
	b = strings.ToLower(b)
	logPrint("Chose: " + b)
	for {
		switch b {
		case "tg":
			fmt.Printf("%v \n", b)
			ui.GetInput("Nice Choice")
			return
		case "x":
			return

		}
	}
	ui.Pause()
	return

}
func playerListChoice(p []player) {
	clearScreen()
	gameTop()
	playersContent(p)
	log.Printf("Call to playerListChoice, Logging game content: %v", p)
	gameBlank()
	gameBottom()
}

func playersContent(p []player) {
	l := len(fmt.Sprintf("%v", "Players:"))

	fmt.Print(vLine + bl + fmt.Sprintf("%v", "Players:") + (strings.Repeat(bl, (gameWidth - l - 1))) + vLine + "\n")
	for i := range p {
		num := strconv.Itoa(i) + ": "
		l1, l2 := ui.LabelLength(1, num, p[i].Name, p[i].Continent.Abr)
		fmt.Print(vLine + bl + fmt.Sprintf("%v", l2) + (strings.Repeat(bl, (gameWidth - l1 - 1))) + vLine + "\n")

	}
}

func playerInfo(p string) {
	clearScreen()
	gameTop()
	playerInfoDisplay(p)
	log.Printf("Call to playerListChoice, Logging game content: %v", p)
	gameBlank()
	gameBottom()
}

func playerInfoDisplay(p string) {
	l := len(fmt.Sprintf("%v", "Player Info:"))
	fmt.Print(vLine + bl + fmt.Sprintf("%v", "Player Info:") + (strings.Repeat(bl, (gameWidth - l - 1))) + vLine + "\n")
	l1, l2 := ui.LabelLength(1, p)
	fmt.Print(vLine + bl + fmt.Sprintf("%v", l2) + (strings.Repeat(bl, (gameWidth - l1 - 1))) + vLine + "\n")

}
