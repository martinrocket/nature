package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/martinrocket/nature/ui"
)

func listCards(b bool) {
	//listCards receives a bool to pause (true) or not pause (false)
	fmt.Println("Selection of Animals:")
	A := readCards()
	log.Println("Call to readCards, Selection of Animal Cards: ", A)
	l := AnimalList(A)
	for i := range l {
		fmt.Println(l[i])
	}
	if b == true {
		ui.Pause()
	}

}

func readCards() []Animal {
	//clearScreen()
	dir, err := ioutil.ReadDir(".")
	checkErr(err)
	cards := []string{} //slice to track Card Names
	var numCards int    //The number of cards we have imported
	for _, file := range dir {
		if strings.Contains(file.Name(), "card_") {
			cards = append(cards, file.Name())
			numCards = numCards + 1
		}
	}
	A := []Animal{}
	for i := range cards {
		file, _ := ioutil.ReadFile(cards[i])

		data := Animal{}

		_ = json.Unmarshal([]byte(file), &data)
		A = append(A, data)
	}
	return A
}
