package main

import "fmt"

func main() {
	// Create a new slice
	slice := []string{"item1", "item2"}
	fmt.Println(slice)

	// Append to slice
	slice = append(slice, "item3")
	fmt.Println(slice)
}
