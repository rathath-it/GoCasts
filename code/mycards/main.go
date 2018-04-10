package main

import (
	"fmt"
)

func main() {
	//slice
	cards := []string{newCard(), newCard()}
	// append does not mutate the slice
	cards = append(cards, "Another card")
	// loop through slice
	for i, card := range cards {
		fmt.Println(i, card)
	}

}

func newCard() string {
	return "Five of Dido"
}
