package main

import (
	"fmt"
)

func main() {
	// ASCII code: 'h', always in single code with single character.
	fmt.Println("h:", 'h')
	fmt.Println("H :", 'H')
	//fmt.Println('hq') // Error, always single character, no double character allowed
	fmt.Println('i')
	fmt.Println('j')
}
