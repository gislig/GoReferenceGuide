package main

import (
	"fmt"
)

func main() {
	fmt.Println("Trying for loops")

	// Classic for loop
	for i := 0; i < 3; i++ {
		fmt.Println("Classic for loop count : ", i)
	}

	// Condition for loop
	j := -3
	for j != 0 {
		fmt.Println("Similar to while : ", j)
		j++
	}

	// Infinate loop
	r := 0
	for {
		fmt.Println("Infinate loop : ", r)
		r++
		if r == 3 {
			break
		}
	}

}
