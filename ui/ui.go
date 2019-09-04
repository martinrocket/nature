package ui

import "strings"

//LabelLength return the lenth of several strings with a SPACE between each string plus all the strings added together.
func LabelLength(sp int, s ...string) (int, string) {
	spaces := (strings.Repeat(" ", sp))
	st := ""
	for _, myString := range s {
		st = st + myString + spaces
	}
	return len(st), st
}
