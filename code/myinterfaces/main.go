package main

import (
	"fmt"
)

type person interface {
	getType() string
}

type student struct{}
type teacher struct{}

func main() {
	ahmad := student{}
	raki := teacher{}
	printType(ahmad)
	printType(raki)
}

func printType(p person) {
	fmt.Println(p.getType())
}

func (s student) getType() string {
	return "STUDENT_TYPE"
}

func (t teacher) getType() string {
	return "TEACHER_TYPE"
}
