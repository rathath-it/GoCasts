package main

func main() {
	//slice
	cards := deck{newCard(), newCard()}
	// append does not mutate the slice
	cards = append(cards, "Another card")
	// loop through slice
	cards.print()

}

func newCard() string {
	return "Five of Dido"
}
