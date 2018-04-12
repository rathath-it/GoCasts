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

	fmt.Println(colors)
}
