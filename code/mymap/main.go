package main

import (
	"fmt"
)

func main() {
	// means all keys are string, all values are string
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#8cd123",
		"white": "#fff",
	}
	// Another way to declare map
	var procolors map[string]string

	// Another way again to declare map
	premcolors := make(map[string]string)
	premcolors["black"] = "#000"

	fmt.Println(colors)
	fmt.Println(procolors)
	fmt.Println(premcolors)
}
