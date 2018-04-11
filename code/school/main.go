package main

import (
	"fmt"
)

func main() {
	ahmad := getStudent("Ahmad")
	fmt.Println(ahmad.toString())
	pointerToAhmed := &ahmad
	pointerToAhmed.updateName("Mr Ahmed")
	fmt.Println(ahmad.toString())
}
