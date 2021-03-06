package main

import (
	"log"
	"os"
	"os/exec"
	"runtime"
)

func clearScreen() {
	switch myOs := runtime.GOOS; myOs {
	case "darwin":
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()
	case "linux":
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()
	default:
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

}

func checkErr(e error) {
	if e != nil {
		log.Printf(" Error here %v /n", e)
	}

}
