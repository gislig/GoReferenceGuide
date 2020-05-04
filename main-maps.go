package main

import "fmt"

func main() {
	fmt.Println("maps")

	var itemlist map[string]string = make(map[string]string)

	itemlist["item1"] = "butter"
	itemlist["item2"] = "pickle"
	itemlist["item3"] = "bread"

	fmt.Println(itemlist)

	for key, _ := range itemlist {
		fmt.Println("The key :", key, "contains the value", itemlist[key])
	}
}
