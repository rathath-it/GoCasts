package main

import (
	"fmt"
)

func main() {
	//slice
	cards := newDeck()
	// loop through slice
	cards.print()
	hand, remainingCard := deal(cards, 5)
	fmt.Println("----")
	fmt.Println("Hand", hand)
	fmt.Println("Remaining", remainingCard)
}
