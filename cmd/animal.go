package main

import "fmt"

type Animal struct {
	Active   bool   `json:"active"`
	Name     string `json:"name"`
	House    string `json:"house"`
	Species  string `json:"species"`
	Life     int    `json:"life"`
	Strength int    `json:"strength"`
	Speed    int    `json:"speed"`
	Lethal   int    `json:"lethal"`
	Flight   bool   `json:"flight"`
	Swim     bool   `json:"swim"`
}

func GetLion() Animal {
	l := Animal{}
	l.Active = true
	l.House = "Africa"
	l.Species = "Cat"
	l.Name = "Lion"
	l.Life = 10
	l.Strength = 10
	l.Speed = 8
	l.Lethal = 10
	l.Flight = false
	l.Swim = false
	return l
}

func PAnimal(a Animal) {
	fmt.Printf("Name: %v \n Information: %+v \n", a.Name, a)

}
