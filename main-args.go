package main

import (
	"fmt"
	"os"
)

func main() {
	// One way
	if len(os.Args) >= 2 {
		fmt.Println("You passed an argument : ", os.Args[1])
	} else {
		fmt.Println("No args passed")
	}
}
