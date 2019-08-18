package main

//Animal is a stuct for animals that a player uses for game play.
type Animal struct {
	Name     string `json:"name"`
	Abr      string `json:"abr"`
	Active   bool   `json:"active"`
	House    string `json:"house"`
	Species  string `json:"species"`
	Life     int    `json:"life"`
	Strength int    `json:"strength"`
	Speed    int    `json:"speed"`
	Lethal   int    `json:"lethal"`
	Flight   bool   `json:"flight"`
	Swim     bool   `json:"swim"`
}

//Animals is a stuct contain Animal.
type Animals struct {
	All []Animal `json:"all animals"`
}

//AnimalList returns a slice of Animal.
func AnimalList(A []Animal) []string {
	var list []string
	for i := range A {
		list = append(list, (A[i].Abr + ") " + A[i].Name))
	}
	return list
}
