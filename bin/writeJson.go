package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func writeJSON(a Animal) {
	d, err := json.MarshalIndent(a, "", "\t")
	if err != nil {
		log.Printf("Error marshalling %v \n", err)
	}
	fmt.Println(string(d))
	openCardFile(a, d)

}

func openCardFile(a Animal, d []byte) {

	filename := "card_" + a.Name + ".dat"
	f, err := os.Create(filename)
	checkErr(err)

	defer f.Close()
	_ = ioutil.WriteFile(filename, d, 0100644)

}
