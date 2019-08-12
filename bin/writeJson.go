package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func writeJson() {
	l := GetLion()
	d, err := json.MarshalIndent(l, "", "\t")
	if err != nil {
		log.Printf("Error marshalling %v \n", err)
	}
	fmt.Println(string(d))
}
