package main

import (
	"fmt"
	"log"
	"regexp"
)

const validstring string = `^@?(\w){1,15}$`

func main() {

	fmt.Println("Testing regex")
	r, err := regexp.Compile(validstring)
	if err != nil {
		log.Fatal(err)
	}

	var invalid, valid string
	var result bool

	invalid = "bob#"
	valid = "bob"

	result = r.MatchString(invalid)
	fmt.Println(result)

	result = r.MatchString(valid)
	fmt.Println(result)
}
