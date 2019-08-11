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
const cardWidth = 40

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
	for i := 0; i < values.Type().NumField(); i = i + 2 {
		field1 := fields.Field(i)
		field2 := fields.Field(i + 1)
		value1 := values.Field(i)
		value2 := values.Field(i + 1)
		s1 := field1.Name
		s2 := field2.Name
		v1 := fmt.Sprintf("%v", value1)
		v2 := fmt.Sprintf("%v", value2)
		l := len(s1) + len(": ") + len(fmt.Sprintf("%v", value1)) + len(s2) + len(": ") + len(fmt.Sprintf("%v", value2)) + 1

		//fmt.Print(s, ": ", value, " ", l, "\n")

		fmt.Print(hLine + bl + field1.Tag.Get("json") + ": " + v1 + (strings.Repeat(bl, (cardWidth - l - 1))) + field2.Tag.Get("json") + ": " + v2 + bl + hLine + "\n")

		//fmt.Printf((val.Type().Field(i).Tag.Get("json")) + "\n")
	}
}

func sLength(a, b string) int {
	return len(a) + len(b)
}
