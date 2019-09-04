package main

import (
	"fmt"
	"strings"

	"github.com/martinrocket/nature/ui"
)

const gameWidth = 60

func gameStatus(p []player) {
	clearScreen()
	gameTop()
	gameContent(p)
	gameBlank()
	gameBottom()
	fmt.Println(p)

}

func gameContent(p []player) {
	l := len(fmt.Sprintf("%v", "Game Status"))

	fmt.Print(vLine + bl + fmt.Sprintf("%v", "Game Status") + (strings.Repeat(bl, (gameWidth - l - 1))) + vLine + "\n")
	for i := range p {
		//l2 := len((p[i].Name + " " + p[i].Continent.Name))
		l1, l2 := ui.LabelLength(4, p[i].Name, p[i].Continent.Name)
		//name := p[i].Name + " " + p[i].Continent.Name
		fmt.Print(vLine + bl + fmt.Sprintf("%v", l2) + (strings.Repeat(bl, (gameWidth - l1 - 1))) + vLine + "\n")

	}
}

func gameTop() {
	fmt.Print(uLeft + strings.Repeat(hLine, gameWidth) + uRight + "\n")
}
func gameBottom() {
	fmt.Print(bLeft + strings.Repeat(hLine, gameWidth) + bRight + "\n")
}

func gameBlank() {
	fmt.Print(vLine + strings.Repeat(bl, gameWidth) + vLine + "\n")
}
