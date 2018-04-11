package main

import (
	"fmt"
)

func main() {
	// ahmad is a reference to a block (address) in the memory
	ahmad := getStudent("Ahmad")
	fmt.Println(ahmad.toString())
	// &var means give me the memory address of the value this var is pointing at
	ahmad.updateName("Mr Ahmed")
	fmt.Println(ahmad.toString())
}
