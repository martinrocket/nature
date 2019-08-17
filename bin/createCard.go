package main

import (
	"fmt"
	"reflect"
	"strings"
)

const hLine = "─"
const uRight = "┐"
const uLeft = "┌"
const space = " "
const vLine = "│"
const bRight = "┘"
const bLeft = "└"
const bl = " "
const cardWidth = 40

func createCard(a Animal) {
	top()
	content(a)
	bottom()

}

func top() {
	fmt.Print(uLeft + strings.Repeat(hLine, cardWidth) + uRight + "\n")
}
func bottom() {
	fmt.Print(bLeft + strings.Repeat(hLine, cardWidth) + bRight + "\n")
}

func content(a Animal) {
	fields := reflect.TypeOf(a)
	values := reflect.ValueOf(a)
	l := len(fmt.Sprintf("%v", values.Field(0)))
	fmt.Print(vLine + bl + fmt.Sprintf("%v", values.Field(0)) + (strings.Repeat(bl, (cardWidth - l - 1))) + vLine + "\n")
	for i := 1; i < values.Type().NumField(); i = i + 2 {
		if i == (values.Type().NumField() - 1) {
			break
		}
		field1 := fields.Field(i)
		field2 := fields.Field(i + 1)
		value1 := values.Field(i)
		value2 := values.Field(i + 1)
		s1 := field1.Name
		s2 := field2.Name
		v1 := fmt.Sprintf("%v", value1)
		v2 := fmt.Sprintf("%v", value2)
		l := len(s1) + len(": ") + len(fmt.Sprintf("%v", value1)) + len(s2) + len(": ") + len(fmt.Sprintf("%v", value2)) + 1
		fmt.Print(vLine + bl + field1.Tag.Get("json") + ": " + v1 + (strings.Repeat(bl, (cardWidth - l - 1))) + field2.Tag.Get("json") + ": " + v2 + bl + vLine + "\n")

	}

}
