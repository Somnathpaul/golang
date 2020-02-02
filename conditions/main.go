package main

import "fmt"

func main() {
	a := 100
	b := 20
	if a > b {
		fmt.Printf("%d is greater then %d \n", a, b)
	} else {
		fmt.Printf("%d is greater than %d \n", b, a)
	}
}
