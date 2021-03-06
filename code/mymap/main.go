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

	printMap(colors)

	// Another way to declare map
	// var procolors map[int]string
	// procolors = map[int]string{
	// 	23: "#f23412",
	// }
	// procolors[10] = "#000"
	// delete(procolors, 10)
	// // Another way again to declare map
	// premcolors := make(map[string]string)
	// premcolors["black"] = "#000"

	fmt.Println(colors)
	// fmt.Println(procolors)
	// fmt.Println(premcolors)
}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Println("key: ", key, "-> value: ", value)
	}
}
