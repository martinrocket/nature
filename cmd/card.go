package main

import (
	"fmt"
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

func buildBox(x int, y int, a animal) {
	b := boxDraw{}

	top := b.box1().uLeft + strings.Repeat(b.box1().vLine, x) + b.box1().uRight
	bottom := b.box1().bLeft + strings.Repeat(b.box1().vLine, x) + b.box1().bRight
	blank := b.box1().hLine + strings.Repeat(b.box1().space, x) + b.box1().hLine

	fmt.Print(top + "\n")
	z := (len(a.name) + len("Life:") + len(strconv.Itoa(a.life)))
	s := strconv.Itoa(len(a.name) + len("Life:") + len(strconv.Itoa(a.life)))
	fmt.Print(middleLeftJustify(x, s))

	fmt.Print(middleLeftJustify(x, (a.name + strings.Repeat(" ", z) + "Life:" + strconv.Itoa(a.life))))
	fmt.Print(middleLeftJustify(x, ("House:" + a.house + ", " + "Species:" + a.species)))

	fmt.Print(blank + "\n")
	fmt.Print(bottom + "\n")

}

func middleLeftJustify(x int, s string) string {
	b := boxDraw{}
	return b.box1().hLine + b.box1().bl + s + strings.Repeat(b.box1().space, (x-1-len(s))) + b.box1().hLine + "\n"
}

func card(a animal) {

}
