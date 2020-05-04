package main

import "fmt"

func main() {
	fmt.Println("Testing arrays")

	// Predefined array contains 3 items
	var arr1 [3]string
	arr1[0] = "item 1"
	arr1[1] = "item 2"
	arr1[2] = "item 3"
	fmt.Printf("arr1 contains : %v\n", arr1)

	var arr2 [2][2]string
	arr2[0][0] = "item 1.1"
	arr2[0][1] = "item 1.2"
	arr2[1][0] = "item 2.1"
	arr2[1][1] = "item 2.2"
	fmt.Printf("arr2 contains : %v\n", arr2)
}
