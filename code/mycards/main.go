package main

import (
	"fmt"
)

func main() {
	//slice
	cards := []string{newCard(), newCard()}
	fmt.Println(cards)
}

func newCard() string {
	return "Five of Dido"
}
