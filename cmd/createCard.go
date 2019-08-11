package main

import (
	"fmt"
	"reflect"
	"strings"
)

const vLine = "─"
const uRight = "┐"
const uLeft = "┌"
const space = " "
const hLine = "│"
const bRight = "┘"
const bLeft = "└"
const bl = " "
const cardWidth = 50

func createCard(a Animal) {

	fmt.Println("Called Create Card")
	fmt.Println(a)
	fmt.Printf("%+v\n", a)

	top()
	content(a)
	bottom()

}

func top() {
	fmt.Print(uLeft + strings.Repeat(vLine, cardWidth) + uRight + "\n")
}
func bottom() {
	fmt.Print(bLeft + strings.Repeat(vLine, cardWidth) + bRight + "\n")
}

func content(a Animal) {
	fields := reflect.TypeOf(a)
	values := reflect.ValueOf(a)
	for i := 0; i < values.Type().NumField(); i++ {
		field := fields.Field(i)
		value := values.Field(i)
		s := field.Name
		v := fmt.Sprintf("%v", value)
		l := len(s) + len(": ") + len(fmt.Sprintf("%v", value))

		//fmt.Print(s, ": ", value, " ", l, "\n")

		fmt.Print(hLine + bl + field.Name + ": " + v + (strings.Repeat(bl, (cardWidth - l - 1))) + hLine + "\n")

		//fmt.Printf((val.Type().Field(i).Tag.Get("json")) + "\n")
	}
}

func sLength(a, b string) int {
	return len(a) + len(b)
}
