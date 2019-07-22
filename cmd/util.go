package main

import (
	"os"
	"os/exec"
)

func clearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
	w := exec.Command("clr")
	w.Stdout = os.Stdout
	w.Run()

}
