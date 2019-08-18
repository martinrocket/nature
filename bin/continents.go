package main

//Continent is struct of a Continents Name and an Abreviation
type Continent struct {
	Name string `json:"Name"`
	Abr  string `json:"Abreviation"`
}

//Continents returns a slice of Continents with Name and Abreviation
func Continents() []Continent {
	var c []Continent
	c = append(c, Continent{"Africa", "af"})
	c = append(c, Continent{"Asia", "as"})
	c = append(c, Continent{"Australia", "au"})
	c = append(c, Continent{"Europe", "eu"})
	c = append(c, Continent{"North America", "na"})
	c = append(c, Continent{"South America", "sa"})
	return c
}
