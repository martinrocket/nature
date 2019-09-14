package main

import (
	"bufio"
	"fmt"
	"os"
)

type player struct {
	Name      string `json:"name"`
	Continent `json:"continent"`
	Herd      []Animal `json:"herd"`
}

func players() player {
	p := player{}

	scanner := bufio.NewScanner(os.Stdin) // Scans in Text for user input
	fmt.Print("Enter your name: \n")
	scanner.Scan()

	p.Name = scanner.Text()
	fmt.Println("The following continets are available:")
	cl := Continents()
	for i := range cl {
		fmt.Print(cl[i].Abr + ") " + cl[i].Name + "\n")
	}
	fmt.Print("Choose your Contient: \n")
	scanner.Scan()
	t := string(scanner.Text())
	switch t {
	case "af":
		p.Continent = cl[0]
	case "as":
		p.Continent = cl[1]
	case "au":
		p.Continent = cl[2]
	case "eu":
		p.Continent = cl[3]
	case "na":
		p.Continent = cl[4]
	case "sa":
		p.Continent = cl[5]
	default:
		players()
		break

	}

	fmt.Println("Choose your Animals, 3 to start:")
	switch p.Continent.Abr {
	case "af":

	}

	scanner.Scan()

	return p

}
