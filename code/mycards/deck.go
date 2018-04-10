package main

import (
	"fmt"
)

// Create new type of 'deck' as slice of strings
type deck []string

// receiver : setup functions to variables type that was created
// "d" because it is the beginning of the type "deck".. one or two letters
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
