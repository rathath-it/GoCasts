package main

import (
	"fmt"
)

// Create new type of 'deck' as slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Banana", "Faroula", "Toffaha"}
	cardValues := []string{"Two", "Three", "Four"}
	for i, suit := range cardSuits {
		for j, val := range cardValues {
			fmt.Println(i, j)
			cards = append(cards, suit+" "+val)
		}
	}
	return cards
}

// receiver : setup functions to variables type that was created
// "d" because it is the beginning of the type "deck".. one or two letters
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
