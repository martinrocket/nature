package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type boxDraw struct {
	vLine  string
	uRight string
	uLeft  string
	space  string
	hLine  string
	bRight string
	bLeft  string
	bl     string
}

func (m boxDraw) box1() boxDraw {
	m.vLine = "─"
	m.uRight = "┐"
	m.uLeft = "┌"
	m.space = " "
	m.hLine = "│"
	m.bRight = "┘"
	m.bLeft = "└"
	m.bl = " "
	return m

}

func buildCard(x int, y int, a animal) {
	b := boxDraw{}

	top := b.box1().uLeft + strings.Repeat(b.box1().vLine, x) + b.box1().uRight
	bottom := b.box1().bLeft + strings.Repeat(b.box1().vLine, x) + b.box1().bRight
	blank := b.box1().hLine + strings.Repeat(b.box1().space, x) + b.box1().hLine

	r := reflect.ValueOf(&a).Elem()
	typeOfT := r.Type()
	for i := 0; i < r.NumField(); i++ {
		fmt.Printf("%v: %v \n", typeOfT.Field(i).Name, r.Field(i))

	}
	//c := cardDetail(a animal)
	f := typeOfT.Field(4).Name
	z := (len(a.name) + len(typeOfT.Field(4).Name) + len(strconv.Itoa(a.life)))
	c := x - z - 4
	sc := strconv.Itoa(a.life)
	fmt.Print(top + "\n")

	fmt.Print(middleLeftJustify(x, (a.name + strings.Repeat(" ", c) + f + ": " + sc)))
	fmt.Print(blank + "\n")
	fmt.Print(middleLeftJustify(x, ("House:" + a.house + ", " + "Species:" + a.species)))
	fmt.Print(middleLeftJustify(x, ("Lethal:" + strconv.Itoa(a.lethal) + ", " + "Strength:" + strconv.Itoa(a.strength))))
	fmt.Print(middleLeftJustify(x, ("Flight:" + strconv.FormatBool(a.flight) + ", " + "Strength:" + strconv.Itoa(a.speed))))

	fmt.Print(bottom + "\n")

}

func middleLeftJustify(x int, s string) string {

	b := boxDraw{}
	return b.box1().hLine + b.box1().bl + s + strings.Repeat(b.box1().space, (x-1-len(s))) + b.box1().hLine + "\n"
}

// func cardDetail(a animal) (f, sc string, c int) {
// 	r := reflect.ValueOf(&a).Elem()
// 	typeOfT := r.Type()
// 	f := typeOfT.Field(4).Name
// 	z := (len(a.name) + len(typeOfT.Field(4).Name) + len(strconv.Itoa(a.life)))
// 	c := x - z - 4
// 	sc := strconv.Itoa(a.life)

// }
