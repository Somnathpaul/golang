package main

import (
	"fmt"
)

func main() {

	// simple if else
	if true {
		fmt.Println("hello golang")
	}

	if 'a' == 'A' {
		fmt.Println("true")
	} else if 'a' == 'a' {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
