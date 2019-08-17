package main

import (
	"fmt"
	"strings"
)

const gameWidth = 60

func gameStatus() {
	gameTop()
	gameContent()
	gameBlank()
	gameBottom()

}

func gameContent() {
	l := len(fmt.Sprintf("%v", "Game Status"))
	fmt.Print(vLine + bl + fmt.Sprintf("%v", "Game Status") + (strings.Repeat(bl, (gameWidth - l - 1))) + vLine + "\n")

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
