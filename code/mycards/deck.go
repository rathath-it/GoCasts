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
	// underscore (_) mean we don't want to use the variable to not throw error
	for _, suit := range cardSuits {
		for _, val := range cardValues {
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

// Return two values
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
