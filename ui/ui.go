package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

//LabelLength return the lenth of several strings with a SPACE between each string plus all the strings added together.
func LabelLength(sp int, s ...string) (int, string) {
	spaces := (strings.Repeat(" ", sp))
	st := ""
	for _, myString := range s {
		st = st + myString + spaces
	}
	return len(st), st
}

//Pause for user to hit enter
func Pause() {
	fmt.Print("Press Enter to Continue: ")
	scanner.Scan()
}

//GetInput allows a question with a string return
func GetInput(s string) string {
	fmt.Print(s + ": ")
	scanner.Scan()
	text := scanner.Text()
	return text
}
