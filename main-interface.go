package main

import "fmt"

// Þetta er classinn
type interfacename interface {
	thisIsATest() string
}

// Þetta eru variables sem á að vinna með fyrir viðeigandi function initið
type athing struct {
	mything string
}

// Þetta eru definitions
func (a athing) thisIsATest() string {
	return a.mything
}

func main() {
	fmt.Println("interface")

	var m = athing{
		mything: "bob",
	}

	items := []interfacename{m}

	fmt.Println(items[0].thisIsATest())
}
