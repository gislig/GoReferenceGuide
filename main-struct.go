package main

import "fmt"

type structname struct {
	name string
	age  int
}

func main() {
	// Create an instance of a struct
	r := structname{}
	r.name = "bob"
	r.age = 28
	fmt.Println(r)

	// Create an array of a struct
	y := [2]structname{}
	y[0].name = "joe"
	y[0].age = 48

	y[1].name = "smith"
	y[1].age = 37

	fmt.Println(y)
}
