package main

import (
	"fmt"
)

func main() {
	ahmad := getStudent("Ahmad")
	fmt.Println(ahmad.toString())
	ahmad.updateName("Mr Ahmed")
	fmt.Println(ahmad.toString())
}
