package main

import (
	"fmt"
)

func getMenu() [][]string {
	m := [][]string{
		{"AC", "Add Cards"},
		{"RC", "Read Cards"},
		{"DC", "Display Cards"},
		{"AP", "Add Players"},
		{"AA", "Add Animals to Player"},
		{"DP", "Dispay Players"},
		{"GS", "Game Status"},
		{"SN", "Start New Game"},
		{"CG", "Continue Game"},
		{"X ", "Exit Game"},
	}
	return m
}

func printMenu(m [][]string) {
	for key, value := range m {
		fmt.Printf("%v: %v %v\n", key, value[0], value[1])
	}
}
