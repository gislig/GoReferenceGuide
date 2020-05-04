package main

import "fmt"

// oneParam : Has only one parameter
func oneParam(x int) {
	fmt.Println(x)
}

// twoParamReturnInt : Takes two parameters and returns any integer
func twoParamReturnInt(x int, y int) int {
	return x + y
}

// twoParamesReturnsTwoInts : Takes two parameters and returns two integers
func twoParamesReturnsTwoInts(x, y int) (z int, t int) {
	return x, y
}

// multiParam : Takes multiple parameters, seperated with comma e.g. multiParam(1, 2, 3)
func multiParam(args ...int) {
	s := 0
	for i := 0; i < len(args); i++ {
		s += args[i]
	}
	fmt.Println(s)
}

func main() {
	fmt.Println("Testing funcs")
	oneParam(1)
	fmt.Println(twoParamReturnInt(1, 2))
	fmt.Println(twoParamesReturnsTwoInts(1, 2))
	multiParam(1, 2, 3, 4)
}
