package main

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

func AnimalList() []string {
	var list []string
	return append(list, (GetLion().Abr + ") " + GetLion().Name), (GetAfricanDog().Abr + ") " + GetAfricanDog().Name), (GetThomGaz().Abr + ") " + GetThomGaz().Name))

}

func GetLion() Animal {
	l := Animal{}
	l.Active = true
	l.House = "Africa"
	l.Species = "Cat"
	l.Name = "Lion"
	l.Abr = "l"
	l.Life = 10
	l.Strength = 10
	l.Speed = 8
	l.Lethal = 10
	l.Flight = false
	l.Swim = false
	return l
}

func GetAfricanDog() Animal {
	l := Animal{}
	l.Active = true
	l.House = "Africa"
	l.Species = "Dog"
	l.Name = "African Dog"
	l.Abr = "d"
	l.Life = 10
	l.Strength = 5
	l.Speed = 5
	l.Lethal = 5
	l.Flight = false
	l.Swim = false
	return l
}

func GetThomGaz() Animal {
	return Animal{"Thompson's Gazelle", "tg", true, "Africa", "Bovidea", 10, 3, 10, 0, false, false}

}
