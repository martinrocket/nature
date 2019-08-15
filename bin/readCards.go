package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

func readCards() {
	clearScreen()
	dir, err := ioutil.ReadDir(".")
	checkErr(err)
	cards := []string{}
	for _, file := range dir {
		if strings.Contains(file.Name(), "card_") {
			cards = append(cards, file.Name())
		}
	}
	for i := range cards {
		file, _ := ioutil.ReadFile(cards[i])

		data := Animal{}

		_ = json.Unmarshal([]byte(file), &data)

		fmt.Printf("read a card %v \n", data)
	}
}
