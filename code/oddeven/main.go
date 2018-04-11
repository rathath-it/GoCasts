package main

import (
	"fmt"
)

func main() {
	numbers := []mynumber{34, 65, 33, 33, 19, 22, 71, 349}
	for _, num := range numbers {
		fmt.Println(num.feedback())
	}
	fmt.Println(numbers)

}
